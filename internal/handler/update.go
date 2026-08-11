package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/SlawaBE/go-metrics-collector/internal/model"
	"github.com/SlawaBE/go-metrics-collector/internal/service"
	"github.com/SlawaBE/go-metrics-collector/internal/utils"
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

	metric, err := getMetric(metricType, metricName, metricValue)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.service.UpdateMetric(r.Context(), *metric)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func getMetric(mType, name, value string) (*model.Metric, error) {
	switch mType {
	case model.Counter:
		intValue, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return nil, errors.New("invalid value")
		}
		return utils.Ptr(model.NewCounterMetric(name, intValue)), nil

	case model.Gauge:
		floatValue, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return nil, errors.New("invalid value")
		}
		return utils.Ptr(model.NewGaugeMetric(name, floatValue)), nil

	default:
		return nil, errors.New("unknown metric type")
	}
}
