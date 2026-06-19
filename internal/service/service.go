package service

import (
	"errors"
	"strconv"

	"github.com/SlawaBE/go-metrics-collector/internal/model"
	"github.com/SlawaBE/go-metrics-collector/internal/storage"
)

type MetricsService struct {
	storage storage.Storage
}

func NewMetricsService(s storage.Storage) *MetricsService {
	return &MetricsService{
		storage: s,
	}
}

func (m *MetricsService) UpdateMetric(mType, name, value string) error {
	if name == "" {
		return errors.New("empty metric name")
	}

	var metric model.Metric

	switch mType {
	case model.Counter:
		intValue, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return errors.New("invalid value")
		}
		metric = model.NewCounterMetric(name, intValue)

	case model.Gauge:
		floatValue, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return errors.New("invalid value")
		}
		metric = model.NewGaugeMetric(name, floatValue)

	default:
		return errors.New("unknown metric type")
	}

	m.storage.UpdateMetric(metric)
	return nil
}
