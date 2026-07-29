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

type JsonFileMetricSaver struct {
	interval int
	fileName string
	storage  storage.Storage
}

func NewJsonFileMetricSaver(interval int, fileName string, storage storage.Storage) *JsonFileMetricSaver {
	return &JsonFileMetricSaver{
		interval: interval,
		fileName: fileName,
		storage:  storage,
	}
}

func (m *JsonFileMetricSaver) Load() error {
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
	err = m.storage.UpdateAll(metrics)
	if err != nil {
		logger.Log.Error("Error updating metrics in storage", zap.Error(err))
		return err
	}
	return nil
}

func (m *JsonFileMetricSaver) save() error {
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

func (m *JsonFileMetricSaver) SaveSync() error {
	if m.interval == 0 {
		return m.save()
	}
	return nil
}

func (m *JsonFileMetricSaver) StartSync(ctx context.Context) {
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
