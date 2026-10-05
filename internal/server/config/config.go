// Package config holds the metrics server configuration, populated from
// command-line flags and environment variables.
package config

import (
	"flag"

	"github.com/caarlos0/env/v11"
)

// Config describes the metrics server startup parameters.
type Config struct {
	// ServerAddress is the HTTP server address and port.
	ServerAddress string `env:"ADDRESS"`
	// StoreInterval is the interval (in seconds) for saving metrics to a file.
	StoreInterval int `env:"STORE_INTERVAL"`
	// FileStoragePath is the path to the file used for persisting metrics.
	FileStoragePath string `env:"FILE_STORAGE_PATH"`
	// Restore indicates whether to restore metrics from the file on startup.
	Restore bool `env:"RESTORE"`
	// DatabaseDSN is the PostgreSQL connection string; when set, the DB is used.
	DatabaseDSN string `env:"DATABASE_DSN"`
	// Key is the secret key for signing and verifying HMAC SHA-256 requests.
	Key string `env:"KEY"`
	// AuditFile is the path to the file for writing audit events.
	AuditFile string `env:"AUDIT_FILE"`
	// AuditURL is the URL for sending audit events.
	AuditURL string `env:"AUDIT_URL"`
	// ProfileEnabled indicates whether pprof profiling is enabled.
	ProfileEnabled bool `env:"PROFILE_ENABLED"`
}

// ReadConfig populates the configuration from command-line flags and
// environment variables (environment takes precedence).
func ReadConfig() (Config, error) {
	var config Config

	flag.StringVar(&config.ServerAddress, "a", "localhost:8080", "address and port to run server")
	flag.IntVar(&config.StoreInterval, "i", 300, "interval for saving metric to file")
	flag.StringVar(&config.FileStoragePath, "f", "storage.json", "file for saving metrics")
	flag.BoolVar(&config.Restore, "r", false, "restore metric from file")
	flag.StringVar(&config.DatabaseDSN, "d", "", "database url")
	flag.StringVar(&config.Key, "k", "", "key for signing request with HMAC SHA-256")
	flag.StringVar(&config.AuditFile, "audit-file", "", "file for writing audit event")
	flag.StringVar(&config.AuditURL, "audit-url", "", "url for sending audit event")
	flag.BoolVar(&config.ProfileEnabled, "p", false, "enable profiling")
	flag.Parse()

	err := env.Parse(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}
