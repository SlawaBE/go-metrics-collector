package service

import (
	"context"
	"testing"

	"github.com/SlawaBE/go-metrics-collector/internal/model"
	"github.com/SlawaBE/go-metrics-collector/internal/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type mockSaver struct {
	saveCalls int
}

func (m *mockSaver) Load() error { return nil }
func (m *mockSaver) SaveSync() error {
	m.saveCalls++
	return nil
}
func (m *mockSaver) StartSync(ctx context.Context) {}

func newServiceWithSaver() (*MetricsService, *mockSaver) {
	stor := storage.NewMemStorage()
	saver := &mockSaver{}
	return NewMetricsService(stor, saver), saver
}

func TestMetricsService_UpdateMetric(t *testing.T) {
	svc, saver := newServiceWithSaver()

	err := svc.UpdateMetric(context.Background(), model.NewCounterMetric("counter-1", 10))
	require.NoError(t, err)
	assert.Equal(t, 1, saver.saveCalls, "saver should be called after a successful update")
}

func TestMetricsService_UpdateMetric_EmptyName(t *testing.T) {
	svc, saver := newServiceWithSaver()

	err := svc.UpdateMetric(context.Background(), model.Metric{ID: "", MType: model.Counter})
	assert.Error(t, err)
	assert.Equal(t, 0, saver.saveCalls, "saver must not be called on validation error")
}

func TestMetricsService_UpdateMetric_StorageErrorNoSave(t *testing.T) {
	svc, saver := newServiceWithSaver()

	err := svc.UpdateMetric(context.Background(), model.Metric{ID: "bad", MType: "unknown"})
	assert.Error(t, err)
	assert.Equal(t, 0, saver.saveCalls, "saver must not be called when storage update fails")
}

func TestMetricsService_UpdateMetrics(t *testing.T) {
	svc, saver := newServiceWithSaver()

	metrics := []model.Metric{
		model.NewCounterMetric("counter-1", 10),
		model.NewGaugeMetric("gauge-1", 1.5),
	}

	err := svc.UpdateMetrics(context.Background(), metrics)
	require.NoError(t, err)
	assert.Equal(t, 1, saver.saveCalls)
}

func TestMetricsService_GetMetric(t *testing.T) {
	svc, _ := newServiceWithSaver()
	ctx := context.Background()

	require.NoError(t, svc.UpdateMetric(ctx, model.NewCounterMetric("counter-1", 10)))

	metric, err := svc.GetMetric(ctx, model.MetricRequest{ID: "counter-1", MType: model.Counter})
	require.NoError(t, err)
	require.NotNil(t, metric)
	assert.Equal(t, int64(10), *metric.Delta)
}

func TestMetricsService_GetMetric_NotFound(t *testing.T) {
	svc, _ := newServiceWithSaver()
	ctx := context.Background()

	require.NoError(t, svc.UpdateMetric(ctx, model.NewCounterMetric("counter-1", 10)))

	_, err := svc.GetMetric(ctx, model.MetricRequest{ID: "missing", MType: model.Counter})
	assert.Error(t, err)
}

func TestMetricsService_GetMetric_TypeMismatch(t *testing.T) {
	svc, _ := newServiceWithSaver()
	ctx := context.Background()

	require.NoError(t, svc.UpdateMetric(ctx, model.NewCounterMetric("counter-1", 10)))

	_, err := svc.GetMetric(ctx, model.MetricRequest{ID: "counter-1", MType: model.Gauge})
	assert.Error(t, err)
}

func TestMetricsService_List_SortedAndFormatted(t *testing.T) {
	svc, _ := newServiceWithSaver()
	ctx := context.Background()

	require.NoError(t, svc.UpdateMetric(ctx, model.NewCounterMetric("counter-2", 10)))
	require.NoError(t, svc.UpdateMetric(ctx, model.NewGaugeMetric("gauge-1", 1.5)))
	require.NoError(t, svc.UpdateMetric(ctx, model.NewCounterMetric("counter-1", 42)))

	list, err := svc.List(ctx)
	require.NoError(t, err)

	assert.Equal(t, []string{"counter-1: 42", "counter-2: 10", "gauge-1: 1.5"}, list)
}

func TestMetricsService_List_Empty(t *testing.T) {
	svc, _ := newServiceWithSaver()

	list, err := svc.List(context.Background())
	require.NoError(t, err)
	assert.Empty(t, list)
}

func TestMetricsService_WithoutSaver(t *testing.T) {
	svc := NewMetricsService(storage.NewMemStorage(), nil)

	err := svc.UpdateMetric(context.Background(), model.NewCounterMetric("counter-1", 10))
	require.NoError(t, err, "a nil saver must not panic")
}