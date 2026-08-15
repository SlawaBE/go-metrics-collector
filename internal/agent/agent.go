package agent

import (
	"context"

	"github.com/SlawaBE/go-metrics-collector/internal/agent/config"
	"github.com/SlawaBE/go-metrics-collector/internal/storage"
)

func Run(config config.Config) {
	storage := storage.NewMemStorage()
	runtimePoller := NewRuntimeMetricsPoller(storage, config.PollInterval)
	gopsutilPoller := NewGopsultilMetricsPoller(storage, config.PollInterval)
	reporter := NewReporter(storage, config.ServerAddress, config.ReportInterval, config.Key, config.RateLimit)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go runtimePoller.Run(ctx)
	go gopsutilPoller.Run(ctx)
	reporter.Run(ctx)
}
