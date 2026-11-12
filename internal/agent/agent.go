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
	pb "github.com/SlawaBE/go-metrics-collector/internal/proto/metrics"
	"github.com/SlawaBE/go-metrics-collector/internal/storage"
	"github.com/SlawaBE/go-metrics-collector/internal/tls"
	"github.com/SlawaBE/go-metrics-collector/internal/utils"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

	var reporter Reporter
	if config.GRPCServerAddress == "" {
		clientIP, ipErr := utils.LocalIP(config.ServerAddress)
		if ipErr != nil {
			logger.Log.Warn("failed to determine local IP, X-Real-IP header will not be set", zap.Error(ipErr))
		}
		reporter = NewHTTPReporter(storage, rsaEncryptor, config.ServerAddress, config.ReportInterval.Duration(), config.Key, config.RateLimit, clientIP)
	} else {
		clientIP, ipErr := utils.LocalIP(config.GRPCServerAddress)
		if ipErr != nil {
			logger.Log.Warn("failed to determine local IP, X-Real-IP header will not be set", zap.Error(ipErr))
		}

		var grpcDialOpts []grpc.DialOption
		if config.TLSCertFile != "" || config.TLSKeyFile != "" || config.TLSCAFile != "" {
			creds, err := tls.LoadClientCredentials(config.TLSCertFile, config.TLSKeyFile, config.TLSCAFile)
			if err != nil {
				return fmt.Errorf("error loading TLS credentials: %w", err)
			}
			grpcDialOpts = append(grpcDialOpts, grpc.WithTransportCredentials(creds))
		} else {
			grpcDialOpts = append(grpcDialOpts, grpc.WithTransportCredentials(insecure.NewCredentials()))
		}

		conn, err := grpc.NewClient(config.GRPCServerAddress, grpcDialOpts...)
		if err != nil {
			return fmt.Errorf("error open grpc connect: %w", err)
		}
		defer func() {
			if closeErr := conn.Close(); closeErr != nil {
				logger.Log.Error("failed to close dial connection", zap.Error(closeErr))
			}
		}()

		metricsClient := pb.NewMetricsClient(conn)
		reporter = NewGRPCReporter(storage, clientIP, metricsClient, config.ReportInterval.Duration(), config.RateLimit)
	}

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
