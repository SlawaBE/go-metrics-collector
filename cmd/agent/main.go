package main

import (
	"fmt"
	"log"

	"github.com/SlawaBE/go-metrics-collector/internal/agent"
	"github.com/SlawaBE/go-metrics-collector/internal/agent/config"
	"github.com/SlawaBE/go-metrics-collector/internal/buildinfo"
)

var (
	buildVersion string
	buildDate    string
	buildCommit  string
)

func main() {
	buildinfo.Print(buildVersion, buildDate, buildCommit)
	if err := run(); err != nil {
		log.Fatalf("Error running agent: %v", err)
	}
}

// run reads the agent configuration and starts the agent. It returns an
// error only when the configuration cannot be loaded; the agent itself keeps
// running until a shutdown signal is received.
func run() error {
	conf, err := config.ReadConfig()
	if err != nil {
		return fmt.Errorf("error reading configuration: %w", err)
	}
	agent.Run(conf)
	return nil
}
