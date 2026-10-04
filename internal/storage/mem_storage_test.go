package storage

import (
	"context"
	"testing"

	"github.com/SlawaBE/go-metrics-collector/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMemStorage_UpdateMetric_Counter(t *testing.T) {
	s := NewMemStorage()
	ctx := context.Background()

	err := s.UpdateMetric(ctx, model.NewCounterMetric("counter-1", 10))
	require.NoError(t, err)

	err = s.UpdateMetric(ctx, model.NewCounterMetric("counter-1", 5))
	require.NoError(t, err)

	metric, err := s.GetMetric(ctx, "counter-1")
	require.NoError(t, err)
	require.NotNil(t, metric.Delta)
	assert.Equal(t, int64(15), *metric.Delta)
	assert.Equal(t, model.Counter, metric.MType)
}

func TestMemStorage_UpdateMetric_Gauge(t *testing.T) {
	s := NewMemStorage()
	ctx := context.Background()

	err := s.UpdateMetric(ctx, model.NewGaugeMetric("gauge-1", 1.5))
	require.NoError(t, err)

	err = s.UpdateMetric(ctx, model.NewGaugeMetric("gauge-1", 2.5))
	require.NoError(t, err)

	metric, err := s.GetMetric(ctx, "gauge-1")
	require.NoError(t, err)
	require.NotNil(t, metric.Value)
	assert.Equal(t, 2.5, *metric.Value)
	assert.Equal(t, model.Gauge, metric.MType)
}

func TestMemStorage_UpdateMetric_UnsupportedType(t *testing.T) {
	s := NewMemStorage()

	metric := model.Metric{ID: "bad", MType: "unknown"}

	err := s.UpdateMetric(context.Background(), metric)
	assert.Error(t, err)
}

func TestMemStorage_GetMetric_NotFound(t *testing.T) {
	s := NewMemStorage()

	_, err := s.GetMetric(context.Background(), "missing")
	assert.Error(t, err)
}

func TestMemStorage_GetValues(t *testing.T) {
	s := NewMemStorage()
	ctx := context.Background()

	require.NoError(t, s.UpdateMetric(ctx, model.NewCounterMetric("counter-1", 10)))
	require.NoError(t, s.UpdateMetric(ctx, model.NewCounterMetric("counter-2", 20)))
	require.NoError(t, s.UpdateMetric(ctx, model.NewGaugeMetric("gauge-1", 1.5)))

	metrics, err := s.GetValues(ctx)
	require.NoError(t, err)
	assert.Len(t, metrics, 3)

	ids := make(map[string]bool)
	for _, m := range metrics {
		ids[m.ID] = true
	}
	assert.True(t, ids["counter-1"])
	assert.True(t, ids["counter-2"])
	assert.True(t, ids["gauge-1"])
}

func TestMemStorage_GetValues_Empty(t *testing.T) {
	s := NewMemStorage()

	metrics, err := s.GetValues(context.Background())
	require.NoError(t, err)
	assert.Empty(t, metrics)
}

func TestMemStorage_UpdateAll(t *testing.T) {
	s := NewMemStorage()
	ctx := context.Background()

	metrics := []model.Metric{
		model.NewCounterMetric("counter-1", 10),
		model.NewGaugeMetric("gauge-1", 1.5),
	}

	err := s.UpdateAll(ctx, metrics)
	require.NoError(t, err)

	err = s.UpdateAll(ctx, []model.Metric{model.NewCounterMetric("counter-1", 5)})
	require.NoError(t, err)

	metric, err := s.GetMetric(ctx, "counter-1")
	require.NoError(t, err)
	require.NotNil(t, metric.Delta)
	assert.Equal(t, int64(15), *metric.Delta)
}

func TestMemStorage_GetValuesAndClear(t *testing.T) {
	s := NewMemStorage()
	ctx := context.Background()

	require.NoError(t, s.UpdateMetric(ctx, model.NewCounterMetric("counter-1", 10)))
	require.NoError(t, s.UpdateMetric(ctx, model.NewGaugeMetric("gauge-1", 1.5)))

	metrics := s.GetValuesAndClear()
	assert.Len(t, metrics, 2)

	remaining, err := s.GetValues(ctx)
	require.NoError(t, err)
	assert.Empty(t, remaining)
}
