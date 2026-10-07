package agent

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/SlawaBE/go-metrics-collector/internal/gzip"
	"github.com/SlawaBE/go-metrics-collector/internal/model"
	"github.com/SlawaBE/go-metrics-collector/internal/storage"
	"github.com/SlawaBE/go-metrics-collector/internal/utils/checksum"
	"github.com/go-resty/resty/v2"
)

// Reporter periodically sends the accumulated metrics to the server in a batch
// via POST /updates, using a pool of workers.
type Reporter struct {
	storage        Storage
	reportInterval int
	client         *resty.Client
	secretKey      []byte
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

// NewReporter creates a reporter that sends metrics to the given address at the
// given interval, signature key and worker limit.
func NewReporter(storage Storage, reportAddress string, reportInterval int, secretKey string, rateLimit int) *Reporter {
	return &Reporter{
		storage:        storage,
		reportInterval: reportInterval,
		client:         httpClient("http://" + reportAddress),
		secretKey:      []byte(secretKey),
		rateLimit:      rateLimit,
	}
}

// Run starts the metric sending loop to the server until the context is
// cancelled.
func (r *Reporter) Run(ctx context.Context) {
	jobs := make(chan []model.Metric, r.rateLimit)
	var wg sync.WaitGroup

	for range r.rateLimit {
		wg.Go(func() {
			r.work(ctx, jobs)
		})
	}

	ticker := time.NewTicker(time.Duration(r.reportInterval) * time.Second)
	for {
		select {
		case <-ctx.Done():
			ticker.Stop()
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

func (r *Reporter) work(ctx context.Context, jobs <-chan []model.Metric) {
	for {
		select {
		case <-ctx.Done():
			return
		case metrics := <-jobs:
			if len(metrics) == 0 {
				continue
			}
			if err := r.sendMetrics(ctx, metrics); err != nil {
				fmt.Println("Error sending metrics:", err)
			}
		}
	}
}

// Report consumes the accumulated metrics and sends them to the server once.
func (r *Reporter) Report() {
	metrics := r.storage.GetValuesAndClear()
	if len(metrics) == 0 {
		return
	}
	if err := r.sendMetrics(context.Background(), metrics); err != nil {
		fmt.Println("Error sending metrics:", err)
	}
}

func (r *Reporter) sendMetrics(ctx context.Context, metrics []model.Metric) error {
	jsonData, err := json.Marshal(metrics)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %v", err)
	}

	data, err := gzip.Compress(jsonData)
	if err != nil {
		return fmt.Errorf("error compress metrics: %v", err)
	}

	req := r.client.R().
		SetContext(ctx)
	if len(r.secretKey) > 0 {
		sign := checksum.Sign(jsonData, r.secretKey)
		req.SetHeader("HashSHA256", hex.EncodeToString(sign))
	}

	res, err := req.SetBody(data).
		Post("/updates")

	if err != nil {
		return fmt.Errorf("error sending metrics: %v", err)
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
