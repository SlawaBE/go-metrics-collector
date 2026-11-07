package asymcrypto

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func generateTestKeys(t *testing.T) (certPath, keyPath string) {
	t.Helper()

	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("failed to generate RSA key: %v", err)
	}

	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			Organization: []string{"Test Org"},
		},
		NotBefore:             time.Now().Add(-time.Hour),
		NotAfter:              time.Now().Add(time.Hour),
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)
	if err != nil {
		t.Fatalf("failed to create certificate: %v", err)
	}

	dir := t.TempDir()
	certPath = filepath.Join(dir, "cert.pem")
	keyPath = filepath.Join(dir, "key.pem")

	certPem := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	if writeErr := os.WriteFile(certPath, certPem, 0644); writeErr != nil {
		t.Fatalf("failed to write cert: %v", writeErr)
	}

	keyBytes, err := x509.MarshalPKCS8PrivateKey(priv)
	if err != nil {
		t.Fatalf("failed to marshal key: %v", err)
	}
	keyPem := pem.EncodeToMemory(&pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: keyBytes,
	})
	if err := os.WriteFile(keyPath, keyPem, 0600); err != nil {
		t.Fatalf("failed to write key: %v", err)
	}

	return certPath, keyPath
}

func TestNewRSAEncryptor_Success(t *testing.T) {
	certPath, _ := generateTestKeys(t)

	enc, err := NewRSAEncryptor(certPath)
	if err != nil {
		t.Fatalf("NewRSAEncryptor returned error: %v", err)
	}
	if enc == nil {
		t.Fatal("NewRSAEncryptor returned nil encryptor")
	}
	if enc.cert == nil {
		t.Fatal("certificate not set")
	}
}

func TestNewRSAEncryptor_FileNotExist(t *testing.T) {
	_, err := NewRSAEncryptor(filepath.Join(t.TempDir(), "nonexistent.pem"))
	if err == nil {
		t.Fatal("expected error for non-existent file, got nil")
	}
}

func TestNewRSAEncryptor_InvalidPEM(t *testing.T) {
	dir := t.TempDir()
	badPath := filepath.Join(dir, "bad.pem")
	if err := os.WriteFile(badPath, []byte("not a pem file"), 0644); err != nil {
		t.Fatal(err)
	}

	_, err := NewRSAEncryptor(badPath)
	if err == nil {
		t.Fatal("expected error for invalid PEM, got nil")
	}
}

func TestNewRSADecryptor_Success(t *testing.T) {
	_, keyPath := generateTestKeys(t)

	dec, err := NewRSADecryptor(keyPath)
	if err != nil {
		t.Fatalf("NewRSADecryptor returned error: %v", err)
	}
	if dec == nil {
		t.Fatal("NewRSADecryptor returned nil decryptor")
	}
	if dec.pKey == nil {
		t.Fatal("private key not set")
	}
}

func TestNewRSADecryptor_FileNotExist(t *testing.T) {
	_, err := NewRSADecryptor(filepath.Join(t.TempDir(), "nonexistent.pem"))
	if err == nil {
		t.Fatal("expected error for non-existent file, got nil")
	}
}

func TestNewRSADecryptor_InvalidPEM(t *testing.T) {
	dir := t.TempDir()
	badPath := filepath.Join(dir, "bad.pem")
	if err := os.WriteFile(badPath, []byte("not a pem file"), 0644); err != nil {
		t.Fatal(err)
	}

	_, err := NewRSADecryptor(badPath)
	if err == nil {
		t.Fatal("expected error for invalid PEM, got nil")
	}
}

func TestEncryptDecrypt_RoundTrip(t *testing.T) {
	certPath, keyPath := generateTestKeys(t)

	enc, err := NewRSAEncryptor(certPath)
	if err != nil {
		t.Fatalf("NewRSAEncryptor: %v", err)
	}
	dec, err := NewRSADecryptor(keyPath)
	if err != nil {
		t.Fatalf("NewRSADecryptor: %v", err)
	}

	cases := map[string][]byte{
		"empty":        {},
		"short text":   []byte("hello"),
		"russian text": []byte("Привет, мир!"),
		"large data":   bytes.Repeat([]byte("A"), 1024*1024), // 1MB
	}

	for name, plaintext := range cases {
		t.Run(name, func(t *testing.T) {
			hct, err := enc.Encrypt(plaintext)
			if err != nil {
				t.Fatalf("Encrypt: %v", err)
			}
			if hct == nil {
				t.Fatal("Encrypt returned nil")
			}
			if len(hct.Nonce) == 0 {
				t.Fatal("Nonce is empty")
			}
			if len(hct.EncryptedKey) == 0 {
				t.Fatal("EncryptedKey is empty")
			}

			if len(plaintext) > 0 && bytes.Equal(hct.CipherText, plaintext) {
				t.Fatal("ciphertext equals plaintext")
			}

			decrypted, err := dec.Decrypt(hct)
			if err != nil {
				t.Fatalf("Decrypt: %v", err)
			}

			if !bytes.Equal(decrypted, plaintext) {
				t.Fatalf("decrypted != plaintext\n got: %q\nwant: %q", decrypted, plaintext)
			}
		})
	}
}

func TestEncrypt_DifferentNonces(t *testing.T) {
	certPath, _ := generateTestKeys(t)
	enc, err := NewRSAEncryptor(certPath)
	if err != nil {
		t.Fatal(err)
	}

	plaintext := []byte("same data")

	hct1, err := enc.Encrypt(plaintext)
	if err != nil {
		t.Fatal(err)
	}
	hct2, err := enc.Encrypt(plaintext)
	if err != nil {
		t.Fatal(err)
	}

	if bytes.Equal(hct1.Nonce, hct2.Nonce) {
		t.Fatal("nonces are equal for two encryptions — nonce reuse detected")
	}
	if bytes.Equal(hct1.CipherText, hct2.CipherText) {
		t.Fatal("ciphertexts are equal for two encryptions — nonce reuse detected")
	}
	if bytes.Equal(hct1.EncryptedKey, hct2.EncryptedKey) {
		t.Fatal("encrypted keys are equal — AES key reuse detected")
	}
}

func TestDecrypt_TamperedCipherText(t *testing.T) {
	certPath, keyPath := generateTestKeys(t)

	enc, err := NewRSAEncryptor(certPath)
	if err != nil {
		t.Fatal(err)
	}
	dec, err := NewRSADecryptor(keyPath)
	if err != nil {
		t.Fatal(err)
	}

	hct, err := enc.Encrypt([]byte("secret message"))
	if err != nil {
		t.Fatal(err)
	}

	tampered := *hct
	tampered.CipherText = append([]byte{}, hct.CipherText...)
	tampered.CipherText[len(tampered.CipherText)-1] ^= 0xFF

	if _, err := dec.Decrypt(&tampered); err == nil {
		t.Fatal("expected authentication error for tampered ciphertext, got nil")
	}
}

func TestDecrypt_TamperedNonce(t *testing.T) {
	certPath, keyPath := generateTestKeys(t)

	enc, err := NewRSAEncryptor(certPath)
	if err != nil {
		t.Fatal(err)
	}
	dec, err := NewRSADecryptor(keyPath)
	if err != nil {
		t.Fatal(err)
	}

	hct, err := enc.Encrypt([]byte("secret message"))
	if err != nil {
		t.Fatal(err)
	}

	tampered := *hct
	tampered.Nonce = append([]byte{}, hct.Nonce...)
	tampered.Nonce[0] ^= 0xFF

	if _, err := dec.Decrypt(&tampered); err == nil {
		t.Fatal("expected error for tampered nonce, got nil")
	}
}

func TestDecrypt_TamperedEncryptedKey(t *testing.T) {
	certPath, keyPath := generateTestKeys(t)

	enc, err := NewRSAEncryptor(certPath)
	if err != nil {
		t.Fatal(err)
	}
	dec, err := NewRSADecryptor(keyPath)
	if err != nil {
		t.Fatal(err)
	}

	hct, err := enc.Encrypt([]byte("secret message"))
	if err != nil {
		t.Fatal(err)
	}

	tampered := *hct
	tampered.EncryptedKey = append([]byte{}, hct.EncryptedKey...)
	tampered.EncryptedKey[0] ^= 0xFF

	if _, err := dec.Decrypt(&tampered); err == nil {
		t.Fatal("expected error for tampered encrypted key, got nil")
	}
}

func TestDecrypt_WrongKey(t *testing.T) {
	certPath, _ := generateTestKeys(t)
	_, otherKeyPath := generateTestKeys(t)

	enc, err := NewRSAEncryptor(certPath)
	if err != nil {
		t.Fatal(err)
	}
	otherDec, err := NewRSADecryptor(otherKeyPath)
	if err != nil {
		t.Fatal(err)
	}

	hct, err := enc.Encrypt([]byte("secret message"))
	if err != nil {
		t.Fatal(err)
	}

	if _, err := otherDec.Decrypt(hct); err == nil {
		t.Fatal("expected error when decrypting with wrong key, got nil")
	}
}

func TestDecrypt_InvalidAESKeySize(t *testing.T) {
	certPath, keyPath := generateTestKeys(t)

	enc, err := NewRSAEncryptor(certPath)
	if err != nil {
		t.Fatal(err)
	}
	dec, err := NewRSADecryptor(keyPath)
	if err != nil {
		t.Fatal(err)
	}

	hct, err := enc.Encrypt([]byte("secret message"))
	if err != nil {
		t.Fatal(err)
	}

	wrongKey := make([]byte, 16)
	_, err = rand.Read(wrongKey)
	if err != nil {
		t.Fatal(err)
	}
	cert, err := readPublicKey(certPath)
	if err != nil {
		t.Fatal(err)
	}
	encKey, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		cert.PublicKey.(*rsa.PublicKey),
		wrongKey,
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}

	_ = encKey
	badHCT := *hct
	badHCT.EncryptedKey = encKey
	_, err = dec.Decrypt(&badHCT)

	if err == nil {
		t.Fatal("expected error for invalid encrypted key, got nil")
	}
	if err.Error() != "invalid AES key size" {
		t.Fatal("expecting 'invalid AES key size' error")
	}
}
