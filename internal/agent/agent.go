package agent

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/SlawaBE/go-metrics-collector/internal/agent/config"
	"github.com/SlawaBE/go-metrics-collector/internal/logger"
	"github.com/SlawaBE/go-metrics-collector/internal/storage"
)

func Run(config config.Config) {
	logger.Initialize("info")

	storage := storage.NewMemStorage()
	runtimePoller := NewRuntimeMetricsPoller(storage, config.PollInterval)
	gopsutilPoller := NewGopsutilMetricsPoller(storage, config.PollInterval)
	reporter := NewReporter(storage, config.ServerAddress, config.ReportInterval, config.Key, config.RateLimit)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
		<-signalChan

		logger.Log.Info("Shutdown signal received")
		cancel()
	}()

	var wg sync.WaitGroup
	wg.Add(3)
	go func() { defer wg.Done(); runtimePoller.Run(ctx) }()
	go func() { defer wg.Done(); gopsutilPoller.Run(ctx) }()
	go func() { defer wg.Done(); reporter.Run(ctx) }()

	wg.Wait()
	logger.Log.Info("Agent stopped")
}
