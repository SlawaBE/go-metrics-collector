package service

import (
	"errors"
	"strconv"
)

type MetricsService struct {
	storage Storage
}

type Storage interface {
	SaveCounter(key string, value int64) error
	SaveGauge(key string, value float64) error
}

func NewMetricsService(s Storage) *MetricsService {
	return &MetricsService{
		storage: s,
	}
}

func (m *MetricsService) UpdateMetric(mType, name, value string) error {
	if name == "" {
		return errors.New("empty metric name")
	}

	switch mType {
	case "counter":
		intValue, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return errors.New("invalid value")
		}
		m.storage.SaveCounter(name, intValue)

	case "gauge":
		floatValue, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return errors.New("invalid value")
		}
		m.storage.SaveGauge(name, floatValue)

	default:
		return errors.New("unknown metric type")
	}

	return nil
}
