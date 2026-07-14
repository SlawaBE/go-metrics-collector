package server

import (
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

func InitRouter() chi.Router {
	storage := storage.NewMemStorage()
	service := service.NewMetricsService(storage)

	r := chi.NewRouter()
	updateHandler := handler.NewUpdateMetricHandler(service)
	getHandler := handler.NewGetMetricHandler(service)
	listHandler := handler.NewListMetricHandler(service)

	r.Handle("GET /", listHandler)
	r.Handle("POST /update/{type}/{name}/{value}", updateHandler)
	r.Handle("GET /value/{type}/{name}", getHandler)
	return r
}

func Run(config config.Config) {
	logger.Initialize("info")
	r := InitRouter()
	m := middleware.RequestLogger(r)

	err := http.ListenAndServe(config.ServerAddress, m)
	if err != nil {
		log.Fatal(err)
	}
}
