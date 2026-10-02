package agent

import (
	"context"
	"testing"
	"time"

	"github.com/SlawaBE/go-metrics-collector/internal/model"
	"github.com/SlawaBE/go-metrics-collector/internal/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPoller_Poll(t *testing.T) {
	stor := storage.NewMemStorage()

	poller := NewPoller(stor, 1, func() []model.Metric {
		return []model.Metric{
			model.NewCounterMetric("c1", 10),
			model.NewGaugeMetric("g1", 1.5),
		}
	})

	poller.Poll()

	metrics, err := stor.GetValues(context.Background())
	require.NoError(t, err)
	require.Len(t, metrics, 2)
}

func TestPoller_Poll_EmptyMetrics(t *testing.T) {
	stor := storage.NewMemStorage()

	poller := NewPoller(stor, 1, func() []model.Metric {
		return nil
	})

	poller.Poll()

	metrics, err := stor.GetValues(context.Background())
	require.NoError(t, err)
	assert.Empty(t, metrics)
}

func TestPoller_Run(t *testing.T) {
	stor := storage.NewMemStorage()

	calls := 0
	poller := NewPoller(stor, 1, func() []model.Metric {
		calls++
		return []model.Metric{model.NewCounterMetric("c1", int64(calls))}
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go poller.Run(ctx)
	time.Sleep(1500 * time.Millisecond)
	cancel()

	assert.GreaterOrEqual(t, calls, 1)

	metrics, err := stor.GetValues(context.Background())
	require.NoError(t, err)
	assert.Len(t, metrics, 1)
}