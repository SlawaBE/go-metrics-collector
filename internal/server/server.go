// Package server is responsible for initializing the HTTP router and starting
// the metrics collection server.
package server

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "net/http/pprof"

	"github.com/SlawaBE/go-metrics-collector/internal/db"
	"github.com/SlawaBE/go-metrics-collector/internal/handler"
	"github.com/SlawaBE/go-metrics-collector/internal/logger"
	"github.com/SlawaBE/go-metrics-collector/internal/middleware"
	"github.com/SlawaBE/go-metrics-collector/internal/server/config"
	"github.com/SlawaBE/go-metrics-collector/internal/service"
	"github.com/SlawaBE/go-metrics-collector/internal/storage"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

// InitRouter builds the chi router with all HTTP endpoints of the server.
//
// Available endpoints:
//   - GET  /                                   - list of all metrics (HTML);
//   - POST /update/{type}/{name}/{value}        - update a metric via URL;
//   - GET  /value/{type}/{name}                 - fetch a metric via URL;
//   - POST /update                              - JSON update of a metric;
//   - POST /value                               - JSON fetch of a metric;
//   - POST /updates                             - JSON batch update;
//   - GET  /ping                                - database connectivity check.
func InitRouter(service *service.MetricsService, database *sql.DB, auditPublisher service.AuditPublisher) chi.Router {
	r := chi.NewRouter()
	updateHandler := handler.NewUpdateMetricHandler(service, auditPublisher)
	getHandler := handler.NewGetMetricHandler(service)
	listHandler := handler.NewListMetricHandler(service)
	jsonUpdateHandler := handler.NewJSONUpdateMetricHandler(service, auditPublisher)
	jsonGetMetricHandler := handler.NewJSONGetMetricHandler(service)
	jsonBatchUpdatesHandler := handler.NewJSONBatchUpdateMetricsHandler(service, auditPublisher)

	r.Handle("GET /", listHandler)
	r.Handle("POST /update/{type}/{name}/{value}", updateHandler)
	r.Handle("GET /value/{type}/{name}", getHandler)

	r.Handle("POST /update", jsonUpdateHandler)
	r.Handle("POST /update/", jsonUpdateHandler)
	r.Handle("POST /value", jsonGetMetricHandler)
	r.Handle("POST /value/", jsonGetMetricHandler)

	r.Handle("GET /ping", handler.NewPingHandler(database))

	r.Handle("POST /updates", jsonBatchUpdatesHandler)
	r.Handle("POST /updates/", jsonBatchUpdatesHandler)

	return r
}

// Run starts the metrics server: initializes the storage (memory or database),
// sets up auditing, attaches the middleware (signature, gzip, logging) and
// listens for incoming requests until a shutdown signal is received.
func Run(config config.Config) {
	if config.ProfileEnabled {
		go func() {
			if err := http.ListenAndServe(":8085", nil); err != nil {
				os.Exit(5)
			}
		}()
	}
	logger.Initialize("info")

	var database *sql.DB
	var metricsService *service.MetricsService

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if config.DatabaseDSN != "" {
		var err error
		database, err = db.NewDB(config.DatabaseDSN)
		if err != nil {
			os.Exit(2)
		}

		pingCtx, pingCancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer pingCancel()
		err = database.PingContext(pingCtx)
		if err != nil {
			logger.Log.Fatal("Error ping database", zap.Error(err))
			os.Exit(3)
		}

		err = db.RunMigrations(database, config.DatabaseDSN)
		if err != nil {
			logger.Log.Fatal("Error migration", zap.Error(err))
			os.Exit(4)
		}
		defer database.Close()
		storageInstance := storage.NewDBStorage(database)
		metricsService = service.NewMetricsService(storageInstance, nil)
	} else {
		storageInstance := storage.NewMemStorage()
		saver := service.NewJSONFileMetricSaver(config.StoreInterval, config.FileStoragePath, storageInstance)
		if config.Restore {
			saver.Load()
		}

		saver.StartSync(ctx)
		metricsService = service.NewMetricsService(storageInstance, saver)
	}

	auditService := service.NewAuditService()
	if config.AuditFile != "" {
		fileAuditSubscriber := service.NewFileAuditSubscriber(config.AuditFile)
		auditService.Subscribe(fileAuditSubscriber)
	}
	if config.AuditURL != "" {
		httpAuditSubscriber := service.NewHTTPAuditSubscriber(config.AuditURL)
		auditService.Subscribe(httpAuditSubscriber)
	}

	var r http.Handler = InitRouter(metricsService, database, auditService)
	if config.Key != "" {
		r = middleware.NewCheckSum(config.Key).CheckSumMiddleware(r)
	}
	gzipper := middleware.GZip(r)
	requestLogger := middleware.RequestLogger(gzipper)

	httpServer := &http.Server{
		Addr:    config.ServerAddress,
		Handler: requestLogger,
	}

	shutdownDone := make(chan struct{})
	go func() {
		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
		<-signalChan

		logger.Log.Info("Shutdown signal received")
		cancel()

		shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer shutdownCancel()
		err := httpServer.Shutdown(shutdownCtx)
		if err != nil {
			logger.Log.Error("Error during server shutdown", zap.Error(err))
		}
		close(shutdownDone)
	}()

	err := httpServer.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}

	<-shutdownDone
	logger.Log.Info("Server stopped")
}
