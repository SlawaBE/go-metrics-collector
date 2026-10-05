package service

import (
	"context"
	"encoding/json"
	"os"
	"time"

	"github.com/SlawaBE/go-metrics-collector/internal/logger"
	"github.com/SlawaBE/go-metrics-collector/internal/model"
	"github.com/SlawaBE/go-metrics-collector/internal/storage"
	"go.uber.org/zap"
)

// JSONFileMetricSaver persists metrics to a file in JSON format, both
// synchronously and with a configurable interval.
type JSONFileMetricSaver struct {
	interval int
	fileName string
	storage  storage.Storage
}

// NewJSONFileMetricSaver creates a metrics saver: with interval == 0 saving is
// synchronous, with a positive interval it happens in the background via a ticker.
func NewJSONFileMetricSaver(interval int, fileName string, storage storage.Storage) *JSONFileMetricSaver {
	return &JSONFileMetricSaver{
		interval: interval,
		fileName: fileName,
		storage:  storage,
	}
}

// Load loads metrics from a file into the storage.
func (m *JSONFileMetricSaver) Load() error {
	data, err := os.ReadFile(m.fileName)
	if err != nil {
		logger.Log.Error("Error loading metrics", zap.Error(err))
		return err
	}
	var metrics []model.Metric
	err = json.Unmarshal(data, &metrics)
	if err != nil {
		logger.Log.Error("Error loading metrics", zap.Error(err))
		return err
	}
	err = m.storage.UpdateAll(context.Background(), metrics)
	if err != nil {
		logger.Log.Error("Error updating metrics in storage", zap.Error(err))
		return err
	}
	return nil
}

func (m *JSONFileMetricSaver) save() error {
	list, err := m.storage.GetValues(context.Background())
	if err != nil {
		logger.Log.Error("Error saving metrics", zap.Error(err))
		return err
	}

	data, err := json.Marshal(list)
	if err != nil {
		logger.Log.Error("Error saving metrics", zap.Error(err))
		return err
	}
	return os.WriteFile(m.fileName, data, 0644)
}

// SaveSync saves metrics to a file if the save interval is 0.
func (m *JSONFileMetricSaver) SaveSync() error {
	if m.interval == 0 {
		return m.save()
	}
	return nil
}

// StartSync starts background persistence of metrics to a file at the given
// interval. No background saving is started for a non-positive interval.
func (m *JSONFileMetricSaver) StartSync(ctx context.Context) {
	if m.interval <= 0 {
		return
	}
	ticker := time.NewTicker(time.Duration(m.interval) * time.Second)

	go func() {
		logger.Log.Info("Start sync")
		for {
			select {
			case <-ctx.Done():
				logger.Log.Info("Stop sync")
				ticker.Stop()
				return

			case <-ticker.C:
				logger.Log.Info("Save metrics")
				err := m.save()
				if err != nil {
					logger.Log.Error("Error saving metrics", zap.Error(err))
				}
			}
		}
	}()
}
