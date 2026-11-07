package middleware

import (
	"bytes"
	"encoding/base64"
	"io"
	"net/http"

	"github.com/SlawaBE/go-metrics-collector/internal/asymcrypto"
	"github.com/SlawaBE/go-metrics-collector/internal/logger"
	"go.uber.org/zap"
)

type Decryptor struct {
	rsaDecryptor *asymcrypto.RSADecryptor
}

func NewDecryptor(fileName string) (*Decryptor, error) {
	rsaDecryptor, err := asymcrypto.NewRSADecryptor(fileName)
	if err != nil {
		return nil, err
	}
	return &Decryptor{
		rsaDecryptor: rsaDecryptor,
	}, nil
}

func (d *Decryptor) MiddlewareDecryptor(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Nonce") == "" && r.Header.Get("X-Encrypted-Key") == "" {
			next.ServeHTTP(w, r)
			return
		}
		nonce, err := base64.StdEncoding.DecodeString(r.Header.Get("X-Nonce"))
		if err != nil {
			logger.Log.Error("error decode nonce", zap.Error(err))
			http.Error(w, "error decode nonce", http.StatusBadRequest)
			return
		}
		encryptedKey, err := base64.StdEncoding.DecodeString(r.Header.Get("X-Encrypted-Key"))
		if err != nil {
			logger.Log.Error("error decode encrypted key", zap.Error(err))
			http.Error(w, "error decode encrypted key", http.StatusBadRequest)
			return
		}
		body, err := io.ReadAll(r.Body)
		if err != nil {
			logger.Log.Error("error reading body", zap.Error(err))
			http.Error(w, "error reading body", http.StatusBadRequest)
			return
		}
		defer func() {
			if closeErr := r.Body.Close(); closeErr != nil {
				logger.Log.Error("failed to close request body", zap.Error(closeErr))
			}
		}()

		hct := &asymcrypto.HybridCipherText{
			CipherText:   body,
			Nonce:        nonce,
			EncryptedKey: encryptedKey,
		}

		body, err = d.rsaDecryptor.Decrypt(hct)
		if err != nil {
			logger.Log.Error("error decrypting body", zap.Error(err))
			http.Error(w, "error decrypting body", http.StatusBadRequest)
			return
		}
		r.Body = io.NopCloser(bytes.NewReader(body))
		next.ServeHTTP(w, r)
	})
}
