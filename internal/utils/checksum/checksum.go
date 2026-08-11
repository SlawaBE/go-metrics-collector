package checksum

import (
	"crypto/hmac"
	"crypto/sha256"
)

func Sign(data []byte, secretKey []byte) []byte {
	h := hmac.New(sha256.New, secretKey)
	h.Write(data)
	return h.Sum(nil)
}

func Check(data []byte, sign []byte, secretKey []byte) bool {
	s := Sign(data, secretKey)
	return hmac.Equal(s, sign)
}
