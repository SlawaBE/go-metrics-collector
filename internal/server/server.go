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

func InitRouter(service *service.MetricsService, database *sql.DB) chi.Router {
	r := chi.NewRouter()
	updateHandler := handler.NewUpdateMetricHandler(service)
	getHandler := handler.NewGetMetricHandler(service)
	listHandler := handler.NewListMetricHandler(service)
	jsonUpdateHandler := handler.NewJsonUpdateMetricHandler(service)
	jsonGetMetricHandler := handler.NewJsonGetMetricHandler(service)
	jsonBatchUpdatesHandler := handler.NewJsonBatchUpdateMetricsHandler(service)

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

func Run(config config.Config) {
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
		saver := service.NewJsonFileMetricSaver(config.StoreInterval, config.FileStoragePath, storageInstance)
		if config.Restore {
			saver.Load()
		}

		saver.StartSync(ctx)
		metricsService = service.NewMetricsService(storageInstance, saver)
	}

	var r http.Handler = InitRouter(metricsService, database)
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
