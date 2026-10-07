// Package retry provides a mechanism for retrying operations with increasing
// backoff on database transport errors.
package retry

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/SlawaBE/go-metrics-collector/internal/logger"
	"github.com/jackc/pgx/v5/pgconn"
	"go.uber.org/zap"
)

// retryIntervals holds increasing retry intervals (in seconds).
var retryIntervals = []time.Duration{1 * time.Second, 3 * time.Second, 5 * time.Second}

// IsTransportError returns true if the error is a PostgreSQL transport error
// (SQLSTATE class 08 - connection errors).
func IsTransportError(err error) bool {
	var pgErr *pgconn.PgError
	return errors.As(err, &pgErr) && strings.HasPrefix(pgErr.Code, "08")
}

// RetryWithBackoff executes fn, retrying it on transport errors with increasing
// backoff, until it succeeds, the context is cancelled, or a non-transport
// error occurs.
func RetryWithBackoff(ctx context.Context, fn func() error) error {
	if err := fn(); err != nil {
		if !IsTransportError(err) {
			return err
		}
		for i, interval := range retryIntervals {
			logger.Log.Warn("transport error, retrying",
				zap.Int("attempt", i+1),
				zap.Duration("wait", interval),
				zap.Error(err))
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(interval):
			}
			if err = fn(); err == nil {
				return nil
			}
			if !IsTransportError(err) {
				return err
			}
		}
		logger.Log.Error("transport error, attempts exceeded",
			zap.Error(err))
		return err
	}
	return nil
}
