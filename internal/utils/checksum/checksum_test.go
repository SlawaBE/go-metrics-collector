package checksum

import (
	"crypto/hmac"
	"crypto/sha256"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSign(t *testing.T) {
	data := []byte("some test data")
	secretKey := []byte("test-secret")

	sign := Sign(data, secretKey)

	h := hmac.New(sha256.New, secretKey)
	h.Write(data)
	assert.Equal(t, h.Sum(nil), sign)
}

func TestSign_Deterministic(t *testing.T) {
	data := []byte("some test data")
	secretKey := []byte("test-secret")

	first := Sign(data, secretKey)
	second := Sign(data, secretKey)

	assert.Equal(t, first, second)
}

func TestSign_DifferentKeyProducesDifferentSignature(t *testing.T) {
	data := []byte("some test data")

	withKeyA := Sign(data, []byte("key-a"))
	withKeyB := Sign(data, []byte("key-b"))

	assert.NotEqual(t, withKeyA, withKeyB)
}

func TestSign_DifferentDataProducesDifferentSignature(t *testing.T) {
	secretKey := []byte("test-secret")

	signA := Sign([]byte("data one"), secretKey)
	signB := Sign([]byte("data two"), secretKey)

	assert.NotEqual(t, signA, signB)
}

func TestCheck(t *testing.T) {
	data := []byte("some test data")
	secretKey := []byte("test-secret")

	sign := Sign(data, secretKey)

	assert.True(t, Check(data, sign, secretKey))
}

func TestCheck_InvalidSignature(t *testing.T) {
	data := []byte("some test data")
	secretKey := []byte("test-secret")

	other := Sign([]byte("other data"), secretKey)

	assert.False(t, Check(data, other, secretKey))
}

func TestCheck_WrongKey(t *testing.T) {
	data := []byte("some test data")

	sign := Sign(data, []byte("real-secret"))

	assert.False(t, Check(data, sign, []byte("wrong-secret")))
}

func TestCheck_TamperedData(t *testing.T) {
	secretKey := []byte("test-secret")

	original := Sign([]byte("original data"), secretKey)

	assert.False(t, Check([]byte("tampered-data"), original, secretKey))
}

func TestCheck_EmptyData(t *testing.T) {
	secretKey := []byte("test-secret")

	sign := Sign(nil, secretKey)

	assert.True(t, Check(nil, sign, secretKey))
}

func TestCheck_EmptySecretKey(t *testing.T) {
	sign := Sign([]byte("data"), nil)

	assert.True(t, Check([]byte("data"), sign, nil))
}
