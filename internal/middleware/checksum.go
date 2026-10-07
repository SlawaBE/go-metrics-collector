package middleware

import (
	"bytes"
	"encoding/hex"
	"io"
	"net/http"

	"github.com/SlawaBE/go-metrics-collector/internal/logger"
	"github.com/SlawaBE/go-metrics-collector/internal/utils/checksum"
	"go.uber.org/zap"
)

// CheckSum verifies the integrity of request bodies and signs responses using
// HMAC SHA-256.
type CheckSum struct {
	secretKey []byte
}

// NewCheckSum creates a signature verification middleware with the given
// secret key.
func NewCheckSum(secretKey string) *CheckSum {
	return &CheckSum{
		secretKey: []byte(secretKey),
	}
}

type responseWriterWrapper struct {
	http.ResponseWriter
	body   *bytes.Buffer
	status int
}

func (ww *responseWriterWrapper) WriteHeader(statusCode int) {
	ww.status = statusCode
}

func (ww *responseWriterWrapper) Write(data []byte) (int, error) {
	return ww.body.Write(data)
}

// CheckSumMiddleware wraps an HTTP handler: for requests (except GET) it
// verifies the HashSHA256 header, and signs responses with the same key.
func (c *CheckSum) CheckSumMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			signatureHeader := r.Header.Get("HashSHA256")
			if signatureHeader == "" {
				logger.Log.Error("Header 'HashSHA256' does not exist")
				http.Error(w, "Header 'HashSHA256' does not exist", http.StatusBadRequest)
				return
			}

			bodyBytes, err := io.ReadAll(r.Body)
			if err != nil {
				logger.Log.Error("Failed to read request body", zap.Error(err))
				http.Error(w, "Failed to read request body", http.StatusInternalServerError)
				return
			}
			r.Body.Close()
			r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

			sign, err := hex.DecodeString(signatureHeader)
			if err != nil {
				logger.Log.Error("Invalid signature format", zap.Error(err))
				http.Error(w, "Invalid signature format", http.StatusBadRequest)
				return
			}

			if !checksum.Check(bodyBytes, sign, c.secretKey) {
				logger.Log.Error("Invalid signature", zap.Error(err))
				http.Error(w, "Invalid signature", http.StatusBadRequest)
				return
			}
		}

		ww := &responseWriterWrapper{
			ResponseWriter: w,
			body:           new(bytes.Buffer),
			status:         http.StatusOK,
		}

		handler.ServeHTTP(ww, r)

		response := ww.body.Bytes()
		respSig := checksum.Sign(response, c.secretKey)
		w.Header().Set("HashSHA256", hex.EncodeToString(respSig))

		w.WriteHeader(ww.status)
		w.Write(response)
	})
}
