package handler

import (
	"net/http"

	"github.com/SlawaBE/go-metrics-collector/internal/service"
)

type GetMetricHandler struct {
	service *service.MetricsService
}

func NewGetMetricHandler(service *service.MetricsService) *GetMetricHandler {
	return &GetMetricHandler{
		service: service,
	}
}

func (h *GetMetricHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")

	if r.Method != http.MethodGet {
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
	}

	metricType := r.PathValue("type")
	metricName := r.PathValue("name")

	if metricName == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	m, err := h.service.GetMetric(metricType, metricName)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(m))
	}
}
