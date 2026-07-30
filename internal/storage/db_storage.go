package storage

import (
	"context"
	"database/sql"

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

const (
	INSERT_METRIC = `INSERT INTO metrics (id, type, delta, value) VALUES ($1, $2, $3, $4)
						ON CONFLICT (id)
						DO UPDATE SET delta = metrics.delta + $3, value = $4;`
	SELECT_METRICS = `SELECT id, type, delta, value
						FROM metrics`
	SELECT_METRIC = `SELECT id, type, delta, value
						FROM metrics
						WHERE id = $1`
)

func (s *DBStorage) UpdateMetric(ctx context.Context, metric model.Metric) error {
	tx, err := s.db.Begin()
	if err != nil {
		logger.Log.Error("error begin transaction", zap.Error(err))
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.PrepareContext(ctx, INSERT_METRIC)
	if err != nil {
		logger.Log.Error("error prepare statement", zap.Error(err))
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, metric.ID, metric.MType, metric.Delta, metric.Value)
	if err != nil {
		logger.Log.Error("error exec statement", zap.Error(err))
		return err
	}

	err = tx.Commit()
	if err != nil {
		logger.Log.Error("error commit transaction", zap.Error(err))
	}
	return err
}

func (s *DBStorage) GetValues(ctx context.Context) ([]model.Metric, error) {
	metrics := make([]model.Metric, 0)
	rows, err := s.db.QueryContext(ctx, SELECT_METRICS)
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
	rows := s.db.QueryRowContext(ctx, SELECT_METRIC, id)

	if err := rows.Scan(&metric.ID, &metric.MType, &metric.Delta, &metric.Value); err != nil {
		logger.Log.Error("error get metric", zap.Error(err))
		return nil, err
	}
	return &metric, nil
}

func (s *DBStorage) UpdateAll(ctx context.Context, metrics []model.Metric) error {
	tx, err := s.db.Begin()
	if err != nil {
		logger.Log.Error("error begin transaction", zap.Error(err))
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.PrepareContext(ctx, INSERT_METRIC)
	if err != nil {
		logger.Log.Error("error prepare statement", zap.Error(err))
		return err
	}
	defer stmt.Close()

	for _, m := range metrics {
		_, err := stmt.ExecContext(ctx, m.ID, m.MType, m.Delta, m.Value)
		if err != nil {
			logger.Log.Error("error exec statement", zap.Error(err))
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		logger.Log.Error("error commit transaction", zap.Error(err))
	}
	return err
}
