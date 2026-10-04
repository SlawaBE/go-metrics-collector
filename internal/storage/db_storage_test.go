package storage

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/SlawaBE/go-metrics-collector/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newDBStorage(t *testing.T) (*DBStorage, sqlmock.Sqlmock) {
	t.Helper()
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	return NewDBStorage(db), mock
}

func TestDBStorage_UpdateMetric(t *testing.T) {
	s, mock := newDBStorage(t)
	defer s.db.Close()

	mock.ExpectBegin()
	mock.ExpectPrepare(`INSERT INTO metrics`)
	mock.ExpectExec(`INSERT INTO metrics`).
		WithArgs("counter-1", model.Counter, sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	err := s.UpdateMetric(context.Background(), model.NewCounterMetric("counter-1", 10))
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDBStorage_UpdateMetric_BeginError(t *testing.T) {
	s, mock := newDBStorage(t)
	defer s.db.Close()

	mock.ExpectBegin().WillReturnError(errors.New("begin failed"))

	err := s.UpdateMetric(context.Background(), model.NewCounterMetric("counter-1", 10))
	assert.Error(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDBStorage_UpdateMetric_ExecError(t *testing.T) {
	s, mock := newDBStorage(t)
	defer s.db.Close()

	mock.ExpectBegin()
	mock.ExpectPrepare(`INSERT INTO metrics`)
	mock.ExpectExec(`INSERT INTO metrics`).
		WillReturnError(errors.New("exec failed"))
	mock.ExpectRollback()

	err := s.UpdateMetric(context.Background(), model.NewCounterMetric("counter-1", 10))
	assert.Error(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDBStorage_GetMetric(t *testing.T) {
	s, mock := newDBStorage(t)
	defer s.db.Close()

	rows := sqlmock.NewRows([]string{"id", "type", "delta", "value"}).
		AddRow("counter-1", model.Counter, int64(10), nil)

	mock.ExpectQuery(`SELECT id, type, delta, value FROM metrics WHERE id = \$1`).
		WithArgs("counter-1").
		WillReturnRows(rows)

	m, err := s.GetMetric(context.Background(), "counter-1")
	assert.NoError(t, err)
	require.NotNil(t, m)
	assert.Equal(t, "counter-1", m.ID)
	require.NotNil(t, m.Delta)
	assert.Equal(t, int64(10), *m.Delta)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDBStorage_GetMetric_NotFound(t *testing.T) {
	s, mock := newDBStorage(t)
	defer s.db.Close()

	mock.ExpectQuery(`SELECT id, type, delta, value FROM metrics WHERE id = \$1`).
		WithArgs("missing").
		WillReturnError(sql.ErrNoRows)

	m, err := s.GetMetric(context.Background(), "missing")
	assert.Error(t, err)
	assert.Nil(t, m)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDBStorage_GetValues(t *testing.T) {
	s, mock := newDBStorage(t)
	defer s.db.Close()

	rows := sqlmock.NewRows([]string{"id", "type", "delta", "value"}).
		AddRow("counter-1", model.Counter, int64(10), nil).
		AddRow("gauge-1", model.Gauge, nil, 1.5)

	mock.ExpectQuery(`SELECT id, type, delta, value FROM metrics`).
		WillReturnRows(rows)

	metrics, err := s.GetValues(context.Background())
	assert.NoError(t, err)
	require.Len(t, metrics, 2)
	assert.Equal(t, "counter-1", metrics[0].ID)
	assert.Equal(t, "gauge-1", metrics[1].ID)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDBStorage_GetValues_Error(t *testing.T) {
	s, mock := newDBStorage(t)
	defer s.db.Close()

	mock.ExpectQuery(`SELECT id, type, delta, value FROM metrics`).
		WillReturnError(errors.New("query failed"))

	_, err := s.GetValues(context.Background())
	assert.Error(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDBStorage_UpdateAll(t *testing.T) {
	s, mock := newDBStorage(t)
	defer s.db.Close()

	metrics := []model.Metric{
		model.NewCounterMetric("counter-1", 10),
		model.NewGaugeMetric("gauge-1", 1.5),
	}

	mock.ExpectBegin()
	mock.ExpectPrepare(`INSERT INTO metrics`)
	mock.ExpectExec(`INSERT INTO metrics`).
		WithArgs("counter-1", model.Counter, sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectExec(`INSERT INTO metrics`).
		WithArgs("gauge-1", model.Gauge, sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	err := s.UpdateAll(context.Background(), metrics)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
