package main

import (
	"github.com/SlawaBE/go-metrics-collector/internal/server"
	"github.com/SlawaBE/go-metrics-collector/internal/server/config"
)

func main() {
	conf := config.ReadConfig()
	server.Run(conf)
}
