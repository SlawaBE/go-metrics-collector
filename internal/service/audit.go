package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net"
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
	a.notify(ctx, model.NewAuditEvent(ip, []string{metric.ID}))
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

type HTTPAuditSubscriber struct {
	url    string
	client *http.Client
}

func auditTransport() *http.Transport {
	return &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   3 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          10,
		MaxIdleConnsPerHost:   10,
		MaxConnsPerHost:       10,
		IdleConnTimeout:       30 * time.Second,
		TLSHandshakeTimeout:   3 * time.Second,
		ResponseHeaderTimeout: 5 * time.Second,
		ExpectContinueTimeout: time.Second,
	}
}

func NewHTTPAuditSubscriber(url string) *HTTPAuditSubscriber {
	return &HTTPAuditSubscriber{
		url: url,
		client: &http.Client{
			Timeout:   5 * time.Second,
			Transport: auditTransport(),
		},
	}
}

func (h *HTTPAuditSubscriber) Name() string {
	return "HttpSubscriber"
}

func (h *HTTPAuditSubscriber) Notify(ctx context.Context, event model.AuditEvent) error {
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
