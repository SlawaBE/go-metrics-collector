package service

import (
	"errors"
	"sort"
	"strconv"

	"github.com/SlawaBE/go-metrics-collector/internal/model"
	"github.com/SlawaBE/go-metrics-collector/internal/storage"
	"github.com/SlawaBE/go-metrics-collector/internal/utils"
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

func (m *MetricsService) GetMetric(mType, name string) (string, error) {
	if name == "" {
		return "", errors.New("empty metric name")
	}

	metric, err := m.storage.GetMetric(name)
	if err != nil || metric.MType != mType {
		return "", errors.New("not found")
	}

	switch mType {
	case model.Counter:
		return utils.ConvertCounter(*metric.Delta), nil
	case model.Gauge:
		return utils.ConvertGauge(*metric.Value), nil
	default:
		return "", errors.New("unknown metric type")
	}
}

func (m *MetricsService) List() []string {
	list := m.storage.GetValues()
	sort.Slice(list, func(i, j int) bool {
		return list[i].ID < list[j].ID
	})
	res := make([]string, len(list))

	for i, v := range list {
		switch v.MType {
		case model.Counter:
			res[i] = v.ID + ": " + utils.ConvertCounter(*v.Delta)
		case model.Gauge:
			res[i] = v.ID + ": " + utils.ConvertGauge(*v.Value)
		}

	}

	return res
}
