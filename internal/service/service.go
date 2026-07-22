package service

import (
	"errors"
	"sort"

	"github.com/SlawaBE/go-metrics-collector/internal/model"
	"github.com/SlawaBE/go-metrics-collector/internal/storage"
	"github.com/SlawaBE/go-metrics-collector/internal/utils"
)

type MetricsService struct {
	storage storage.Storage
	saver   FileMetricSaver
}

func NewMetricsService(s storage.Storage, saver FileMetricSaver) *MetricsService {
	return &MetricsService{
		storage: s,
		saver:   saver,
	}
}

func (m *MetricsService) UpdateMetric(metric model.Metric) error {
	if metric.ID == "" {
		return errors.New("empty metric name")
	}
	err := m.storage.UpdateMetric(metric)
	if err == nil {
		m.saver.SaveSync()
	}
	return err
}

func (m *MetricsService) GetMetric(request model.MetricRequest) (*model.Metric, error) {
	metric, err := m.storage.GetMetric(request.ID)
	if err != nil || metric.MType != request.MType {
		return nil, errors.New("not found")
	}
	return metric, nil
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
