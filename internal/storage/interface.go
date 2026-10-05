package storage

import (
	"context"

	"github.com/SlawaBE/go-metrics-collector/internal/model"
)

// Storage defines the single contract for a metric storage.
type Storage interface {
	// UpdateMetric updates a single metric.
	UpdateMetric(ctx context.Context, metric model.Metric) error
	// GetValues returns all metrics from the storage.
	GetValues(ctx context.Context) ([]model.Metric, error)
	// GetMetric returns a metric by identifier.
	GetMetric(ctx context.Context, id string) (*model.Metric, error)
	// UpdateAll batch updates a set of metrics.
	UpdateAll(ctx context.Context, metrics []model.Metric) error
}
