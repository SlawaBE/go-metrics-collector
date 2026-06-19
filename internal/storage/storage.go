package storage

import (
	"errors"
	"sync"

	"github.com/SlawaBE/go-metrics-collector/internal/model"
)

type MemStorage struct {
	mutex      sync.RWMutex
	metrics map[string]model.Metric
}

func NewMemStorage() *MemStorage {
	return &MemStorage{
	    mutex: sync.RWMutex{},
		metrics: make(map[string]model.Metric),
	}
}

func (s *MemStorage) SaveCounter(key string, value int64) error {
	s.UpdateMetric(model.NewCounterMetric(key, value))
	return nil
}

func (s *MemStorage) SaveGauge(key string, value float64) error {
	s.UpdateMetric(model.NewGaugeMetric(key, value))
	return nil
}

func (s *MemStorage) UpdateMetric(metric model.Metric) error {
    s.mutex.Lock()
    defer s.mutex.Unlock()

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

func (s *MemStorage) GetValues() []model.Metric {
	res := make([]model.Metric, 0, len(s.metrics))
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	for k := range s.metrics {
		res = append(res, s.metrics[k])
	}

	return res
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

func (s *MemStorage) GetMetric(id string) (*model.Metric, error) {
    s.mutex.RLock()
    defer s.mutex.RUnlock()

    m, ok := s.metrics[id]
    if !ok {
        return nil, errors.New("not found")
    }
    return &m, nil
}
