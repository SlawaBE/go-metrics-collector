package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/SlawaBE/go-metrics-collector/internal/logger"
	"github.com/SlawaBE/go-metrics-collector/internal/model"
	"github.com/SlawaBE/go-metrics-collector/internal/utils"
	"go.uber.org/zap"
)

// AuditService asynchronously delivers audit events to registered subscribers.
//
// All subscribers must be registered before the first event is published.
// Events are buffered in an internal queue. If the queue is full, new events
// are dropped.
//
// Shutdown drains all queued events before closing the registered subscribers.
type AuditService struct {
	subscribers []AuditSubscriber
	events      chan model.AuditEvent

	wg        sync.WaitGroup
	mu        sync.RWMutex
	closed    bool
	closeOnce sync.Once
}

// NewAuditService creates an AuditService with the specified event queue size.
//
// queueSize must be greater than zero.
func NewAuditService(queueSize int) (*AuditService, error) {
	if queueSize <= 0 {
		return nil, errors.New("audit queue size must be greater than zero")
	}

	service := &AuditService{
		subscribers: make([]AuditSubscriber, 0),
		events:      make(chan model.AuditEvent, queueSize),
	}

	service.wg.Add(1)
	go service.run()

	return service, nil
}

// Subscribe registers a subscriber for audit events.
//
// All subscribers must be registered before the first event is published.
func (a *AuditService) Subscribe(subscriber AuditSubscriber) {
	if subscriber == nil {
		return
	}

	a.mu.Lock()
	defer a.mu.Unlock()

	if a.closed {
		return
	}

	a.subscribers = append(a.subscribers, subscriber)
}

// SendMetric publishes an audit event for a single metric.
//
// The event is dropped if the audit queue is full or the service has already
// been shut down.
func (a *AuditService) SendMetric(address string, metric model.Metric) {
	ip := a.extractIP(address)

	a.dispatch(model.NewAuditEvent(
		ip,
		[]string{metric.ID},
	))
}

// SendMetrics publishes an audit event for multiple metrics.
//
// The event is dropped if the audit queue is full or the service has already
// been shut down.
func (a *AuditService) SendMetrics(address string, metrics []model.Metric) {
	ip := a.extractIP(address)

	metricNames := make([]string, len(metrics))
	for i, metric := range metrics {
		metricNames[i] = metric.ID
	}

	a.dispatch(model.NewAuditEvent(ip, metricNames))
}

// Shutdown stops accepting new audit events, drains the event queue, and
// closes all registered subscribers.
//
// Shutdown is safe to call multiple times.
func (a *AuditService) Shutdown() {
	a.closeOnce.Do(func() {
		a.mu.Lock()
		a.closed = true
		close(a.events)
		a.mu.Unlock()

		a.wg.Wait()
	})
}

func (a *AuditService) dispatch(event model.AuditEvent) {
	a.mu.RLock()
	defer a.mu.RUnlock()

	if a.closed {
		logger.Log.Warn(
			"Audit service is shut down, dropping event",
			zap.Any("event", event),
		)
		return
	}

	select {
	case a.events <- event:
	default:
		logger.Log.Warn(
			"Audit queue is full, dropping event",
			zap.Any("event", event),
		)
	}
}

func (a *AuditService) run() {
	defer a.wg.Done()

	for event := range a.events {
		a.notifySubscribers(event)
	}

	a.closeSubscribers()
}

func (a *AuditService) notifySubscribers(event model.AuditEvent) {
	logger.Log.Info(
		"Sending audit event",
		zap.Any("event", event),
	)

	for _, subscriber := range a.subscribers {
		if err := subscriber.Notify(event); err != nil {
			logger.Log.Error(
				"Failed to notify audit subscriber",
				zap.String("audit_subscriber", subscriber.Name()),
				zap.Error(err),
			)
		}
	}
}

func (a *AuditService) closeSubscribers() {
	for _, subscriber := range a.subscribers {
		if err := subscriber.Close(); err != nil {
			logger.Log.Error(
				"Failed to close audit subscriber",
				zap.String("audit_subscriber", subscriber.Name()),
				zap.Error(err),
			)
		}
	}
}

func (a *AuditService) extractIP(address string) string {
	ip, err := utils.GetIPFromAddress(address)
	if err != nil {
		logger.Log.Warn(
			"Failed to extract IP from remote address",
			zap.String("address", address),
			zap.Error(err),
		)
		return ""
	}

	return ip
}

// FileAuditSubscriber writes audit events to a file in JSON format. The file is
// opened once at construction and appended to on every Notify call until Close
// is invoked.
type FileAuditSubscriber struct {
	file *os.File
}

// NewFileAuditSubscriber creates a subscriber that writes audit events to the
// given file, opening the file in append mode. The returned subscriber must be
// closed via Close when it is no longer needed.
func NewFileAuditSubscriber(path string) (*FileAuditSubscriber, error) {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	return &FileAuditSubscriber{file: file}, nil
}

// Name returns the subscriber name for logging.
func (f *FileAuditSubscriber) Name() string {
	return "FileAuditSubscriber"
}

// Close flushes and closes the underlying file. It must be called when the
// subscriber is no longer needed to ensure all events are written to disk.
func (f *FileAuditSubscriber) Close() error {
	if err := f.file.Sync(); err != nil {
		return fmt.Errorf("failed to sync audit file: %w", err)
	}
	if err := f.file.Close(); err != nil {
		return fmt.Errorf("failed to close audit file: %w", err)
	}
	return nil
}

// Notify appends an audit event to the file in JSON format.
func (f *FileAuditSubscriber) Notify(event model.AuditEvent) error {
	if err := json.NewEncoder(f.file).Encode(event); err != nil {
		return fmt.Errorf("failed to encode json to file: %w", err)
	}

	return nil
}

// HTTPAuditSubscriber sends audit events to a remote HTTP endpoint.
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

// NewHTTPAuditSubscriber creates a subscriber that sends audit events to the
// given URL via a POST request.
func NewHTTPAuditSubscriber(url string) *HTTPAuditSubscriber {
	return &HTTPAuditSubscriber{
		url: url,
		client: &http.Client{
			Timeout:   5 * time.Second,
			Transport: auditTransport(),
		},
	}
}

// Name returns the subscriber name for logging.
func (h *HTTPAuditSubscriber) Name() string {
	return "HttpSubscriber"
}

// Close closes idle connections held by the underlying HTTP client.
func (h *HTTPAuditSubscriber) Close() error {
	h.client.CloseIdleConnections()
	return nil
}

// Notify sends an audit event to the remote HTTP endpoint.
func (h *HTTPAuditSubscriber) Notify(event model.AuditEvent) error {
	jsonData, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("failed to marshal event: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, h.url, bytes.NewBuffer(jsonData))
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
