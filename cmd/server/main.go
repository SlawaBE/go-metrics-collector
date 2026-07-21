package main

import (
	"fmt"
	"os"

	"github.com/SlawaBE/go-metrics-collector/internal/server"
	"github.com/SlawaBE/go-metrics-collector/internal/server/config"
)

func main() {
	conf, err := config.ReadConfig()
	if err != nil {
		fmt.Println("Error reading configuration")
		os.Exit(1)
	}
	server.Run(conf)
}
