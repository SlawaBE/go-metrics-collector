package config

import (
	"flag"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	ServerAddress  string `env:"ADDRESS"`
	PollInterval   int    `env:"POLL_INTERVAL"`
	ReportInterval int    `env:"REPORT_INTERVAL"`
}

func ReadConfig() (Config, error) {
	var config Config

	flag.StringVar(&config.ServerAddress, "a", "localhost:8080", "address and port of metric server")
	flag.IntVar(&config.PollInterval, "p", 2, "metrics polling interval")
	flag.IntVar(&config.ReportInterval, "r", 10, "metrics report interval")

	flag.Parse()

	err := env.Parse(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}
