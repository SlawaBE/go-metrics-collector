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
	defaultServerAddress     = "localhost:8080"
	defaultGRPCServerAddress = "localhost:9090"
	defaultFileStoragePath   = "storage.json"
	defaultDatabaseDSN       = ""
	defaultKey               = ""
	defaultAuditFile         = ""
	defaultAuditURL          = ""
	defaultProfileAddress    = "localhost:8085"
	defaultCryptoKey         = ""
	defaultStoreInterval     = config.Duration(time.Second * 300)
	defaultAuditQueueSize    = 100
	defaultRestore           = false
	defaultProfileEnabled    = false
	defaultConfig            = ""
	defaultTrustedSubnet     = ""
	defaultGRPCCertFile      = ""
	defaultGRPCKeyFile       = ""
	defaultGRPCCAFile        = ""
)

// Config describes the metrics server startup parameters.
// generate:reset
type Config struct {
	// ServerAddress is the HTTP server address and port.
	ServerAddress string `env:"ADDRESS" json:"address"`
	// GRPCServerAddress is the GRPC server address and port
	GRPCServerAddress string `env:"GRPC_ADDRESS" json:"grpc_address"`
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
	// TrustedSubnet is a CIDR string; when set, agent requests must come from
	// an IP within this subnet (checked via X-Real-IP header).
	TrustedSubnet string `env:"TRUSTED_SUBNET" json:"trusted_subnet"`
	// TLSCertFile is the path to the TLS certificate file for the gRPC server.
	TLSCertFile string `env:"GRPC_TLS_CERT_FILE" json:"grpc_tls_cert_file"`
	// TLSKeyFile is the path to the TLS key file for the gRPC server.
	TLSKeyFile string `env:"GRPC_TLS_KEY_FILE" json:"grpc_tls_key_file"`
	// TLSCAFile is the path to the CA certificate file for client verification (mTLS).
	TLSCAFile string `env:"GRPC_TLS_CA_FILE" json:"grpc_tls_ca_file"`
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

	flag.StringVar(&flagConfig.ServerAddress, "a", defaultServerAddress, "address and port to run HTTP server")
	flag.StringVar(&flagConfig.GRPCServerAddress, "g", defaultGRPCServerAddress, "address and port to run GRPC server")
	flag.StringVar(&flagConfig.FileStoragePath, "f", defaultFileStoragePath, "file for saving metrics")
	flag.StringVar(&flagConfig.DatabaseDSN, "d", defaultDatabaseDSN, "database url")
	flag.StringVar(&flagConfig.Key, "k", defaultKey, "key for signing request with HMAC SHA-256")
	flag.StringVar(&flagConfig.AuditFile, "audit-file", defaultAuditFile, "file for writing audit event")
	flag.StringVar(&flagConfig.AuditURL, "audit-url", defaultAuditURL, "url for sending audit event")
	flag.StringVar(&flagConfig.ProfileAddress, "pa", defaultProfileAddress, "address and port to run profile server")
	flag.StringVar(&flagConfig.CryptoKey, "crypto-key", defaultCryptoKey, "path to file with private key for decryption agent's reports")
	flag.StringVar(&flagConfig.TrustedSubnet, "t", defaultTrustedSubnet, "CIDR of trusted agent subnets")
	flag.Var(&flagConfig.StoreInterval, "i", "interval for saving metric to file")
	flag.IntVar(&flagConfig.AuditQueueSize, "audit-queue-size", defaultAuditQueueSize, "size of the audit event queue")
	flag.BoolVar(&flagConfig.Restore, "r", defaultRestore, "restore metric from file")
	flag.BoolVar(&flagConfig.ProfileEnabled, "p", defaultProfileEnabled, "enable profiling")
	flag.StringVar(&flagConfig.TLSCertFile, "grpc-tls-cert", defaultGRPCCertFile, "path to TLS certificate file for gRPC server")
	flag.StringVar(&flagConfig.TLSKeyFile, "grpc-tls-key", defaultGRPCKeyFile, "path to TLS key file for gRPC server")
	flag.StringVar(&flagConfig.TLSCAFile, "grpc-tls-ca", defaultGRPCCAFile, "path to CA certificate file for gRPC client verification")

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
		ServerAddress:     defaultServerAddress,
		GRPCServerAddress: defaultGRPCServerAddress,
		FileStoragePath:   defaultFileStoragePath,
		DatabaseDSN:       defaultDatabaseDSN,
		Key:               defaultKey,
		AuditFile:         defaultAuditFile,
		AuditURL:          defaultAuditURL,
		ProfileAddress:    defaultProfileAddress,
		CryptoKey:         defaultCryptoKey,
		StoreInterval:     defaultStoreInterval,
		AuditQueueSize:    defaultAuditQueueSize,
		Restore:           defaultRestore,
		ProfileEnabled:    defaultProfileEnabled,
		TrustedSubnet:     defaultTrustedSubnet,
		TLSCertFile:       defaultGRPCCertFile,
		TLSKeyFile:        defaultGRPCKeyFile,
		TLSCAFile:         defaultGRPCCAFile,
	}
	if configPath != "" {
		if err := loadJSON(configPath, &config); err != nil {
			return config, fmt.Errorf("load config file %q: %w", configPath, err)
		}
	}

	if setFlags["a"] {
		config.ServerAddress = flagConfig.ServerAddress
	}
	if setFlags["g"] {
		config.GRPCServerAddress = flagConfig.GRPCServerAddress
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
	if setFlags["t"] {
		config.TrustedSubnet = flagConfig.TrustedSubnet
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
	if setFlags["grpc-tls-cert"] {
		config.TLSCertFile = flagConfig.TLSCertFile
	}
	if setFlags["grpc-tls-key"] {
		config.TLSKeyFile = flagConfig.TLSKeyFile
	}
	if setFlags["grpc-tls-ca"] {
		config.TLSCAFile = flagConfig.TLSCAFile
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
