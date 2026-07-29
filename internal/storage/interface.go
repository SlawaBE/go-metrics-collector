package storage

import (
	"context"

	"github.com/SlawaBE/go-metrics-collector/internal/model"
)

type Storage interface {
	UpdateMetric(ctx context.Context, metric model.Metric) error
	GetValues(ctx context.Context, ) ([]model.Metric, error)
	GetMetric(ctx context.Context, id string) (*model.Metric, error)
	UpdateAll(metrics []model.Metric) error
}
