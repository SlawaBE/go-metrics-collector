package main

import "github.com/SlawaBE/go-metrics-collector/internal/agent"

func main() {
	parseFlags()
	agent.Run(flagRunAddress, flagPollInterval, flagReportInterval)
}
