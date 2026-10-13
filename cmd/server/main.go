package main

import (
	"fmt"
	"log"

	"github.com/SlawaBE/go-metrics-collector/internal/server"
	"github.com/SlawaBE/go-metrics-collector/internal/server/config"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("Error running server: %v", err)
	}
}

func run() error {
	conf, err := config.ReadConfig()
	if err != nil {
		return fmt.Errorf("error reading configuration: %w", err)
	}
	server.Run(conf)
	return nil
}
