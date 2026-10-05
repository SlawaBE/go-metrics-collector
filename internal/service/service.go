package service

import (
	"context"
	"errors"
	"sort"

	"github.com/SlawaBE/go-metrics-collector/internal/model"
	"github.com/SlawaBE/go-metrics-collector/internal/storage"
	"github.com/SlawaBE/go-metrics-collector/internal/utils"
)

// MetricsService encapsulates the business logic for working with metrics
// on top of a Storage implementation.
type MetricsService struct {
	storage storage.Storage
	saver   FileMetricSaver
}

// NewMetricsService creates a metrics service backed by the given storage and
// an optional file saver.
func NewMetricsService(s storage.Storage, saver FileMetricSaver) *MetricsService {
	return &MetricsService{
		storage: s,
		saver:   saver,
	}
}

// UpdateMetric updates a single metric and, when a saver is set, synchronously
// persists the state to a file.
func (m *MetricsService) UpdateMetric(ctx context.Context, metric model.Metric) error {
	if metric.ID == "" {
		return errors.New("empty metric name")
	}
	err := m.storage.UpdateMetric(ctx, metric)
	if err == nil && m.saver != nil {
		m.saver.SaveSync()
	}
	return err
}

// UpdateMetrics updates a set of metrics and, when a saver is set, synchronously
// persists the state to a file.
func (m *MetricsService) UpdateMetrics(ctx context.Context, metrics []model.Metric) error {
	err := m.storage.UpdateAll(ctx, metrics)
	if err == nil && m.saver != nil {
		m.saver.SaveSync()
	}
	return err
}

// GetMetric returns the metric matching the request; if the requested metric
// type does not match the stored one, an error "not found" is returned.
func (m *MetricsService) GetMetric(ctx context.Context, request model.MetricRequest) (*model.Metric, error) {
	metric, err := m.storage.GetMetric(ctx, request.ID)
	if err != nil || metric.MType != request.MType {
		return nil, errors.New("not found")
	}
	return metric, nil
}

// List returns a sorted list of string representations of all metrics.
func (m *MetricsService) List(ctx context.Context) ([]string, error) {
	list, err := m.storage.GetValues(ctx)
	if err != nil {
		return nil, err
	}
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

	return res, nil
}
