package storage

import (
	"context"

	"github.com/SlawaBE/go-metrics-collector/internal/model"
)

// Storage is an interface for working with the metric storage.
type Storage interface {
	// Update metric in storage.
	//
	// ctx may usage for cancelling operation.
	UpdateMetric(ctx context.Context, metric model.Metric) error

	// Get all metrics from storage.
	//
	// ctx may usage for cancelling operation.
	GetValues(ctx context.Context) ([]model.Metric, error)

	// Get metric from storage by id
	//
	// ctx may usage for cancelling operation.
	GetMetric(ctx context.Context, id string) (*model.Metric, error)

	// Update metrics from slice in storage.
	//
	// ctx may usage for cancelling operation.
	UpdateAll(ctx context.Context, metrics []model.Metric) error
}
