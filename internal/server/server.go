package server

import (
	"log"
	"net/http"

	"github.com/SlawaBE/go-metrics-collector/internal/handler"
	"github.com/SlawaBE/go-metrics-collector/internal/service"
	"github.com/SlawaBE/go-metrics-collector/internal/storage"
)

func Run() {
	storage := storage.NewMemStorage()
	service := service.NewMetricsService(storage)

	mux := http.NewServeMux()
	mux.Handle("POST /update/{type}/{name}/{value}", handler.NewUpdateMetricHandler(service))

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
