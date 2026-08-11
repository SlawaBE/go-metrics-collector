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

type CheckSum struct {
	secretKey []byte
}

func NewCheckSum(secretKey string) *CheckSum {
	return &CheckSum{
		secretKey: []byte(secretKey),
	}
}

type responseWriterWrapper struct {
	http.ResponseWriter
	body *bytes.Buffer
}

func (ww *responseWriterWrapper) Write(data []byte) (int, error) {
	return ww.body.Write(data)
}

func (c *CheckSum) CheckSumMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		signatureHeader := r.Header.Get("HashSHA256")
		if signatureHeader != "" {
			// if signatureHeader == "" {
			// 	logger.Log.Error("Header 'HashSHA256' does not exist", zap.Error(err))
			// 	http.Error(w, "Header 'HashSHA256' does not exist", http.StatusBadRequest)
			// 	return
			// }

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
		}

		handler.ServeHTTP(ww, r)

		response := ww.body.Bytes()
		respSig := checksum.Sign(response, c.secretKey)
		w.Header()["HashSHA256"] = []string{hex.EncodeToString(respSig)}

		w.Write(response)
	})
}
