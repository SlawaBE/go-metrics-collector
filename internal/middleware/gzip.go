package middleware

import (
	"io"
	"net/http"
	"strings"

	"github.com/SlawaBE/go-metrics-collector/internal/gzip"
)

func GZip(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.Header.Get("Content-Encoding"), "gzip") {
			gzipReader, err := gzip.NewCompressReader(r.Body)
			if err != nil {
				http.Error(w, "Failed to decompress request body", http.StatusBadRequest)
				return
			}
			defer gzipReader.Close()
			r.Body = io.NopCloser(gzipReader)
		}

		if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			gzipWriter := gzip.NewCompressWriter(w)
			defer gzipWriter.Close()
			w.Header().Set("Content-Encoding", "gzip")

			handler.ServeHTTP(gzipWriter, r)
			return
		}
		handler.ServeHTTP(w, r)
	})
}
