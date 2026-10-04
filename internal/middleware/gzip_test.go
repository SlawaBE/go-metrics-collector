package middleware

import (
	"bytes"
	gzstd "compress/gzip"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SlawaBE/go-metrics-collector/internal/gzip"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGZip_ResponseCompression(t *testing.T) {
	tests := []struct {
		name           string
		acceptEncoding string
		contentType    string
		body           string
		wantCompressed bool
	}{
		{
			name:           "json response is compressed",
			acceptEncoding: "gzip",
			contentType:    "application/json",
			body:           `{"id":"name","type":"gauge","value":1.1}`,
			wantCompressed: true,
		},
		{
			name:           "html response is compressed",
			acceptEncoding: "gzip",
			contentType:    "text/html",
			body:           "<html><body>hi</body></html>",
			wantCompressed: true,
		},
		{
			name:           "non-media response is not compressed",
			acceptEncoding: "gzip",
			contentType:    "text/plain",
			body:           "plain text",
			wantCompressed: false,
		},
		{
			name:           "client does not accept gzip",
			acceptEncoding: "",
			contentType:    "application/json",
			body:           `{"id":"name","type":"gauge","value":1.1}`,
			wantCompressed: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", tt.contentType)
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write([]byte(tt.body))
			})

			req := httptest.NewRequest(http.MethodGet, "/", nil)
			if tt.acceptEncoding != "" {
				req.Header.Set("Accept-Encoding", tt.acceptEncoding)
			}

			rec := httptest.NewRecorder()
			GZip(handler).ServeHTTP(rec, req)

			assert.Equal(t, http.StatusOK, rec.Code)

			if tt.wantCompressed {
				assert.Equal(t, "gzip", rec.Header().Get("Content-Encoding"))

				decompressed, err := decompress(rec.Body.Bytes())
				require.NoError(t, err)
				assert.Equal(t, tt.body, string(decompressed))
			} else {
				assert.Equal(t, "", rec.Header().Get("Content-Encoding"))
				assert.Equal(t, tt.body, rec.Body.String())
			}
		})
	}
}

func decompress(data []byte) ([]byte, error) {
	reader, err := gzstd.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	return io.ReadAll(reader)
}

func TestGZip_DecompressRequestBody(t *testing.T) {
	raw := []byte(`{"id":"name","type":"gauge","value":1.1}`)
	compressed, err := gzip.Compress(raw)
	require.NoError(t, err)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		require.NoError(t, err)
		w.WriteHeader(http.StatusOK)
		w.Write(body)
	})

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(compressed))
	req.Header.Set("Content-Encoding", "gzip")

	rec := httptest.NewRecorder()
	GZip(handler).ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, string(raw), rec.Body.String())
}
