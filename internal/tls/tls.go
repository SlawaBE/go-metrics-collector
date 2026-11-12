// Package tls provides utilities for generating self-signed certificates and
// loading gRPC TLS credentials.
package tls

import (
	crt "crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"os"

	"google.golang.org/grpc/credentials"
)

// LoadServerCredentials loads TLS credentials for the gRPC server from cert and key files.
// If caFile is non-empty, it enables mTLS by loading the CA cert for client verification.
func LoadServerCredentials(certFile, keyFile, caFile string) (credentials.TransportCredentials, error) {
	return loadCredentials(certFile, keyFile, caFile, false)
}

// LoadClientCredentials loads TLS credentials for the gRPC client from cert, key, and CA files.
func LoadClientCredentials(certFile, keyFile, caFile string) (credentials.TransportCredentials, error) {
	return loadCredentials(certFile, keyFile, caFile, true)
}

// loadCredentials builds gRPC TransportCredentials from cert/key files with optional CA.
// When isClient is true and no CA is given, InsecureSkipVerify is used so agent can
// connect with only cert+key and no mTLS requirement. When isClient is false, the
// server exposes its cert but skips client verification unless a CA is provided.
func loadCredentials(certFile, keyFile, caFile string, isClient bool) (credentials.TransportCredentials, error) {
	if certFile == "" || keyFile == "" {
		return nil, errors.New("tls: cert_file and key_file are required")
	}

	cert, err := crt.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, fmt.Errorf("tls: load key pair: %w", err)
	}

	cfg := &crt.Config{
		Certificates: []crt.Certificate{cert},
		MinVersion:   crt.VersionTLS12,
	}

	if caFile != "" {
		caCert, err := os.ReadFile(caFile)
		if err != nil {
			return nil, fmt.Errorf("tls: read CA cert: %w", err)
		}
		caCertPool := x509.NewCertPool()
		if !caCertPool.AppendCertsFromPEM(caCert) {
			return nil, errors.New("tls: failed to parse CA certificate")
		}
		if isClient {
			cfg.RootCAs = caCertPool
		} else {
			cfg.ClientAuth = crt.RequireAndVerifyClientCert
			cfg.ClientCAs = caCertPool
		}
	} else if isClient {
		cfg.InsecureSkipVerify = true
	}

	return credentials.NewTLS(cfg), nil
}
