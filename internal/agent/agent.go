package agent

import (
	"context"
	"fmt"
	"os/signal"
	"sync"
	"syscall"

	"github.com/SlawaBE/go-metrics-collector/internal/agent/config"
	"github.com/SlawaBE/go-metrics-collector/internal/asymcrypto"
	"github.com/SlawaBE/go-metrics-collector/internal/logger"
	"github.com/SlawaBE/go-metrics-collector/internal/storage"
	"github.com/SlawaBE/go-metrics-collector/internal/utils"
	"go.uber.org/zap"
)

// Run starts the agent: launches the runtime poller, the gopsutil poller and
// the reporter, which run until a shutdown signal is received.
func Run(config config.Config) error {
	err := logger.Initialize("info")
	if err != nil {
		return fmt.Errorf("error initializing logger: %w", err)
	}

	storage := storage.NewMemStorage()
	runtimePoller := NewRuntimeMetricsPoller(storage, config.PollInterval.Duration())
	gopsutilPoller := NewGopsutilMetricsPoller(storage, config.PollInterval.Duration())
	var rsaEncryptor *asymcrypto.RSAEncryptor
	if config.CryptoKey != "" {
		rsaEncryptor, err = asymcrypto.NewRSAEncryptor(config.CryptoKey)
		if err != nil {
			return fmt.Errorf("error initializing rsaEncryptor: %w", err)
		}
	}
	clientIP, ipErr := utils.LocalIPv4(config.ServerAddress)
	if ipErr != nil {
		logger.Log.Warn("failed to determine local IP, X-Real-IP header will not be set", zap.Error(ipErr))
	}
	reporter := NewReporter(storage, rsaEncryptor, config.ServerAddress, config.ReportInterval.Duration(), config.Key, config.RateLimit, clientIP)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	defer stop()

	var wg sync.WaitGroup
	wg.Add(3)
	go func() { defer wg.Done(); runtimePoller.Run(ctx) }()
	go func() { defer wg.Done(); gopsutilPoller.Run(ctx) }()
	go func() { defer wg.Done(); reporter.Run(ctx) }()

	<-ctx.Done()
	logger.Log.Info("Shutdown signal received")

	wg.Wait()
	logger.Log.Info("Agent stopped")
	return nil
}
