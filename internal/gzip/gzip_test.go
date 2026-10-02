package gzip

import (
	"bytes"
	gzstd "compress/gzip"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func decompress(t *testing.T, data []byte) []byte {
	t.Helper()
	reader, err := gzstd.NewReader(bytes.NewReader(data))
	require.NoError(t, err)
	defer reader.Close()
	raw, err := io.ReadAll(reader)
	require.NoError(t, err)
	return raw
}

func TestCompress(t *testing.T) {
	data := []byte("hello world")

	compressed, err := Compress(data)
	require.NoError(t, err)

	assert.NotEqual(t, data, compressed)
	assert.Equal(t, data, decompress(t, compressed))
}

func TestCompress_EmptyData(t *testing.T) {
	compressed, err := Compress(nil)
	require.NoError(t, err)

	assert.Empty(t, decompress(t, compressed))
}

func TestCompressWriter_Write(t *testing.T) {
	rec := httptest.NewRecorder()
	writer := NewCompressWriter(rec)

	_, err := writer.Write([]byte("payload"))
	require.NoError(t, err)
	require.NoError(t, writer.Close())

	assert.Equal(t, "payload", string(decompress(t, rec.Body.Bytes())))
}

func TestCompressWriter_WriteHeader_SuccessStatus(t *testing.T) {
	rec := httptest.NewRecorder()
	writer := NewCompressWriter(rec)

	writer.WriteHeader(http.StatusOK)
	_, _ = writer.Write([]byte("body"))
	require.NoError(t, writer.Close())

	assert.Equal(t, "gzip", rec.Header().Get("Content-Encoding"))
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestCompressWriter_WriteHeader_ErrorStatus(t *testing.T) {
	rec := httptest.NewRecorder()
	writer := NewCompressWriter(rec)

	writer.WriteHeader(http.StatusInternalServerError)
	_, _ = writer.Write([]byte("body"))
	require.NoError(t, writer.Close())

	assert.Equal(t, "", rec.Header().Get("Content-Encoding"))
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
}

func TestCompressReader_Read(t *testing.T) {
	plain := []byte("compressed payload")

	compressed, err := Compress(plain)
	require.NoError(t, err)

	reader, err := NewCompressReader(io.NopCloser(bytes.NewReader(compressed)))
	require.NoError(t, err)

	out, err := io.ReadAll(reader)
	require.NoError(t, err)
	require.NoError(t, reader.Close())

	assert.Equal(t, plain, out)
}

func TestCompressReader_InvalidData(t *testing.T) {
	reader, err := NewCompressReader(io.NopCloser(bytes.NewReader([]byte("not gzip data"))))
	assert.Error(t, err)
	assert.Nil(t, reader)
}

func TestCompressWriter_ImplementsHeaderInterface(t *testing.T) {
	rec := httptest.NewRecorder()
	writer := NewCompressWriter(rec)

	writer.Header().Set("X-Custom", "value")
	assert.Equal(t, "value", rec.Header().Get("X-Custom"))
}