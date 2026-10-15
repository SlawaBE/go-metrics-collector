package handler

import (
	"net/http"

	"github.com/SlawaBE/go-metrics-collector/internal/logger"
	"github.com/SlawaBE/go-metrics-collector/internal/model"
	"github.com/SlawaBE/go-metrics-collector/internal/service"
	"github.com/SlawaBE/go-metrics-collector/internal/utils"
)

// GetMetricHandler handles fetching a metric via the URL path:
// GET /value/{type}/{name}.
type GetMetricHandler struct {
	service *service.MetricsService
}

// NewGetMetricHandler creates a handler for fetching a metric via the URL path.
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

	m, err := h.service.GetMetric(r.Context(), model.MetricRequest{MType: metricType, ID: metricName})

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
		switch metricType {
		case model.Gauge:
			h.writeResponseBody(w, utils.ConvertGauge(*m.Value))
		case model.Counter:
			h.writeResponseBody(w, utils.ConvertCounter(*m.Delta))
		}
	}
}

func (h *GetMetricHandler) writeResponseBody(w http.ResponseWriter, body string) {
	if _, err := w.Write([]byte(body)); err != nil {
		logger.Log.Error("error write response body")
	}
}
