package main

import (
	"fmt"
	"log"

	"github.com/SlawaBE/go-metrics-collector/internal/buildinfo"
	"github.com/SlawaBE/go-metrics-collector/internal/server"
	"github.com/SlawaBE/go-metrics-collector/internal/server/config"
)

var (
	buildVersion string
	buildDate    string
	buildCommit  string
)

func main() {
	buildinfo.Print(buildVersion, buildDate, buildCommit)
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
