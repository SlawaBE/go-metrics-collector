package service

import "context"

type FileMetricSaver interface {
	Load() error
	SaveSync() error
	StartSync(ctx context.Context)
}
