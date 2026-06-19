package main

import (
	"github.com/SlawaBE/go-metrics-collector/internal/server"
)

func main() {
	parseFlags()
	server.Run(flagRunAddress)
}
