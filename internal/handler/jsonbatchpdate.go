package handler

import (
	"encoding/json"
	"net/http"

	"github.com/SlawaBE/go-metrics-collector/internal/logger"
	"github.com/SlawaBE/go-metrics-collector/internal/model"
	"github.com/SlawaBE/go-metrics-collector/internal/service"
	"go.uber.org/zap"
)

type JsonBatchUpdateMetricsHandler struct {
	service *service.MetricsService
}

func NewJsonBatchUpdateMetricsHandler(service *service.MetricsService) *JsonBatchUpdateMetricsHandler {
	return &JsonBatchUpdateMetricsHandler{
		service: service,
	}
}

func (h *JsonBatchUpdateMetricsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var metrics []model.Metric
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&metrics); err != nil {
		logger.Log.Error("cannot decode request JSON body", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	for _, m := range metrics {
		if m.ID == "" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if m.MType != model.Counter && m.MType != model.Gauge {
			w.WriteHeader(http.StatusBadRequest)
			return
		} 
	}

	err := h.service.UpdateMetrics(r.Context(), metrics)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{}"))
	}
}
