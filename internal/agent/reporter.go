package agent

import (
	"context"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/SlawaBE/go-metrics-collector/internal/asymcrypto"
	"github.com/SlawaBE/go-metrics-collector/internal/gzip"
	"github.com/SlawaBE/go-metrics-collector/internal/logger"
	"github.com/SlawaBE/go-metrics-collector/internal/model"
	pb "github.com/SlawaBE/go-metrics-collector/internal/proto/metrics"
	"github.com/SlawaBE/go-metrics-collector/internal/storage"
	"github.com/SlawaBE/go-metrics-collector/internal/utils/checksum"
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

type Reporter interface {
	Run(ctx context.Context)
}

// HTTPReporter periodically sends the accumulated metrics to the server in a batch
// via POST /updates, using a pool of workers.
type HTTPReporter struct {
	storage        Storage
	rsaEncryptor   *asymcrypto.RSAEncryptor
	client         *resty.Client
	clientIP       string
	secretKey      []byte
	reportInterval time.Duration
	rateLimit      int
}

// Storage defines the reporter's requirement on the storage: in addition to the
// regular Storage operations, GetValuesAndClear is required to consume metrics
// between reports.
type Storage interface {
	storage.Storage

	// GetValuesAndClear returns all saved metrics and cleans the storage.
	GetValuesAndClear() []model.Metric
}

// retryIntervals holds increasing send retry intervals (in seconds).
var retryIntervals = []time.Duration{1 * time.Second, 3 * time.Second, 5 * time.Second}

// NewHTTPReporter creates a reporter that sends metrics to the given address at the
// given interval, signature key and worker limit.
func NewHTTPReporter(storage Storage, rsaEncryptor *asymcrypto.RSAEncryptor, reportAddress string, reportInterval time.Duration, secretKey string, rateLimit int, clientIP string) *HTTPReporter {
	return &HTTPReporter{
		storage:        storage,
		rsaEncryptor:   rsaEncryptor,
		reportInterval: reportInterval,
		client:         httpClient("http://" + reportAddress),
		secretKey:      []byte(secretKey),
		rateLimit:      rateLimit,
		clientIP:       clientIP,
	}
}

// Run starts the metric sending loop to the server until the context is
// cancelled.
func (r *HTTPReporter) Run(ctx context.Context) {
	jobs := make(chan []model.Metric, r.rateLimit)
	var wg sync.WaitGroup

	for range r.rateLimit {
		wg.Go(func() {
			r.work(ctx, jobs)
		})
	}

	ticker := time.NewTicker(r.reportInterval)
	for {
		select {
		case <-ctx.Done():
			ticker.Stop()
			close(jobs)
			wg.Wait()
			return
		case <-ticker.C:
			select {
			case <-ctx.Done():
			case jobs <- r.storage.GetValuesAndClear():
			}
		}
	}
}

func (r *HTTPReporter) work(ctx context.Context, jobs <-chan []model.Metric) {
	for metrics := range jobs {
		if len(metrics) == 0 {
			continue
		}
		if err := r.sendMetrics(ctx, metrics); err != nil {
			fmt.Println("Error sending metrics:", err)
		}
	}
}

// Report consumes the accumulated metrics and sends them to the server once.
func (r *HTTPReporter) Report() {
	metrics := r.storage.GetValuesAndClear()
	if len(metrics) == 0 {
		return
	}
	if err := r.sendMetrics(context.Background(), metrics); err != nil {
		fmt.Println("Error sending metrics:", err)
	}
}

func (r *HTTPReporter) sendMetrics(ctx context.Context, metrics []model.Metric) error {
	jsonData, err := json.Marshal(metrics)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	if r.clientIP != "" {
		r.client.SetHeader("X-Real-IP", r.clientIP)
	}

	data, err := gzip.Compress(jsonData)
	if err != nil {
		return fmt.Errorf("error compress metrics: %w", err)
	}

	req := r.client.R().
		SetContext(ctx)
	if len(r.secretKey) > 0 {
		sign := checksum.Sign(jsonData, r.secretKey)
		req.SetHeader("HashSHA256", hex.EncodeToString(sign))
	}

	if r.rsaEncryptor != nil {
		hct, encErr := r.rsaEncryptor.Encrypt(data)
		if encErr != nil {
			return fmt.Errorf("error encrypting metrics: %w", encErr)
		}
		data = hct.CipherText
		req.SetHeader("X-Nonce", base64.StdEncoding.EncodeToString(hct.Nonce))
		req.SetHeader("X-Encrypted-Key", base64.StdEncoding.EncodeToString(hct.EncryptedKey))
	}

	res, err := req.SetBody(data).
		Post("/updates")

	if err != nil {
		return fmt.Errorf("error sending metrics: %w", err)
	}

	if res.StatusCode() != http.StatusOK {
		return fmt.Errorf("error in server response: %d", res.StatusCode())
	}

	return nil
}

func httpClient(baseURL string) *resty.Client {
	client := resty.New().
		SetBaseURL(baseURL).
		SetHeader("Content-Type", "application/json").
		SetHeader("Content-Encoding", "gzip").
		SetHeader("Accept-Encoding", "gzip").
		SetRetryCount(len(retryIntervals)).
		SetRetryAfter(func(c *resty.Client, r *resty.Response) (time.Duration, error) {
			return retryIntervals[r.Request.Attempt-1], nil
		}).
		SetRetryMaxWaitTime(retryIntervals[len(retryIntervals)-1])
	return client
}

// GRPCReporter periodically sends the accumulated metrics to the server in a batch
// via Metrics.UpdateMetrics, using a pool of workers.
type GRPCReporter struct {
	storage        Storage
	client         pb.MetricsClient
	clientIP       string
	reportInterval time.Duration
	rateLimit      int
}

func NewGRPCReporter(storage Storage, clientIP string, client pb.MetricsClient, reportInterval time.Duration, rateLimit int) *GRPCReporter {
	return &GRPCReporter{
		storage:        storage,
		clientIP:       clientIP,
		client:         client,
		reportInterval: reportInterval,
		rateLimit:      rateLimit,
	}
}

// Run starts the metric sending loop to the server until the context is
// cancelled.
func (r *GRPCReporter) Run(ctx context.Context) {
	jobs := make(chan []model.Metric, r.rateLimit)
	var wg sync.WaitGroup

	for range r.rateLimit {
		wg.Go(func() {
			r.work(ctx, jobs)
		})
	}

	ticker := time.NewTicker(r.reportInterval)
	for {
		select {
		case <-ctx.Done():
			ticker.Stop()
			close(jobs)
			wg.Wait()
			return
		case <-ticker.C:
			select {
			case <-ctx.Done():
			case jobs <- r.storage.GetValuesAndClear():
			}
		}
	}
}

func (r *GRPCReporter) work(ctx context.Context, jobs <-chan []model.Metric) {
	for metrics := range jobs {
		if len(metrics) == 0 {
			continue
		}
		if err := r.sendMetrics(ctx, metrics); err != nil {
			logger.Log.Error("error sending metrics", zap.Error(err))
		}
	}
}

func (r *GRPCReporter) sendMetrics(ctx context.Context, metrics []model.Metric) error {
	grpcMetrics := make([]*pb.Metric, len(metrics))
	for i, m := range metrics {
		switch m.MType {
		case model.Counter:
			grpcMetrics[i] = pb.Metric_builder{
				Id:    m.ID,
				Type:  pb.Metric_COUNTER,
				Delta: *m.Delta,
			}.Build()
		case model.Gauge:
			grpcMetrics[i] = pb.Metric_builder{
				Id:    m.ID,
				Type:  pb.Metric_GAUGE,
				Value: *m.Value,
			}.Build()
		default:
			return fmt.Errorf("unknown metric type %s", m.MType)
		}
	}
	req := pb.UpdateMetricsRequest_builder{
		Metrics: grpcMetrics,
	}.Build()
	if r.clientIP != "" {
		md := metadata.New(map[string]string{"x-real-ip": r.clientIP})
		ctx = metadata.NewOutgoingContext(ctx, md)
	}
	_, err := r.client.UpdateMetrics(ctx, req)
	if err != nil {
		return fmt.Errorf("error from grpc: %w", err)
	}
	return nil
}
