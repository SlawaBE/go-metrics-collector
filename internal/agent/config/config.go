// Package config holds the metrics collection agent configuration, populated
// from command-line flags and environment variables.
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
	defaultGRPCServerAddress = ""
	defaultPollInterval      = config.Duration(time.Second * 2)
	defaultReportInterval    = config.Duration(time.Second * 10)
	defaultKey               = ""
	defaultRateLimit         = 1
	defaultCryptoKey         = ""
	defaultConfig            = ""
	defaultGRPCCertFile      = ""
	defaultGRPCKeyFile       = ""
	defaultGRPCCAFile        = ""
)

// Config describes the metrics collection agent startup parameters.
// generate:reset
type Config struct {
	// ServerAddress is the metrics server address and port.
	ServerAddress string `env:"ADDRESS" json:"address"`
	// GRPCServerAddress is the GRPC server address and port
	GRPCServerAddress string `env:"GRPC_ADDRESS" json:"grpc_address"`
	// Key is the secret key for signing requests with HMAC SHA-256.
	Key string `env:"KEY" json:"key"`
	// CryptoKey is the path to file with public key for encryption reports
	CryptoKey string `env:"CRYPTO_KEY" json:"crypto_key"`
	// ConfigFile is the path to config json-file
	Config string `env:"CONFIG" json:"-"`
	// TLSCertFile is the path to the TLS certificate file for gRPC client.
	TLSCertFile string `env:"GRPC_TLS_CERT_FILE" json:"grpc_tls_cert_file"`
	// TLSKeyFile is the path to the TLS key file for gRPC client.
	TLSKeyFile string `env:"GRPC_TLS_KEY_FILE" json:"grpc_tls_key_file"`
	// TLSCAFile is the path to the CA certificate file for server verification.
	TLSCAFile string `env:"GRPC_TLS_CA_FILE" json:"grpc_tls_ca_file"`
	// PollInterval is the metrics polling interval (in seconds).
	PollInterval config.Duration `env:"POLL_INTERVAL" json:"poll_interval"`
	// ReportInterval is the interval (in seconds) for sending reports to the server.
	ReportInterval config.Duration `env:"REPORT_INTERVAL" json:"report_interval"`
	// RateLimit is the limit of workers sending requests to the server.
	RateLimit int `env:"RATE_LIMIT" json:"rate_limit"`
}

// ReadConfig populates the configuration from command-line flags and
// environment variables (environment takes precedence).
func ReadConfig() (Config, error) {
	flagConfig := Config{
		PollInterval:   defaultPollInterval,
		ReportInterval: defaultReportInterval,
	}

	flag.StringVar(&flagConfig.ServerAddress, "a", defaultServerAddress, "address and port of metric server")
	flag.StringVar(&flagConfig.GRPCServerAddress, "g", defaultGRPCServerAddress, "address and port to run GRPC server")
	flag.Var(&flagConfig.PollInterval, "p", "metrics polling interval")
	flag.Var(&flagConfig.ReportInterval, "r", "metrics report interval")
	flag.StringVar(&flagConfig.Key, "k", defaultKey, "key for signing request with HMAC SHA-256")
	flag.IntVar(&flagConfig.RateLimit, "l", defaultRateLimit, "rate limit for sending requests to server")
	flag.StringVar(&flagConfig.CryptoKey, "crypto-key", defaultCryptoKey, "path to file with public key for encryption reports")
	flag.StringVar(&flagConfig.TLSCertFile, "grpc-tls-cert", defaultGRPCCertFile, "path to TLS certificate file for gRPC client")
	flag.StringVar(&flagConfig.TLSKeyFile, "grpc-tls-key", defaultGRPCKeyFile, "path to TLS key file for gRPC client")
	flag.StringVar(&flagConfig.TLSCAFile, "grpc-tls-ca", defaultGRPCCAFile, "path to CA certificate file for gRPC server verification")

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
		Key:               defaultKey,
		CryptoKey:         defaultCryptoKey,
		PollInterval:      defaultPollInterval,
		ReportInterval:    defaultReportInterval,
		RateLimit:         defaultRateLimit,
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
	if setFlags["k"] {
		config.Key = flagConfig.Key
	}
	if setFlags["crypto-key"] {
		config.CryptoKey = flagConfig.CryptoKey
	}
	if setFlags["p"] {
		config.PollInterval = flagConfig.PollInterval
	}
	if setFlags["r"] {
		config.ReportInterval = flagConfig.ReportInterval
	}
	if setFlags["l"] {
		config.RateLimit = flagConfig.RateLimit
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
