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
	Notify(ctx context.Context, event model.AuditEvent) error
	// Name returns the subscriber name for logging.
	Name() string
}

// AuditPublisher publishes audit events to subscribers.
type AuditPublisher interface {
	// Subscribe registers a subscriber for audit events.
	Subscribe(subscriber AuditSubscriber)
	// SendMetric publishes an audit event for a single metric.
	SendMetric(ctx context.Context, ip string, metric model.Metric)
	// SendMetrics publishes an audit event for a set of metrics.
	SendMetrics(ctx context.Context, ip string, metrics []model.Metric)
}
