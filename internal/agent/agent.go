package agent

import (
	"github.com/SlawaBE/go-metrics-collector/internal/agent/config"
	"github.com/SlawaBE/go-metrics-collector/internal/storage"
)

func Run(config config.Config) {
	storage := storage.NewMemStorage()
	poller := NewPoller(storage, config.PollInterval)
	reporter := NewReporter(storage, config.ServerAddress, config.ReportInterval)

	go poller.Run()
	reporter.Run()
}
