package server

import (
	"context"
	"log"
	"net/http"

	"github.com/SlawaBE/go-metrics-collector/internal/handler"
	"github.com/SlawaBE/go-metrics-collector/internal/logger"
	"github.com/SlawaBE/go-metrics-collector/internal/middleware"
	"github.com/SlawaBE/go-metrics-collector/internal/server/config"
	"github.com/SlawaBE/go-metrics-collector/internal/service"
	"github.com/SlawaBE/go-metrics-collector/internal/storage"
	"github.com/go-chi/chi/v5"
)

func InitRouter(service *service.MetricsService) chi.Router {
	r := chi.NewRouter()
	updateHandler := handler.NewUpdateMetricHandler(service)
	getHandler := handler.NewGetMetricHandler(service)
	listHandler := handler.NewListMetricHandler(service)
	jsonUpdateHandler := handler.NewJsonUpdateMetricHandler(service)
	jsonGetMetricHandler := handler.NewJsonGetMetricHandler(service)

	r.Handle("GET /", listHandler)
	r.Handle("POST /update/{type}/{name}/{value}", updateHandler)
	r.Handle("GET /value/{type}/{name}", getHandler)

	r.Handle("POST /update", jsonUpdateHandler)
	r.Handle("POST /update/", jsonUpdateHandler)
	r.Handle("POST /value", jsonGetMetricHandler)
	r.Handle("POST /value/", jsonGetMetricHandler)

	return r
}

func Run(config config.Config) {
	logger.Initialize("info")
	storage := storage.NewMemStorage()
	saver := service.NewJsonFileMetricSaver(config.StoreInterval, config.FileStoragePath, storage)
	if config.Restore {
		saver.Load()
	}

	ctx, cancel := context.WithCancel(context.Background())
	saver.StartSync(ctx)
	defer cancel()
	//TODO разобраться с Graceful Shutdown иначе это смысла не имеет

	service := service.NewMetricsService(storage, saver)

	r := InitRouter(service)
	gzipper := middleware.GZip(r)
	logger := middleware.RequestLogger(gzipper)

	err := http.ListenAndServe(config.ServerAddress, logger)
	if err != nil {
		log.Fatal(err)
	}
}
