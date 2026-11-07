// Package config holds the metrics server configuration, populated from
// command-line flags and environment variables.
package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/SlawaBE/go-metrics-collector/internal/utils/config"
	"github.com/caarlos0/env/v11"
)

const (
	defaultServerAddress   = "localhost:8080"
	defaultFileStoragePath = "storage.json"
	defaultDatabaseDSN     = ""
	defaultKey             = ""
	defaultAuditFile       = ""
	defaultAuditURL        = ""
	defaultProfileAddress  = "localhost:8085"
	defaultCryptoKey       = ""
	defaultStoreInterval   = config.Duration(time.Second * 300)
	defaultAuditQueueSize  = 100
	defaultRestore         = false
	defaultProfileEnabled  = false
	defaultConfig          = ""
)

// Config describes the metrics server startup parameters.
// generate:reset
type Config struct {
	// ServerAddress is the HTTP server address and port.
	ServerAddress string `env:"ADDRESS" json:"address"`
	// FileStoragePath is the path to the file used for persisting metrics.
	FileStoragePath string `env:"FILE_STORAGE_PATH" json:"storage_pass"`
	// DatabaseDSN is the PostgreSQL connection string; when set, the DB is used.
	DatabaseDSN string `env:"DATABASE_DSN" json:"database_dsn"`
	// Key is the secret key for signing and verifying HMAC SHA-256 requests.
	Key string `env:"KEY" json:"key"`
	// AuditFile is the path to the file for writing audit events.
	AuditFile string `env:"AUDIT_FILE" json:"audit_file"`
	// AuditURL is the URL for sending audit events.
	AuditURL string `env:"AUDIT_URL" json:"audit_url"`
	// ProfileAddress is the HTTP profiling server address and port.
	ProfileAddress string `env:"PROFILE_ADDRESS" json:"profile_address"`
	// CryptoKey is the path to file with private key for decryption agent's reports
	CryptoKey string `env:"CRYPTO_KEY" json:"crypto_key"`
	// ConfigFile is the path to config json-file
	Config string `env:"CONFIG" json:"-"`
	// StoreInterval is the interval (in seconds) for saving metrics to a file.
	StoreInterval config.Duration `env:"STORE_INTERVAL" json:"store_interval"`
	// AuditQueueSize limits the size of the audit event queue.
	AuditQueueSize int `env:"AUDIT_QUEUE_SIZE" json:"audit_queue_size"`
	// Restore indicates whether to restore metrics from the file on startup.
	Restore bool `env:"RESTORE" json:"restore"`
	// ProfileEnabled indicates whether pprof profiling is enabled.
	ProfileEnabled bool `env:"PROFILE_ENABLED" json:"profile_enabled"`
}

// ReadConfig populates the configuration from command-line flags and
// environment variables (environment takes precedence).
func ReadConfig() (Config, error) {
	flagConfig := Config{
		StoreInterval: defaultStoreInterval,
	}

	flag.StringVar(&flagConfig.ServerAddress, "a", defaultServerAddress, "address and port to run server")
	flag.StringVar(&flagConfig.FileStoragePath, "f", defaultFileStoragePath, "file for saving metrics")
	flag.StringVar(&flagConfig.DatabaseDSN, "d", defaultDatabaseDSN, "database url")
	flag.StringVar(&flagConfig.Key, "k", defaultKey, "key for signing request with HMAC SHA-256")
	flag.StringVar(&flagConfig.AuditFile, "audit-file", defaultAuditFile, "file for writing audit event")
	flag.StringVar(&flagConfig.AuditURL, "audit-url", defaultAuditURL, "url for sending audit event")
	flag.StringVar(&flagConfig.ProfileAddress, "pa", defaultProfileAddress, "address and port to run profile server")
	flag.StringVar(&flagConfig.CryptoKey, "crypto-key", defaultCryptoKey, "path to file with private key for decryption agent's reports")
	flag.Var(&flagConfig.StoreInterval, "i", "interval for saving metric to file")
	flag.IntVar(&flagConfig.AuditQueueSize, "audit-queue-size", defaultAuditQueueSize, "size of the audit event queue")
	flag.BoolVar(&flagConfig.Restore, "r", defaultRestore, "restore metric from file")
	flag.BoolVar(&flagConfig.ProfileEnabled, "p", defaultProfileEnabled, "enable profiling")

	flag.StringVar(&flagConfig.Config, "c", defaultConfig, "path to file with config")
	flag.StringVar(&flagConfig.Config, "config", defaultConfig, "path to file with config")

	flag.Parse()

	setFlags := map[string]bool{}
	flag.Visit(func(f *flag.Flag) { setFlags[f.Name] = true })

	configPath := ""
	if setFlags["c"] {
		configPath = flagConfig.Config
	}
	if setFlags["config"] {
		configPath = flagConfig.Config
	}
	if v, ok := os.LookupEnv("CONFIG"); ok && v != "" {
		configPath = v
	}

	config := Config{
		ServerAddress:   defaultServerAddress,
		FileStoragePath: defaultFileStoragePath,
		DatabaseDSN:     defaultDatabaseDSN,
		Key:             defaultKey,
		AuditFile:       defaultAuditFile,
		AuditURL:        defaultAuditURL,
		ProfileAddress:  defaultProfileAddress,
		CryptoKey:       defaultCryptoKey,
		StoreInterval:   defaultStoreInterval,
		AuditQueueSize:  defaultAuditQueueSize,
		Restore:         defaultRestore,
		ProfileEnabled:  defaultProfileEnabled,
	}
	if configPath != "" {
		if err := loadJSON(configPath, &config); err != nil {
			return config, fmt.Errorf("load config file %q: %w", configPath, err)
		}
	}

	if setFlags["a"] {
		config.ServerAddress = flagConfig.ServerAddress
	}
	if setFlags["f"] {
		config.FileStoragePath = flagConfig.FileStoragePath
	}
	if setFlags["d"] {
		config.DatabaseDSN = flagConfig.DatabaseDSN
	}
	if setFlags["k"] {
		config.Key = flagConfig.Key
	}
	if setFlags["audit-file"] {
		config.AuditFile = flagConfig.AuditFile
	}
	if setFlags["audit-url"] {
		config.AuditURL = flagConfig.AuditURL
	}
	if setFlags["pa"] {
		config.ProfileAddress = flagConfig.ProfileAddress
	}
	if setFlags["crypto-key"] {
		config.CryptoKey = flagConfig.CryptoKey
	}
	if setFlags["i"] {
		config.StoreInterval = flagConfig.StoreInterval
	}
	if setFlags["audit-queue-size"] {
		config.AuditQueueSize = flagConfig.AuditQueueSize
	}
	if setFlags["r"] {
		config.Restore = flagConfig.Restore
	}
	if setFlags["p"] {
		config.ProfileEnabled = flagConfig.ProfileEnabled
	}
	config.Config = configPath

	err := env.Parse(&config)
	if err != nil {
		return config, fmt.Errorf("error parse environment variables: %w", err)
	}

	return config, nil
}

func loadJSON(path string, cfg *Config) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, cfg); err != nil {
		return err
	}
	return nil
}
