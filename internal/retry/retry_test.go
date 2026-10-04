package retry

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/assert"
)

func TestIsTransportError(t *testing.T) {
	transportErr := &pgconn.PgError{Code: "08001"}
	assert.True(t, IsTransportError(transportErr))

	assert.False(t, IsTransportError(errors.New("some error")))

	assert.False(t, IsTransportError(&pgconn.PgError{Code: "23505"}))
	assert.False(t, IsTransportError(nil))
}

func TestRetryWithBackoff_SuccessFirstTry(t *testing.T) {
	calls := 0
	err := RetryWithBackoff(context.Background(), func() error {
		calls++
		return nil
	})
	assert.NoError(t, err)
	assert.Equal(t, 1, calls)
}

func TestRetryWithBackoff_NonTransportErrorNoRetry(t *testing.T) {
	calls := 0
	expected := errors.New("permanent error")
	err := RetryWithBackoff(context.Background(), func() error {
		calls++
		return expected
	})
	assert.ErrorIs(t, err, expected)
	assert.Equal(t, 1, calls, "non-transport errors must not be retried")
}

func TestRetryWithBackoff_TransportErrorThenSuccess(t *testing.T) {
	original := retryIntervals
	retryIntervals = []time.Duration{1 * time.Millisecond}
	defer func() { retryIntervals = original }()

	calls := 0
	err := RetryWithBackoff(context.Background(), func() error {
		calls++
		if calls == 1 {
			return &pgconn.PgError{Code: "08001"}
		}
		return nil
	})
	assert.NoError(t, err)
	assert.Equal(t, 2, calls)
}

func TestRetryWithBackoff_TransportErrorExhausted(t *testing.T) {
	original := retryIntervals
	retryIntervals = []time.Duration{1 * time.Millisecond, 1 * time.Millisecond}
	defer func() { retryIntervals = original }()

	calls := 0
	transportErr := &pgconn.PgError{Code: "08001"}
	err := RetryWithBackoff(context.Background(), func() error {
		calls++
		return transportErr
	})
	assert.Equal(t, transportErr, err)
	assert.Equal(t, 1+len(retryIntervals), calls)
}

func TestRetryWithBackoff_ContextCancelled(t *testing.T) {
	original := retryIntervals
	retryIntervals = []time.Duration{50 * time.Millisecond}
	defer func() { retryIntervals = original }()

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	transportErr := &pgconn.PgError{Code: "08001"}
	err := RetryWithBackoff(ctx, func() error {
		return transportErr
	})
	assert.ErrorIs(t, err, context.Canceled)
}
