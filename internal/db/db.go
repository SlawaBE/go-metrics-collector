// Package db handles establishing a PostgreSQL connection and running the
// embedded migrations.
package db

import (
	"database/sql"
	"embed"
	"fmt"

	"github.com/SlawaBE/go-metrics-collector/internal/logger"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "github.com/jackc/pgx/v5/stdlib"
	"go.uber.org/zap"
)

// NewDB opens a PostgreSQL connection using the given DSN string.
func NewDB(databaseDSN string) (*sql.DB, error) {
	db, err := sql.Open("pgx", databaseDSN)
	if err != nil {
		logger.Log.Fatal("error open db connect", zap.Error(err))
		return nil, err
	}

	return db, nil
}

//go:embed migrations/*.sql
var migrationFiles embed.FS

// RunMigrations applies the embedded SQL migrations to the database connection.
func RunMigrations(database *sql.DB, databaseDSN string) error {
	srcDriver, err := iofs.New(migrationFiles, "migrations")
	if err != nil {
		return fmt.Errorf("error create source driver %w", err)
	}

	_, err = postgres.WithInstance(database, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("error create db driver %w", err)
	}

	m, err := migrate.NewWithSourceInstance("iofs", srcDriver, databaseDSN)
	if err != nil {
		return fmt.Errorf("error create source %w", err)
	}
	defer func() {
		if _, err := m.Close(); err != nil {
			logger.Log.Error("error close migration source", zap.Error(err))
		}
	}()

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("error migrate %w", err)
	}
	return nil
}
