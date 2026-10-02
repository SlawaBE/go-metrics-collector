package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/SlawaBE/go-metrics-collector/internal/logger"
	"github.com/SlawaBE/go-metrics-collector/internal/model"
	"go.uber.org/zap"
)

type AuditService struct {
	subscribers []AuditSubscriber
}

func NewAuditService() *AuditService {
	return &AuditService{
		subscribers: make([]AuditSubscriber, 0),
	}
}

func (a *AuditService) Subscribe(subscriber AuditSubscriber) {
	a.subscribers = append(a.subscribers, subscriber)
}

func (a *AuditService) SendMetric(ctx context.Context, ip string, metric model.Metric) {
	metricNames := make([]string, 1)
	metricNames[0] = metric.ID
	a.notify(ctx, model.NewAuditEvent(ip, metricNames))
}

func (a *AuditService) SendMetrics(ctx context.Context, ip string, metrics []model.Metric) {
	metricNames := make([]string, len(metrics))
	for i, metric := range metrics {
		metricNames[i] = metric.ID
	}
	a.notify(ctx, model.NewAuditEvent(ip, metricNames))
}

func (a *AuditService) notify(ctx context.Context, event model.AuditEvent) {
	logger.Log.Info("Send audit event", zap.Any("event", event))
	for _, subscriber := range a.subscribers {
		err := subscriber.Notify(ctx, event)
		if err != nil {
			logger.Log.Error("Error send audit event", zap.String("audit_notifier", subscriber.Name()), zap.Error(err))
		}
	}
}

type FileAuditSubscriber struct {
	filePath string
}

func NewFileAuditSubscriber(path string) *FileAuditSubscriber {
	return &FileAuditSubscriber{filePath: path}
}

func (f *FileAuditSubscriber) Name() string {
	return "FileAuditSubscriber"
}

func (f *FileAuditSubscriber) Notify(ctx context.Context, event model.AuditEvent) error {
	file, err := os.OpenFile(f.filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	if err := json.NewEncoder(file).Encode(event); err != nil {
		return fmt.Errorf("failed to encode json to file: %w", err)
	}

	return nil
}

type HttpAuditSubscriber struct {
	url    string
	client *http.Client
}

func NewHttpAuditSubscriber(url string) *HttpAuditSubscriber {
	return &HttpAuditSubscriber{
		url: url,
		client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

func (h *HttpAuditSubscriber) Name() string {
	return "HttpSubscriber"
}

func (h *HttpAuditSubscriber) Notify(ctx context.Context, event model.AuditEvent) error {
	jsonData, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("failed to marshal event: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, h.url, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := h.client.Do(req)
	if err != nil {
		return fmt.Errorf("http request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("unexpected http status: %d", resp.StatusCode)
	}

	return nil
}
