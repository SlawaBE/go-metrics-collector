// Package grpchandler contains GRPC handlers for pdating metrics
package grpchandler

import (
	"context"

	"github.com/SlawaBE/go-metrics-collector/internal/logger"
	"github.com/SlawaBE/go-metrics-collector/internal/model"
	pb "github.com/SlawaBE/go-metrics-collector/internal/proto/metrics"
	"github.com/SlawaBE/go-metrics-collector/internal/service"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MetricsServer struct {
	pb.UnimplementedMetricsServer
	metricsService *service.MetricsService
}

func NewMetricsServer(service *service.MetricsService) pb.MetricsServer {
	return &MetricsServer{
		metricsService: service,
	}
}

func (s *MetricsServer) UpdateMetrics(ctx context.Context, req *pb.UpdateMetricsRequest) (*pb.UpdateMetricsResponse, error) {
	grpcMetrics := req.GetMetrics()
	metrics := make([]model.Metric, len(grpcMetrics))
	for i, m := range grpcMetrics {
		switch m.GetType() {
		case pb.Metric_COUNTER:
			metrics[i] = model.NewCounterMetric(m.GetId(), m.GetDelta())
		case pb.Metric_GAUGE:
			metrics[i] = model.NewGaugeMetric(m.GetId(), m.GetValue())
		default:
			return nil, status.Errorf(codes.InvalidArgument, "unknown metric type value=%d", m.GetType())
		}
	}

	err := s.metricsService.UpdateMetrics(ctx, metrics)
	if err != nil {
		logger.Log.Error("error updating metrics", zap.Error(err))
		return nil, status.Error(codes.Internal, "internal server error")
	}
	return &pb.UpdateMetricsResponse{}, nil
}
