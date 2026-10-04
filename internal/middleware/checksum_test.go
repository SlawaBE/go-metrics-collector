package middleware

import (
	"bytes"
	"encoding/hex"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SlawaBE/go-metrics-collector/internal/utils/checksum"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCheckSumMiddleware_RequestSignature(t *testing.T) {
	const secret = "test-secret"
	middleware := NewCheckSum(secret)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	tests := []struct {
		name       string
		body       []byte
		signature  func() string
		wantStatus int
	}{
		{
			name: "valid signature",
			body: []byte(`{"id":"name","type":"gauge","value":1.1}`),
			signature: func() string {
				return hex.EncodeToString(checksum.Sign([]byte(`{"id":"name","type":"gauge","value":1.1}`), []byte(secret)))
			},
			wantStatus: http.StatusOK,
		},
		{
			name:       "missing header",
			body:       []byte(`{"id":"name","type":"gauge","value":1.1}`),
			signature:  func() string { return "" },
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "invalid hex format",
			body:       []byte(`{"id":"name","type":"gauge","value":1.1}`),
			signature:  func() string { return "not-a-hex" },
			wantStatus: http.StatusBadRequest,
		},
		{
			name: "invalid signature",
			body: []byte(`{"id":"name","type":"gauge","value":1.1}`),
			signature: func() string {
				return hex.EncodeToString(checksum.Sign([]byte(`{"id":"name","type":"gauge","value":2.2}`), []byte(secret)))
			},
			wantStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var req *http.Request
			if tt.signature() == "" {
				req = httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(tt.body))
			} else {
				req = httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(tt.body))
				req.Header.Set("HashSHA256", tt.signature())
			}

			rec := httptest.NewRecorder()
			middleware.CheckSumMiddleware(handler).ServeHTTP(rec, req)

			assert.Equal(t, tt.wantStatus, rec.Code)
		})
	}
}

func TestCheckSumMiddleware_ResponseSignature(t *testing.T) {
	const secret = "test-secret"
	middleware := NewCheckSum(secret)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"id":"name","type":"gauge","value":1.1}`))
	})

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	middleware.CheckSumMiddleware(handler).ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	respSig, err := hex.DecodeString(rec.Header().Get("HashSHA256"))
	require.NoError(t, err)

	expected := checksum.Sign([]byte(`{"id":"name","type":"gauge","value":1.1}`), []byte(secret))
	assert.Equal(t, expected, respSig, "response signature does not match body")
	assert.JSONEq(t, `{"id":"name","type":"gauge","value":1.1}`, rec.Body.String())
}
