package service

type FileMetricSaver interface {
	Load() error
	SaveSync() error
	StartSync()
}
