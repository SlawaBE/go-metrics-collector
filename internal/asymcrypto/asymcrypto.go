// Package asymcrypto provides functionality for asymmetric encryption.
//
// A hybrid scheme with AES‑GCM symmetric key generation is used, since RSA can only encrypt data with a length of no more than 190 bytes.
package asymcrypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"errors"
	"io"

	"crypto/rand"
	"encoding/pem"
	"os"
)

// HybridCipherText is structure for hybrid encryption components
type HybridCipherText struct {
	// AES key, encrypted with RSA
	EncryptedKey []byte
	// Nonce for AES-GCM
	Nonce []byte
	// Cipher text
	CipherText []byte
}

type RSAEncryptor struct {
	cert *x509.Certificate
}

func NewRSAEncryptor(fileName string) (*RSAEncryptor, error) {
	cert, err := readPublicKey(fileName)
	if err != nil {
		return nil, err
	}
	return &RSAEncryptor{
		cert: cert,
	}, nil
}

func readPublicKey(fileName string) (*x509.Certificate, error) {
	certificateBytes, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	certificatePemBlock, _ := pem.Decode(certificateBytes)
	if certificatePemBlock == nil {
		return nil, errors.New("failed to decode PEM block containing certificate")
	}
	certificate, err := x509.ParseCertificate(certificatePemBlock.Bytes)
	if err != nil {
		return nil, err
	}
	return certificate, nil
}

// Encrypt encrypts data with AES-GCM, and encrypts AES key using a public RSA key
func (re *RSAEncryptor) Encrypt(plaintext []byte) (*HybridCipherText, error) {
	aesKey := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, aesKey); err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, err
	}
	cipherText := gcm.Seal(nil, nonce, plaintext, nil)

	encryptedKey, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		re.cert.PublicKey.(*rsa.PublicKey),
		aesKey,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &HybridCipherText{
		EncryptedKey: encryptedKey,
		Nonce:        nonce,
		CipherText:   cipherText,
	}, nil
}

type RSADecryptor struct {
	pKey *rsa.PrivateKey
}

func NewRSADecryptor(fileName string) (*RSADecryptor, error) {
	pKey, err := readPrivateKey(fileName)
	if err != nil {
		return nil, err
	}
	return &RSADecryptor{
		pKey: pKey,
	}, nil
}

func readPrivateKey(fileName string) (*rsa.PrivateKey, error) {
	privateKeyBytes, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	privateKeyPemBlock, _ := pem.Decode(privateKeyBytes)
	if privateKeyPemBlock == nil {
		return nil, errors.New("failed to decode PEM block containing private key")
	}
	privateKeyAny, err := x509.ParsePKCS8PrivateKey(privateKeyPemBlock.Bytes)
	if err != nil {
		return nil, err
	}
	privateKey, ok := privateKeyAny.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("decoded key is not an RSA private key")
	}
	return privateKey, nil
}

// Decrypt decrypts AES key using a private RSA key and then decrypts data with AES-GCM.
func (rd *RSADecryptor) Decrypt(hct *HybridCipherText) ([]byte, error) {
	aesKey, err := rsa.DecryptOAEP(
		sha256.New(),
		rand.Reader,
		rd.pKey,
		hct.EncryptedKey,
		nil,
	)
	if err != nil {
		return nil, err
	}
	if len(aesKey) != 32 {
		return nil, errors.New("invalid AES key size")
	}

	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	plaintext, err := gcm.Open(nil, hct.Nonce, hct.CipherText, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
