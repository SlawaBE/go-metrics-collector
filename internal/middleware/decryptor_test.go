package middleware

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/base64"
	"encoding/pem"
	"io"
	"math/big"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/SlawaBE/go-metrics-collector/internal/asymcrypto"
)

func generateTestKeys(t *testing.T) (certPath, keyPath string) {
	t.Helper()

	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("failed to generate RSA key: %v", err)
	}

	template := x509.Certificate{
		SerialNumber:          big.NewInt(1),
		Subject:               pkix.Name{Organization: []string{"Test"}},
		NotBefore:             time.Now().Add(-time.Hour),
		NotAfter:              time.Now().Add(time.Hour),
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
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
	if err := os.WriteFile(certPath, certPem, 0644); err != nil {
		t.Fatal(err)
	}
	keyPem := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(priv),
	})
	if err := os.WriteFile(keyPath, keyPem, 0600); err != nil {
		t.Fatal(err)
	}

	return certPath, keyPath
}

func encryptBody(t *testing.T, certPath string, plaintext []byte) *asymcrypto.HybridCipherText {
	t.Helper()
	enc, err := asymcrypto.NewRSAEncryptor(certPath)
	if err != nil {
		t.Fatalf("NewRSAEncryptor: %v", err)
	}
	hct, err := enc.Encrypt(plaintext)
	if err != nil {
		t.Fatalf("Encrypt: %v", err)
	}
	return hct
}

func makeEncryptedRequest(t *testing.T, hct *asymcrypto.HybridCipherText) *http.Request {
	t.Helper()
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(hct.CipherText))
	req.Header.Set("X-Nonce", base64.StdEncoding.EncodeToString(hct.Nonce))
	req.Header.Set("X-Encrypted-Key", base64.StdEncoding.EncodeToString(hct.EncryptedKey))
	return req
}

func TestNewDecryptor_Success(t *testing.T) {
	_, keyPath := generateTestKeys(t)

	d, err := NewDecryptor(keyPath)
	if err != nil {
		t.Fatalf("NewDecryptor: %v", err)
	}
	if d == nil || d.rsaDecryptor == nil {
		t.Fatal("NewDecryptor returned empty decryptor")
	}
}

func TestNewDecryptor_FileNotExist(t *testing.T) {
	_, err := NewDecryptor(filepath.Join(t.TempDir(), "nonexistent.pem"))
	if err == nil {
		t.Fatal("expected error for non-existent file, got nil")
	}
}

func TestMiddlewareDecryptor_NoHeaders_PassThrough(t *testing.T) {
	_, keyPath := generateTestKeys(t)
	d, err := NewDecryptor(keyPath)
	if err != nil {
		t.Fatal(err)
	}

	originalBody := []byte("plain body")
	var gotBody []byte
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, _ := io.ReadAll(r.Body)
		gotBody = b
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(originalBody))
	rec := httptest.NewRecorder()

	d.MiddlewareDecryptor(next).ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
	if !bytes.Equal(gotBody, originalBody) {
		t.Fatalf("body not passed through: got %q, want %q", gotBody, originalBody)
	}
}

func TestMiddlewareDecryptor_Success(t *testing.T) {
	certPath, keyPath := generateTestKeys(t)
	plaintext := []byte(`{"id":"gauge","value":42}`)
	hct := encryptBody(t, certPath, plaintext)

	d, err := NewDecryptor(keyPath)
	if err != nil {
		t.Fatal(err)
	}

	var gotBody []byte
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			t.Errorf("read body: %v", err)
		}
		gotBody = b
		w.WriteHeader(http.StatusOK)
	})

	req := makeEncryptedRequest(t, hct)
	rec := httptest.NewRecorder()

	d.MiddlewareDecryptor(next).ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200, body=%q", rec.Code, rec.Body.String())
	}
	if !bytes.Equal(gotBody, plaintext) {
		t.Fatalf("decrypted body mismatch\n got: %q\nwant: %q", gotBody, plaintext)
	}
}

func TestMiddlewareDecryptor_OnlyNonce_Fail(t *testing.T) {
	_, keyPath := generateTestKeys(t)
	d, err := NewDecryptor(keyPath)
	if err != nil {
		t.Fatal(err)
	}

	originalBody := []byte("plain text")
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.ReadAll(r.Body)
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(originalBody))
	req.Header.Set("X-Nonce", "abc")
	rec := httptest.NewRecorder()

	d.MiddlewareDecryptor(next).ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want 400", rec.Code)
	}
}

func TestMiddlewareDecryptor_OnlyEncryptedKey_PassThrough(t *testing.T) {
	_, keyPath := generateTestKeys(t)
	d, err := NewDecryptor(keyPath)
	if err != nil {
		t.Fatal(err)
	}

	originalBody := []byte("plain text")
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.ReadAll(r.Body)
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(originalBody))
	req.Header.Set("X-Encrypted-Key", "abc")
	rec := httptest.NewRecorder()

	d.MiddlewareDecryptor(next).ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want 400", rec.Code)
	}
}

func TestMiddlewareDecryptor_BadNonceBase64(t *testing.T) {
	_, keyPath := generateTestKeys(t)
	d, err := NewDecryptor(keyPath)
	if err != nil {
		t.Fatal(err)
	}

	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Fatal("next must not be called")
	})

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader([]byte("body")))
	req.Header.Set("X-Nonce", "!!! not base64 !!!")
	req.Header.Set("X-Encrypted-Key", base64.StdEncoding.EncodeToString([]byte("key")))
	rec := httptest.NewRecorder()

	d.MiddlewareDecryptor(next).ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want 400", rec.Code)
	}
}

func TestMiddlewareDecryptor_BadEncryptedKeyBase64(t *testing.T) {
	_, keyPath := generateTestKeys(t)
	d, err := NewDecryptor(keyPath)
	if err != nil {
		t.Fatal(err)
	}

	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Fatal("next must not be called")
	})

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader([]byte("body")))
	req.Header.Set("X-Nonce", base64.StdEncoding.EncodeToString([]byte("nonce")))
	req.Header.Set("X-Encrypted-Key", "!!! not base64 !!!")
	rec := httptest.NewRecorder()

	d.MiddlewareDecryptor(next).ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want 400", rec.Code)
	}
}

func TestMiddlewareDecryptor_TamperedCipherText(t *testing.T) {
	certPath, keyPath := generateTestKeys(t)
	hct := encryptBody(t, certPath, []byte("secret"))

	tampered := &asymcrypto.HybridCipherText{
		Nonce:        hct.Nonce,
		EncryptedKey: hct.EncryptedKey,
		CipherText:   append([]byte{}, hct.CipherText...),
	}
	tampered.CipherText[len(tampered.CipherText)-1] ^= 0xFF

	d, err := NewDecryptor(keyPath)
	if err != nil {
		t.Fatal(err)
	}

	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Fatal("next must not be called")
	})

	req := makeEncryptedRequest(t, tampered)
	rec := httptest.NewRecorder()

	d.MiddlewareDecryptor(next).ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want 400", rec.Code)
	}
}

func TestMiddlewareDecryptor_WrongKey(t *testing.T) {
	certPath, _ := generateTestKeys(t)
	_, otherKeyPath := generateTestKeys(t)

	hct := encryptBody(t, certPath, []byte("secret"))

	d, err := NewDecryptor(otherKeyPath)
	if err != nil {
		t.Fatal(err)
	}

	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Fatal("next must not be called")
	})

	req := makeEncryptedRequest(t, hct)
	rec := httptest.NewRecorder()

	d.MiddlewareDecryptor(next).ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want 400", rec.Code)
	}
}

func TestMiddlewareDecryptor_ReplacesBody(t *testing.T) {
	certPath, keyPath := generateTestKeys(t)
	plaintext := []byte("payload")
	hct := encryptBody(t, certPath, plaintext)

	d, err := NewDecryptor(keyPath)
	if err != nil {
		t.Fatal(err)
	}

	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			t.Errorf("read body in next: %v", err)
		}
		if !bytes.Equal(b, plaintext) {
			t.Errorf("next got wrong body: %q", b)
		}
		w.WriteHeader(http.StatusOK)
	})

	req := makeEncryptedRequest(t, hct)
	rec := httptest.NewRecorder()
	d.MiddlewareDecryptor(next).ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
}
