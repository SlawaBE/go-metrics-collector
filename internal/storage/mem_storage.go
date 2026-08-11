package storage

import (
	"context"
	"errors"
	"sync"

	"github.com/SlawaBE/go-metrics-collector/internal/model"
)

type MemStorage struct {
	mutex   sync.RWMutex
	metrics map[string]model.Metric
}

func NewMemStorage() *MemStorage {
	return &MemStorage{
		mutex:   sync.RWMutex{},
		metrics: make(map[string]model.Metric),
	}
}

func (s *MemStorage) UpdateMetric(ctx context.Context, metric model.Metric) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.update(metric)
}

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

func (s *MemStorage) GetValues(ctx context.Context) ([]model.Metric, error) {
	res := make([]model.Metric, 0, len(s.metrics))
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	for k := range s.metrics {
		res = append(res, s.metrics[k])
	}

	return res, nil
}

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

func (s *MemStorage) GetMetric(ctx context.Context, id string) (*model.Metric, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	m, ok := s.metrics[id]
	if !ok {
		return nil, errors.New("not found")
	}
	return &m, nil
}
