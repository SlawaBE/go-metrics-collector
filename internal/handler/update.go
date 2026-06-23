package handler

import (
	"net/http"

	"github.com/SlawaBE/go-metrics-collector/internal/service"
)

type UpdateMetricHandler struct {
	service *service.MetricsService
}

func NewUpdateMetricHandler(service *service.MetricsService) *UpdateMetricHandler {
	return &UpdateMetricHandler{
		service: service,
	}
}

func (h *UpdateMetricHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")

	if r.Method != http.MethodPost {
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
	}

	metricType := r.PathValue("type")
	metricName := r.PathValue("name")
	metricValue := r.PathValue("value")

	if metricName == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err := h.service.UpdateMetric(metricType, metricName, metricValue)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}
