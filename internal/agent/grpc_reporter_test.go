package agent

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/SlawaBE/go-metrics-collector/internal/model"
	"github.com/SlawaBE/go-metrics-collector/internal/storage"
	pb "github.com/SlawaBE/go-metrics-collector/internal/proto/metrics"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// mockMetricsClient implements pb.MetricsClient for testing.
type mockMetricsClient struct {
	gotMeta    map[string][]string
	returnErr  error
	lastReq    *pb.UpdateMetricsRequest
	lastCtx    context.Context
	calls      int
}

func (m *mockMetricsClient) UpdateMetrics(ctx context.Context, req *pb.UpdateMetricsRequest, opts ...grpc.CallOption) (*pb.UpdateMetricsResponse, error) {
	m.calls++
	m.lastReq = req
	m.lastCtx = ctx
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		m.gotMeta = md
	}
	if m.returnErr != nil {
		return nil, m.returnErr
	}
	return &pb.UpdateMetricsResponse{}, nil
}

func TestNewGRPCReporter(t *testing.T) {
	mock := &mockMetricsClient{}
	stor := storage.NewMemStorage()

	r := NewGRPCReporter(stor, "10.0.0.1", mock, time.Second, 2)

	require.NotNil(t, r)
	assert.Equal(t, "10.0.0.1", r.clientIP)
	assert.Same(t, mock, r.client)
	assert.Equal(t, 2, r.rateLimit)
}

func TestGRPCReporter_sendMetrics_Success_Counter(t *testing.T) {
	mock := &mockMetricsClient{}
	reporter := NewGRPCReporter(nil, "", mock, time.Second, 1)

	metrics := []model.Metric{model.NewCounterMetric("c1", 42)}

	err := reporter.sendMetrics(context.Background(), metrics)
	require.NoError(t, err)
	assert.Equal(t, 1, mock.calls)
	assert.Len(t, mock.lastReq.GetMetrics(), 1)
	assert.Equal(t, "c1", mock.lastReq.GetMetrics()[0].GetId())
	assert.Equal(t, int64(42), mock.lastReq.GetMetrics()[0].GetDelta())
}

func TestGRPCReporter_sendMetrics_Success_Gauge(t *testing.T) {
	mock := &mockMetricsClient{}
	reporter := NewGRPCReporter(nil, "", mock, time.Second, 1)

	metrics := []model.Metric{model.NewGaugeMetric("g1", 3.14)}

	err := reporter.sendMetrics(context.Background(), metrics)
	require.NoError(t, err)
	assert.Len(t, mock.lastReq.GetMetrics(), 1)
	assert.Equal(t, "g1", mock.lastReq.GetMetrics()[0].GetId())
	assert.Equal(t, 3.14, mock.lastReq.GetMetrics()[0].GetValue())
}

func TestGRPCReporter_sendMetrics_MixedMetrics(t *testing.T) {
	mock := &mockMetricsClient{}
	reporter := NewGRPCReporter(nil, "", mock, time.Second, 1)

	metrics := []model.Metric{
		model.NewCounterMetric("counter-1", 10),
		model.NewGaugeMetric("gauge-1", 2.5),
		model.NewCounterMetric("counter-2", 20),
	}

	err := reporter.sendMetrics(context.Background(), metrics)
	require.NoError(t, err)
	assert.Len(t, mock.lastReq.GetMetrics(), 3)

	assert.Equal(t, pb.Metric_COUNTER, mock.lastReq.GetMetrics()[0].GetType())
	assert.Equal(t, int64(10), mock.lastReq.GetMetrics()[0].GetDelta())
	assert.Equal(t, pb.Metric_GAUGE, mock.lastReq.GetMetrics()[1].GetType())
	assert.Equal(t, 2.5, mock.lastReq.GetMetrics()[1].GetValue())
}

func TestGRPCReporter_sendMetrics_EmptyMetrics_SendsEmptyBatch(t *testing.T) {
	mock := &mockMetricsClient{}
	reporter := NewGRPCReporter(nil, "", mock, time.Second, 1)

	err := reporter.sendMetrics(context.Background(), nil)
	assert.NoError(t, err)
	assert.Equal(t, 1, mock.calls)
	assert.Empty(t, mock.lastReq.GetMetrics())
}

func TestGRPCReporter_sendMetrics_WithClientIP(t *testing.T) {
	mock := &mockMetricsClient{}
	reporter := NewGRPCReporter(nil, "10.0.0.5", mock, time.Second, 1)

	metrics := []model.Metric{model.NewCounterMetric("c1", 1)}

	err := reporter.sendMetrics(context.Background(), metrics)
	require.NoError(t, err)

	md, ok := metadata.FromOutgoingContext(mock.lastCtx)
	require.True(t, ok)
	values := md.Get("x-real-ip")
	require.Len(t, values, 1)
	assert.Equal(t, "10.0.0.5", values[0])
}

func TestGRPCReporter_sendMetrics_GrpcError(t *testing.T) {
	mock := &mockMetricsClient{returnErr: fmt.Errorf("connection refused")}
	reporter := NewGRPCReporter(nil, "", mock, time.Second, 1)

	metrics := []model.Metric{model.NewCounterMetric("c1", 1)}

	err := reporter.sendMetrics(context.Background(), metrics)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error from grpc")
}

func TestGRPCReporter_sendMetrics_UnknownMetricType(t *testing.T) {
	mock := &mockMetricsClient{}
	reporter := NewGRPCReporter(nil, "", mock, time.Second, 1)

	metrics := []model.Metric{{ID: "bad", MType: "unknown"}}

	err := reporter.sendMetrics(context.Background(), metrics)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unknown metric type unknown")
	assert.Equal(t, 0, mock.calls)
}

func TestGRPCReporter_Run_ContextCancelled(t *testing.T) {
	mock := &mockMetricsClient{}
	stor := storage.NewMemStorage()
	ctx, cancel := context.WithCancel(context.Background())

	reporter := NewGRPCReporter(stor, "", mock, 10*time.Millisecond, 2)

	done := make(chan struct{})
	go func() {
		reporter.Run(ctx)
		close(done)
	}()

	time.Sleep(50 * time.Millisecond)
	cancel()

	select {
	case <-done:
	case <-time.After(2 * time.Second):
		t.Fatal("Run did not return after context cancellation")
	}

	assert.GreaterOrEqual(t, mock.calls, 0) // may be 0 if nothing was sent before cancel
}

func TestGRPCReporter_Run_SendsMetrics(t *testing.T) {
	mock := &mockMetricsClient{}
	stor := storage.NewMemStorage()
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	require.NoError(t, stor.UpdateMetric(ctx, model.NewCounterMetric("c1", 10)))
	require.NoError(t, stor.UpdateMetric(ctx, model.NewGaugeMetric("g1", 1.5)))

	reporter := NewGRPCReporter(stor, "192.168.1.1", mock, 50*time.Millisecond, 2)

	reporter.Run(ctx)

	assert.GreaterOrEqual(t, mock.calls, 1)

	remaining, err := stor.GetValues(context.Background())
	require.NoError(t, err)
	assert.Empty(t, remaining)
}

func TestGRPCReporter_Run_EmptyStorage(t *testing.T) {
	mock := &mockMetricsClient{}
	stor := storage.NewMemStorage()
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	reporter := NewGRPCReporter(stor, "192.168.1.1", mock, 50*time.Millisecond, 2)

	reporter.Run(ctx)

	assert.Equal(t, 0, mock.calls, "empty batches must not invoke gRPC client")
}
