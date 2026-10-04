package agent

import (
	"compress/gzip"
	"context"
	"encoding/hex"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

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
	defer reader.Close()
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
	reporter := NewReporter(stor, address, 1, secret, 2)

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
	reporter := NewReporter(stor, address, 1, "", 1)

	reporter.Report()
}

func TestReporter_sendMetrics_ServerError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer srv.Close()

	address := strings.TrimPrefix(srv.URL, "http://")
	reporter := NewReporter(nil, address, 1, "", 1)

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
	reporter := NewReporter(nil, address, 1, "", 1)

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
	reporter := NewReporter(nil, address, 1, secret, 1)

	require.NoError(t, reporter.sendMetrics(context.Background(), []model.Metric{model.NewCounterMetric("c1", 10)}))

	expectedSig := checksum.Sign(body, []byte(secret))
	assert.Equal(t, expectedSig, signature)
}
