package storage

import (
	"context"
	"errors"
	"sync"

	"github.com/SlawaBE/go-metrics-collector/internal/model"
)

// MemStorage is a thread-safe in-memory metric storage based on a map.
// Counter accumulates via +=, Gauge is overwritten.
type MemStorage struct {
	mutex   sync.RWMutex
	metrics map[string]model.Metric
}

// NewMemStorage creates an empty in-memory metric storage.
func NewMemStorage() *MemStorage {
	return &MemStorage{
		mutex:   sync.RWMutex{},
		metrics: make(map[string]model.Metric),
	}
}

// UpdateMetric updates a single metric.
func (s *MemStorage) UpdateMetric(ctx context.Context, metric model.Metric) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.update(metric)
}

// UpdateAll batch updates a set of metrics under a single lock.
func (s *MemStorage) UpdateAll(ctx context.Context, metrics []model.Metric) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for _, m := range metrics {
		s.update(m)
	}
	return nil
}

func (s *MemStorage) update(metric model.Metric) error {
	switch metric.MType {
	case model.Counter:
		if v, ok := s.metrics[metric.ID]; ok {
			*v.Delta += *metric.Delta
			s.metrics[metric.ID] = v
		} else {
			s.metrics[metric.ID] = metric
		}
	case model.Gauge:
		s.metrics[metric.ID] = metric
	default:
		return errors.New("unsupported metric type")
	}
	return nil
}

// GetValues returns a slice with all metrics from the storage.
func (s *MemStorage) GetValues(ctx context.Context) ([]model.Metric, error) {
	res := make([]model.Metric, 0, len(s.metrics))
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	for k := range s.metrics {
		res = append(res, s.metrics[k])
	}

	return res, nil
}

// GetValuesAndClear returns all metrics and clears the storage.
// It is used by the agent to consume metrics between reports.
func (s *MemStorage) GetValuesAndClear() []model.Metric {
	res := make([]model.Metric, 0, len(s.metrics))
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for k := range s.metrics {
		res = append(res, s.metrics[k])
	}
	s.metrics = make(map[string]model.Metric)

	return res
}

// GetMetric returns a metric by identifier.
func (s *MemStorage) GetMetric(ctx context.Context, id string) (*model.Metric, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	m, ok := s.metrics[id]
	if !ok {
		return nil, errors.New("not found")
	}
	return &m, nil
}
