package db

import (
	"database/sql"

	"github.com/SlawaBE/go-metrics-collector/internal/logger"
	_ "github.com/jackc/pgx/v5/stdlib"
	"go.uber.org/zap"
)

func NewDB(databaseDSN string) (*sql.DB, error) {
	db, err := sql.Open("pgx", databaseDSN)
	if err != nil {
		logger.Log.Fatal("error open db connect", zap.Error(err))
		return nil, err
	}

	return db, nil
}
