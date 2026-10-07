package middleware

import (
	"compress/gzip"
	"io"
	"mime"
	"net/http"
	"strings"
	"sync"

	internalgzip "github.com/SlawaBE/go-metrics-collector/internal/gzip"
	"github.com/SlawaBE/go-metrics-collector/internal/logger"
	"go.uber.org/zap"
)

var gzipWriterPool = sync.Pool{
	New: func() any {
		w, _ := gzip.NewWriterLevel(io.Discard, gzip.BestSpeed)
		return w
	},
}

type gzipResponseWriter struct {
	http.ResponseWriter
	compress bool
	zw       *gzip.Writer
	started  bool
}

func newGzipResponseWriter(w http.ResponseWriter) *gzipResponseWriter {
	return &gzipResponseWriter{ResponseWriter: w}
}

func (rw *gzipResponseWriter) start() {
	if rw.started {
		return
	}
	rw.started = true

	if isCompressible(rw.Header().Get("Content-Type")) {
		rw.compress = true
		rw.Header().Set("Content-Encoding", "gzip")
		rw.Header().Del("Content-Length")
		rw.zw = gzipWriterPool.Get().(*gzip.Writer)
		rw.zw.Reset(rw.ResponseWriter)
	}
}

func (rw *gzipResponseWriter) WriteHeader(statusCode int) {
	if rw.started {
		return
	}
	rw.start()
	rw.ResponseWriter.WriteHeader(statusCode)
}

func (rw *gzipResponseWriter) Write(b []byte) (int, error) {
	rw.start()
	if rw.compress {
		return rw.zw.Write(b)
	}
	return rw.ResponseWriter.Write(b)
}

func (rw *gzipResponseWriter) Close() error {
	if rw.zw == nil {
		return nil
	}
	err := rw.zw.Close()
	gzipWriterPool.Put(rw.zw)
	rw.zw = nil
	return err
}

func isCompressible(contentType string) bool {
	if contentType == "" {
		return false
	}
	mediaType, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		return false
	}
	return mediaType == "application/json" || mediaType == "text/html"
}

// GZip wraps an HTTP handler: decompresses compressed requests and compresses
// responses when the client supports gzip.
func GZip(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.Header.Get("Content-Encoding"), "gzip") {
			gzipReader, err := internalgzip.NewCompressReader(r.Body)
			if err != nil {
				http.Error(w, "Failed to decompress request body", http.StatusBadRequest)
				return
			}
			defer gzipReader.Close()
			r.Body = io.NopCloser(gzipReader)
		}

		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			handler.ServeHTTP(w, r)
			return
		}

		rw := newGzipResponseWriter(w)
		handler.ServeHTTP(rw, r)
		if rw.zw != nil {
			if err := rw.Close(); err != nil {
				logger.Log.Error("failed to compress response", zap.Error(err))
			}
		}
	})
}
