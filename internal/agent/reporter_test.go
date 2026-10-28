package agent

import (
	"bytes"
	"compress/gzip"
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"io"
	"math/big"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/SlawaBE/go-metrics-collector/internal/asymcrypto"
	"github.com/SlawaBE/go-metrics-collector/internal/model"
	"github.com/SlawaBE/go-metrics-collector/internal/storage"
	"github.com/SlawaBE/go-metrics-collector/internal/utils/checksum"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func readGzipBody(t *testing.T, r *http.Request) []byte {
	t.Helper()
	reader, err := gzip.NewReader(r.Body)
	require.NoError(t, err)
	defer func() { _ = reader.Close() }()
	raw, err := io.ReadAll(reader)
	require.NoError(t, err)
	return raw
}

func TestReporter_Report_Success(t *testing.T) {
	const secret = "test-secret"

	var received []model.Metric
	var rawBody []byte
	var receivedSig []byte

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/updates", r.URL.Path)

		rawBody = readGzipBody(t, r)
		require.NoError(t, json.Unmarshal(rawBody, &received))

		sigHex := r.Header.Get("HashSHA256")
		sig, err := hex.DecodeString(sigHex)
		require.NoError(t, err)
		receivedSig = sig

		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	stor := storage.NewMemStorage()
	require.NoError(t, stor.UpdateMetric(context.Background(), model.NewCounterMetric("c1", 10)))
	require.NoError(t, stor.UpdateMetric(context.Background(), model.NewGaugeMetric("g1", 1.5)))

	address := strings.TrimPrefix(srv.URL, "http://")
	reporter := NewReporter(stor, nil, address, 1, secret, 2)

	reporter.Report()

	assert.Len(t, received, 2)

	expectedSig := checksum.Sign(rawBody, []byte(secret))
	assert.Equal(t, expectedSig, receivedSig, "signature should match the uncompressed JSON body")

	remaining, err := stor.GetValues(context.Background())
	require.NoError(t, err)
	assert.Empty(t, remaining)
}

func TestReporter_Report_EmptyStorage(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Error("should not send any metrics when storage is empty")
	}))
	defer srv.Close()

	stor := storage.NewMemStorage()
	address := strings.TrimPrefix(srv.URL, "http://")
	reporter := NewReporter(stor, nil, address, 1, "", 1)

	reporter.Report()
}

func TestReporter_sendMetrics_ServerError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer srv.Close()

	address := strings.TrimPrefix(srv.URL, "http://")
	reporter := NewReporter(nil, nil, address, 1, "", 1)

	err := reporter.sendMetrics(context.Background(), []model.Metric{model.NewCounterMetric("c1", 10)})
	assert.Error(t, err)
}

func TestReporter_sendMetrics_SuccessNoSignature(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/updates", r.URL.Path)
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	address := strings.TrimPrefix(srv.URL, "http://")
	reporter := NewReporter(nil, nil, address, 1, "", 1)

	err := reporter.sendMetrics(context.Background(), []model.Metric{model.NewCounterMetric("c1", 10)})
	assert.NoError(t, err)
}

func TestReporter_sendMetrics_WithSignature(t *testing.T) {
	const secret = "test-secret"

	var body []byte
	var signature []byte

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body = readGzipBody(t, r)
		signature, _ = hex.DecodeString(r.Header.Get("HashSHA256"))
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	address := strings.TrimPrefix(srv.URL, "http://")
	reporter := NewReporter(nil, nil, address, 1, secret, 1)

	require.NoError(t, reporter.sendMetrics(context.Background(), []model.Metric{model.NewCounterMetric("c1", 10)}))

	expectedSig := checksum.Sign(body, []byte(secret))
	assert.Equal(t, expectedSig, signature)
}

func generateTestKeys(t *testing.T) (certPath, keyPath string) {
	t.Helper()

	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("failed to generate RSA key: %v", err)
	}

	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			Organization: []string{"Test Org"},
		},
		NotBefore:             time.Now().Add(-time.Hour),
		NotAfter:              time.Now().Add(time.Hour),
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)
	if err != nil {
		t.Fatalf("failed to create certificate: %v", err)
	}

	dir := t.TempDir()
	certPath = filepath.Join(dir, "cert.pem")
	keyPath = filepath.Join(dir, "key.pem")

	certPem := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	if writeErr := os.WriteFile(certPath, certPem, 0644); writeErr != nil {
		t.Fatalf("failed to write cert: %v", writeErr)
	}

	keyBytes, err := x509.MarshalPKCS8PrivateKey(priv)
	if err != nil {
		t.Fatalf("failed to marshal key: %v", err)
	}
	keyPem := pem.EncodeToMemory(&pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: keyBytes,
	})
	if err := os.WriteFile(keyPath, keyPem, 0600); err != nil {
		t.Fatalf("failed to write key: %v", err)
	}

	return certPath, keyPath
}

func TestReporter_Report_Success_WithEncryption(t *testing.T) {
	const secret = "test-secret"

	certPath, keyPath := generateTestKeys(t)

	encryptor, err := asymcrypto.NewRSAEncryptor(certPath)
	require.NoError(t, err)
	decryptor, err := asymcrypto.NewRSADecryptor(keyPath)
	require.NoError(t, err)

	var (
		received        []model.Metric
		receivedSig     []byte
		receivedNonce   []byte
		receivedEncKey  []byte
		receivedBody    []byte
		receivedBodyRaw []byte
	)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/updates", r.URL.Path)

		nonceB64 := r.Header.Get("X-Nonce")
		encKeyB64 := r.Header.Get("X-Encrypted-Key")
		require.NotEmpty(t, nonceB64, "X-Nonce must be set")
		require.NotEmpty(t, encKeyB64, "X-Encrypted-Key must be set")

		var handlerErr error
		receivedNonce, handlerErr = base64.StdEncoding.DecodeString(nonceB64)
		require.NoError(t, handlerErr)
		receivedEncKey, handlerErr = base64.StdEncoding.DecodeString(encKeyB64)
		require.NoError(t, handlerErr)

		receivedBodyRaw, handlerErr = io.ReadAll(r.Body)
		require.NoError(t, handlerErr)
		require.NotEmpty(t, receivedBodyRaw)

		hct := &asymcrypto.HybridCipherText{
			CipherText:   receivedBodyRaw,
			Nonce:        receivedNonce,
			EncryptedKey: receivedEncKey,
		}
		decrypted, handlerErr := decryptor.Decrypt(hct)
		require.NoError(t, handlerErr, "server must decrypt the body")

		r.Body = io.NopCloser(bytes.NewReader(decrypted))
		receivedBody = readGzipBody(t, r)

		require.NoError(t, json.Unmarshal(receivedBody, &received))

		sigHex := r.Header.Get("HashSHA256")
		sig, handlerErr := hex.DecodeString(sigHex)
		require.NoError(t, handlerErr)
		receivedSig = sig

		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	stor := storage.NewMemStorage()
	require.NoError(t, stor.UpdateMetric(context.Background(), model.NewCounterMetric("c1", 10)))
	require.NoError(t, stor.UpdateMetric(context.Background(), model.NewGaugeMetric("g1", 1.5)))

	address := strings.TrimPrefix(srv.URL, "http://")
	reporter := NewReporter(stor, encryptor, address, 1, secret, 2)

	reporter.Report()

	assert.Len(t, received, 2)

	expectedSig := checksum.Sign(receivedBody, []byte(secret))
	assert.Equal(t, expectedSig, receivedSig,
		"signature must match the uncompressed, unencrypted JSON body")

	assert.NotEqual(t, receivedBody, receivedBodyRaw,
		"raw body must be ciphertext, not plaintext")
	assert.False(t, bytes.HasPrefix(receivedBodyRaw, []byte{0x1f, 0x8b}),
		"raw body must not start with gzip magic bytes")
	assert.False(t, json.Valid(receivedBodyRaw),
		"raw body must not be valid JSON")

	remaining, err := stor.GetValues(context.Background())
	require.NoError(t, err)
	assert.Empty(t, remaining)
}
