package storage

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/SlawaBE/go-metrics-collector/internal/logger"
	"github.com/SlawaBE/go-metrics-collector/internal/model"
	"go.uber.org/zap"
)

type DBStorage struct {
	db *sql.DB
}

func NewDBStorage(db *sql.DB) *DBStorage {
	return &DBStorage{
		db: db,
	}
}

func (s *DBStorage) UpdateMetric(ctx context.Context, metric model.Metric) error {
	_, err := s.db.ExecContext(ctx, `INSERT INTO metrics (id, type, delta, value) values ($1, $2, $3, $4)
							ON CONFLICT (id)
							DO UPDATE SET delta = metrics.delta + $3, value = $4;`,
		metric.ID, metric.MType, metric.Delta, metric.Value)
	if err != nil {
		logger.Log.Error("error save metric", zap.Error(err))
		return err
	}
	return nil
}

func (s *DBStorage) GetValues(ctx context.Context) ([]model.Metric, error) {
	metrics := make([]model.Metric, 0)
	rows, err := s.db.QueryContext(ctx, "SELECT id, type, delta, value from metrics")
	if err != nil {
		logger.Log.Error("error get all metric", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var metric model.Metric
		if err := rows.Scan(&metric.ID, &metric.MType, &metric.Delta, &metric.Value); err != nil {
			logger.Log.Error("error get metric", zap.Error(err))
			return nil, err
		}
		metrics = append(metrics, metric)
	}

	err = rows.Err()
	if err != nil {
		logger.Log.Error("error get all metric", zap.Error(err))
		return nil, err
	}

	return metrics, nil
}

func (s *DBStorage) GetMetric(ctx context.Context, id string) (*model.Metric, error) {
	var metric model.Metric
	rows := s.db.QueryRowContext(ctx, "SELECT id, type, delta, value FROM metrics WHERE id = $1", id)

	if err := rows.Scan(&metric.ID, &metric.MType, &metric.Delta, &metric.Value); err != nil {
		logger.Log.Error("error get metric", zap.Error(err))
		return nil, err
	}
	return &metric, nil
}

func (s *DBStorage) UpdateAll(metrics []model.Metric) error {
	return fmt.Errorf("unsupported operation 'UpdateAll'")
}
