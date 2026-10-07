// Package checksum provides functions for signing and verifying data using
// HMAC SHA-256.
package checksum

import (
	"crypto/hmac"
	"crypto/sha256"
)

// Sign computes the HMAC SHA-256 signature of the data using the given secret key.
func Sign(data []byte, secretKey []byte) []byte {
	h := hmac.New(sha256.New, secretKey)
	h.Write(data)
	return h.Sum(nil)
}

// Check verifies that the signature of the data matches the expected one.
func Check(data []byte, sign []byte, secretKey []byte) bool {
	s := Sign(data, secretKey)
	return hmac.Equal(s, sign)
}
