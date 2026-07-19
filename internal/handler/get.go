package handler

import (
	"net/http"

	"github.com/SlawaBE/go-metrics-collector/internal/model"
	"github.com/SlawaBE/go-metrics-collector/internal/service"
	"github.com/SlawaBE/go-metrics-collector/internal/utils"
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

	m, err := h.service.GetMetric(model.MetricRequest{MType: metricType, ID: metricName})

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
		switch metricType {
		case model.Gauge:
			w.Write([]byte(utils.ConvertGauge(*m.Value)))
		case model.Counter:
			w.Write([]byte(utils.ConvertCounter(*m.Delta)))
		}
	}
}
