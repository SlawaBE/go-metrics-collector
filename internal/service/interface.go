package service

import (
	"context"

	"github.com/SlawaBE/go-metrics-collector/internal/model"
)

// FileMetricSaver defines the contract for synchronous and background
// persistence of metrics to a file.
type FileMetricSaver interface {
	// Load loads metrics from a file into the storage.
	Load() error
	// SaveSync synchronously saves metrics to a file (if the interval is 0).
	SaveSync() error
	// StartSync starts background periodic persistence of metrics to a file.
	StartSync(ctx context.Context)
}

// AuditSubscriber receives and processes audit events.
type AuditSubscriber interface {
	// Notify processes a single audit event.
	Notify(event model.AuditEvent) error
	// Name returns the subscriber name for logging.
	Name() string
	// Close releases resources held by the subscriber.
	Close() error
}

// AuditPublisher publishes audit events to subscribers. Events are delivered asynchronously.
type AuditPublisher interface {
	// Subscribe registers a subscriber for audit events.
	Subscribe(subscriber AuditSubscriber)
	// SendMetric publishes an audit event for a single metric.
	SendMetric(address string, metric model.Metric)
	// SendMetrics publishes an audit event for a set of metrics.
	SendMetrics(address string, metrics []model.Metric)
}
