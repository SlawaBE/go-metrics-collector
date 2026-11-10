package grpchandler

import (
	"context"
	"testing"

	"github.com/SlawaBE/go-metrics-collector/internal/storage"
	pb "github.com/SlawaBE/go-metrics-collector/internal/proto/metrics"
	"github.com/SlawaBE/go-metrics-collector/internal/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func newTestServer() (*MetricsServer, *storage.MemStorage) {
	stor := storage.NewMemStorage()
	svc := service.NewMetricsService(stor, nil)
	return NewMetricsServer(svc, nil).(*MetricsServer), stor
}

func TestNewMetricsServer(t *testing.T) {
	stor := storage.NewMemStorage()
	svc := service.NewMetricsService(stor, nil)

	srv := NewMetricsServer(svc, nil)
	require.NotNil(t, srv)

	ms, ok := srv.(*MetricsServer)
	require.True(t, ok)
	assert.NotNil(t, ms.metricsService)
}

func TestUpdateMetrics_Success_Counter(t *testing.T) {
	srv, stor := newTestServer()

	req := pb.UpdateMetricsRequest_builder{
		Metrics: []*pb.Metric{
			pb.Metric_builder{
				Id:    "c1",
				Type:  pb.Metric_COUNTER,
				Delta: 42,
			}.Build(),
		},
	}.Build()

	resp, err := srv.UpdateMetrics(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, resp)

	metric, err := stor.GetMetric(context.Background(), "c1")
	require.NoError(t, err)
	assert.Equal(t, int64(42), *metric.Delta)
}

func TestUpdateMetrics_Success_Gauge(t *testing.T) {
	srv, stor := newTestServer()

	req := pb.UpdateMetricsRequest_builder{
		Metrics: []*pb.Metric{
			pb.Metric_builder{
				Id:    "g1",
				Type:  pb.Metric_GAUGE,
				Value: 3.14,
			}.Build(),
		},
	}.Build()

	resp, err := srv.UpdateMetrics(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, resp)

	metric, err := stor.GetMetric(context.Background(), "g1")
	require.NoError(t, err)
	assert.Equal(t, 3.14, *metric.Value)
}

func TestUpdateMetrics_Success_MixedBatch(t *testing.T) {
	srv, _ := newTestServer()

	req := pb.UpdateMetricsRequest_builder{
		Metrics: []*pb.Metric{
			pb.Metric_builder{Id: "counter-1", Type: pb.Metric_COUNTER, Delta: 10}.Build(),
			pb.Metric_builder{Id: "gauge-1", Type: pb.Metric_GAUGE, Value: 1.5}.Build(),
			pb.Metric_builder{Id: "counter-2", Type: pb.Metric_COUNTER, Delta: 20}.Build(),
			pb.Metric_builder{Id: "gauge-2", Type: pb.Metric_GAUGE, Value: 0.0}.Build(),
		},
	}.Build()

	resp, err := srv.UpdateMetrics(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, resp)
}

func TestUpdateMetrics_EmptyMetrics_NoError(t *testing.T) {
	srv, _ := newTestServer()

	req := pb.UpdateMetricsRequest_builder{}.Build()

	resp, err := srv.UpdateMetrics(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, resp)
}

func TestUpdateMetrics_UnknownMetricType(t *testing.T) {
	srv, _ := newTestServer()

	req := pb.UpdateMetricsRequest_builder{
		Metrics: []*pb.Metric{
			pb.Metric_builder{
				Id:   "bad",
				Type: pb.Metric_MType(99),
			}.Build(),
		},
	}.Build()

	_, err := srv.UpdateMetrics(context.Background(), req)
	assert.Error(t, err)

	st, ok := status.FromError(err)
	require.True(t, ok)
	assert.Equal(t, codes.InvalidArgument, st.Code())
	assert.Contains(t, st.Message(), "unknown metric type")
}

func TestUpdateMetrics_AccumulatesCounter(t *testing.T) {
	srv, stor := newTestServer()

	req1 := pb.UpdateMetricsRequest_builder{
		Metrics: []*pb.Metric{
			pb.Metric_builder{Id: "cpu_total", Type: pb.Metric_COUNTER, Delta: 100}.Build(),
		},
	}.Build()
	_, err := srv.UpdateMetrics(context.Background(), req1)
	require.NoError(t, err)

	req2 := pb.UpdateMetricsRequest_builder{
		Metrics: []*pb.Metric{
			pb.Metric_builder{Id: "cpu_total", Type: pb.Metric_COUNTER, Delta: 50}.Build(),
		},
	}.Build()
	_, err = srv.UpdateMetrics(context.Background(), req2)
	require.NoError(t, err)

	metric, err := stor.GetMetric(context.Background(), "cpu_total")
	require.NoError(t, err)
	assert.Equal(t, int64(150), *metric.Delta)
}

func TestUpdateMetrics_OverwritesGauge(t *testing.T) {
	srv, stor := newTestServer()

	req1 := pb.UpdateMetricsRequest_builder{
		Metrics: []*pb.Metric{
			pb.Metric_builder{Id: "temp", Type: pb.Metric_GAUGE, Value: 100}.Build(),
		},
	}.Build()
	_, err := srv.UpdateMetrics(context.Background(), req1)
	require.NoError(t, err)

	req2 := pb.UpdateMetricsRequest_builder{
		Metrics: []*pb.Metric{
			pb.Metric_builder{Id: "temp", Type: pb.Metric_GAUGE, Value: 200}.Build(),
		},
	}.Build()
	_, err = srv.UpdateMetrics(context.Background(), req2)
	require.NoError(t, err)

	metric, err := stor.GetMetric(context.Background(), "temp")
	require.NoError(t, err)
	assert.Equal(t, 200.0, *metric.Value)
}

func TestUpdateMetrics_NilRequest_MetricsNil(t *testing.T) {
	srv, _ := newTestServer()

	req := &pb.UpdateMetricsRequest{}

	resp, err := srv.UpdateMetrics(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, resp)
}

func BenchmarkUpdateMetrics_Single(b *testing.B) {
	srv, _ := newTestServer()
	ctx := context.Background()
	req := pb.UpdateMetricsRequest_builder{
		Metrics: []*pb.Metric{
			pb.Metric_builder{Id: "bench", Type: pb.Metric_GAUGE, Value: 1.0}.Build(),
		},
	}.Build()

	b.ReportAllocs()
	for b.Loop() {
		_, _ = srv.UpdateMetrics(ctx, req)
	}
}

func BenchmarkUpdateMetrics_Batch10(b *testing.B) {
	srv, _ := newTestServer()
	ctx := context.Background()
	metrics := make([]*pb.Metric, 10)
	for i := 0; i < 10; i++ {
		metrics[i] = &pb.Metric{}
		metrics[i].SetId(string(rune('a' + i)))
		metrics[i].SetType(pb.Metric_GAUGE)
		metrics[i].SetValue(float64(i))
	}
	req := pb.UpdateMetricsRequest_builder{Metrics: metrics}.Build()

	b.ReportAllocs()
	for b.Loop() {
		_, _ = srv.UpdateMetrics(ctx, req)
	}
}
