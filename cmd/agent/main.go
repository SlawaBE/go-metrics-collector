package main

import (
	"fmt"
	"os"

	"github.com/SlawaBE/go-metrics-collector/internal/agent"
	"github.com/SlawaBE/go-metrics-collector/internal/agent/config"
)

func main() {
	conf, err := config.ReadConfig()
	if err != nil {
		fmt.Println("Error reading configuration")
		os.Exit(1)
	}
	agent.Run(conf)
}
