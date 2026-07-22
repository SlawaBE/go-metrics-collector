package config

import (
	"flag"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	ServerAddress   string `env:"ADDRESS"`
	StoreInterval   int    `env:"STORE_INTERVAL"`
	FileStoragePath string `env:"FILE_STORAGE_PATH"`
	Restore         bool   `env:"RESTORE"`
}

func ReadConfig() (Config, error) {
	var config Config

	flag.StringVar(&config.ServerAddress, "a", "localhost:8080", "address and port to run server")
	flag.IntVar(&config.StoreInterval, "i", 300, "interval for saving metric to file")
	flag.StringVar(&config.FileStoragePath, "f", "storage.json", "file for saving metrics")
	flag.BoolVar(&config.Restore, "r", false, "restore metric from file")
	flag.Parse()

	err := env.Parse(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}
