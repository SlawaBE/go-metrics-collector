package server

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
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
	if config.DatabaseDSN != "" {
		var err error
		database, err = db.NewDB(config.DatabaseDSN)
		if err != nil {
			os.Exit(2)
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		err = database.PingContext(ctx)
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
		storage := storage.NewDBStorage(database)
		metricsService = service.NewMetricsService(storage, nil)
	} else {
		storage := storage.NewMemStorage()
		saver := service.NewJsonFileMetricSaver(config.StoreInterval, config.FileStoragePath, storage)
		if config.Restore {
			saver.Load()
		}

		ctx, cancel := context.WithCancel(context.Background())
		saver.StartSync(ctx)
		defer cancel()
		metricsService = service.NewMetricsService(storage, saver)
	}
	//TODO разобраться с Graceful Shutdown иначе это смысла не имеет

	r := InitRouter(metricsService, database)
	gzipper := middleware.GZip(r)
	logger := middleware.RequestLogger(gzipper)

	err := http.ListenAndServe(config.ServerAddress, logger)
	if err != nil {
		log.Fatal(err)
	}
}
