package storage

import "github.com/SlawaBE/go-metrics-collector/internal/model"

type Storage interface {
	UpdateMetric(metric model.Metric) error
	GetValues() []model.Metric
}
