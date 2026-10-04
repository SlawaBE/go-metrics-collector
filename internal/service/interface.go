package service

import (
	"context"

	"github.com/SlawaBE/go-metrics-collector/internal/model"
)

// FileMetricServer is an interface for synchronizing stored metrics to a file.
//
// It is used only when working with the application’s internal in‑memory storage.
type FileMetricSaver interface {
	// Load loads metric values from a file when the application starts.
	Load() error

	// SaveSync saves the values of all metrics from the storage in a file.
	SaveSync() error

	// StartSync initiates synchronization.
	StartSync(ctx context.Context)
}

// AuditSubscriber is an interface for listening to audit events and performing useful work when they are received.
type AuditSubscriber interface {
	// Notify using for processing audit event.
	Notify(ctx context.Context, event model.AuditEvent) error

	// Name return name of subscriber. It is using for errors message in logs.
	Name() string
}

// AuditPublisher is an interface for publishing audit events when metrics are received.
type AuditPublisher interface {
	// Subscribe subscribes for listening audit events
	Subscribe(subscriber AuditSubscriber)

	// SendMetric notified all subscribers about one obtained metric
	SendMetric(ctx context.Context, ip string, metric model.Metric)

	// SendMetric notified all subscribers about many obtained metrics
	SendMetrics(ctx context.Context, ip string, metrics []model.Metric)
}
