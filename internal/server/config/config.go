package config

import (
	"flag"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	ServerAddress string `env:"ADDRESS"`
}

func ReadConfig() Config {
	var config Config

	flag.StringVar(&config.ServerAddress, "a", "localhost:8080", "address and port to run server")
	flag.Parse()

	err := env.Parse(&config)
	if err != nil {
		panic(err)
	}

	return config
}
