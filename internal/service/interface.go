package service

import (
	"context"

	"github.com/SlawaBE/go-metrics-collector/internal/model"
)

type FileMetricSaver interface {
	Load() error
	SaveSync() error
	StartSync(ctx context.Context)
}

type AuditSubscriber interface {
	Notify(ctx context.Context, event model.AuditEvent) error
	Name() string
}

type AuditPublisher interface {
	Subscribe(subscriber AuditSubscriber)
	SendMetric(ctx context.Context, ip string, metric model.Metric)
	SendMetrics(ctx context.Context, ip string, metrics []model.Metric)
}
