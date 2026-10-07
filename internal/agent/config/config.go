// Package config holds the metrics collection agent configuration, populated
// from command-line flags and environment variables.
package config

import (
	"flag"

	"github.com/caarlos0/env/v11"
)

// Config describes the metrics collection agent startup parameters.
type Config struct {
	// ServerAddress is the metrics server address and port.
	ServerAddress string `env:"ADDRESS"`
	// PollInterval is the metrics polling interval (in seconds).
	PollInterval int `env:"POLL_INTERVAL"`
	// ReportInterval is the interval (in seconds) for sending reports to the server.
	ReportInterval int `env:"REPORT_INTERVAL"`
	// Key is the secret key for signing requests with HMAC SHA-256.
	Key string `env:"KEY"`
	// RateLimit is the limit of workers sending requests to the server.
	RateLimit int `env:"RATE_LIMIT"`
}

// ReadConfig populates the configuration from command-line flags and
// environment variables (environment takes precedence).
func ReadConfig() (Config, error) {
	var config Config

	flag.StringVar(&config.ServerAddress, "a", "localhost:8080", "address and port of metric server")
	flag.IntVar(&config.PollInterval, "p", 2, "metrics polling interval")
	flag.IntVar(&config.ReportInterval, "r", 10, "metrics report interval")
	flag.StringVar(&config.Key, "k", "", "key for signing request with HMAC SHA-256")
	flag.IntVar(&config.RateLimit, "l", 1, "rate limit for sending requests to server")

	flag.Parse()

	err := env.Parse(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}
