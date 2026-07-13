package main

import (
	"github.com/SlawaBE/go-metrics-collector/internal/agent"
	"github.com/SlawaBE/go-metrics-collector/internal/agent/config"
)

func main() {
	conf := config.ReadConfig()
	agent.Run(conf)
}
