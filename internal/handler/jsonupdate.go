package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/SlawaBE/go-metrics-collector/internal/logger"
	"github.com/SlawaBE/go-metrics-collector/internal/model"
	"github.com/SlawaBE/go-metrics-collector/internal/service"
	"go.uber.org/zap"
)

type JSONUpdateMetricHandler struct {
	service        *service.MetricsService
	auditPublisher service.AuditPublisher
}

func NewJSONUpdateMetricHandler(service *service.MetricsService, auditPublisher service.AuditPublisher) *JSONUpdateMetricHandler {
	return &JSONUpdateMetricHandler{
		service:        service,
		auditPublisher: auditPublisher,
	}
}

func (h *JSONUpdateMetricHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var metric model.Metric
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&metric); err != nil {
		logger.Log.Error("cannot decode request JSON body", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if metric.ID == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err := h.service.UpdateMetric(r.Context(), metric)

	h.auditPublisher.SendMetric(r.Context(), strings.Split(r.RemoteAddr, ":")[0], metric)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{}"))
	}
}
