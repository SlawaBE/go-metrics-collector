package handler

import (
	"encoding/json"
	"net/http"

	"github.com/SlawaBE/go-metrics-collector/internal/logger"
	"github.com/SlawaBE/go-metrics-collector/internal/model"
	"github.com/SlawaBE/go-metrics-collector/internal/service"
	"go.uber.org/zap"
)

type JsonGetMetricHandler struct {
	service *service.MetricsService
}

func NewJsonGetMetricHandler(service *service.MetricsService) *JsonGetMetricHandler {
	return &JsonGetMetricHandler{
		service: service,
	}
}

func (h *JsonGetMetricHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var request model.MetricRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		logger.Log.Error("cannot decode request JSON body", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if request.ID == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	m, err := h.service.GetMetric(request)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
		enc := json.NewEncoder(w)
		if err := enc.Encode(m); err != nil {
			logger.Log.Error("error encoding response", zap.Error(err))
			return
		}
	}
}
