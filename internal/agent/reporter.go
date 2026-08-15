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

type Reporter struct {
	storage        Storage
	reportInterval int
	client         *resty.Client
	secretKey      []byte
	rateLimit      int
}

type Storage interface {
	storage.Storage
	GetValuesAndClear() []model.Metric
}

var retryIntervals = []time.Duration{1 * time.Second, 3 * time.Second, 5 * time.Second}

func NewReporter(storage Storage, reportAddress string, reportInterval int, secretKey string, rateLimit int) *Reporter {
	return &Reporter{
		storage:        storage,
		reportInterval: reportInterval,
		client:         httpClient("http://" + reportAddress),
		secretKey:      []byte(secretKey),
		rateLimit:      rateLimit,
	}
}

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
			jobs <- r.storage.GetValuesAndClear()
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
			if err := r.sendMetrics(metrics); err != nil {
				fmt.Println("Error sending metrics:", err)
			}
		}
	}
}

func (r *Reporter) Report() {
	metrics := r.storage.GetValuesAndClear()
	if len(metrics) == 0 {
		return
	}
	if err := r.sendMetrics(metrics); err != nil {
		fmt.Println("Error sending metrics:", err)
	}
}

func (r *Reporter) sendMetrics(metrics []model.Metric) error {
	jsonData, err := json.Marshal(metrics)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %v", err)
	}

	data, err := gzip.Compress(jsonData)
	if err != nil {
		return fmt.Errorf("error compress metrics: %v", err)
	}

	req := r.client.R()
	if len(r.secretKey) > 0 {
		sign := checksum.Sign(jsonData, r.secretKey)
		req.SetHeaderVerbatim("HashSHA256", hex.EncodeToString(sign))
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

func httpClient(baseUrl string) *resty.Client {
	client := resty.New().
		SetBaseURL(baseUrl).
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
