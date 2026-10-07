package service

import (
	"errors"
	"sync"
	"testing"
	"time"

	"github.com/SlawaBE/go-metrics-collector/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type mockAuditSubscriber struct {
	name string

	mu           sync.Mutex
	events       []model.AuditEvent
	closeCounter int

	notifyErr error

	entered   chan struct{}
	block     <-chan struct{}
	enterOnce sync.Once
}

func (m *mockAuditSubscriber) Notify(event model.AuditEvent) error {
	if m.entered != nil {
		m.enterOnce.Do(func() {
			close(m.entered)
		})
	}

	if m.block != nil {
		<-m.block
	}

	m.mu.Lock()
	m.events = append(m.events, event)
	m.mu.Unlock()

	return m.notifyErr
}

func (m *mockAuditSubscriber) Name() string { return m.name }

func (m *mockAuditSubscriber) Close() error {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.closeCounter++
	return nil
}

func (m *mockAuditSubscriber) received() []model.AuditEvent {
	m.mu.Lock()
	defer m.mu.Unlock()
	return append([]model.AuditEvent(nil), m.events...)
}

func (m *mockAuditSubscriber) closeCount() int {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.closeCounter
}

func TestNewAuditService_InvalidQueueSize(t *testing.T) {
	tests := []struct {
		name      string
		queueSize int
	}{
		{
			name:      "zero",
			queueSize: 0,
		},
		{
			name:      "negative",
			queueSize: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service, err := NewAuditService(tt.queueSize)

			if err == nil /*  && tt.queueSize <= 0 */ {
				t.Fatal("expected error, got nil")
			}

			if service != nil /*  && tt.queueSize <= 0 */ {
				t.Fatalf("expected nil service, got %#v", service)
			}
		})
	}
}

func TestAuditService_SendMetric(t *testing.T) {
	svc, err := NewAuditService(10)
	if err != nil {
		t.Fatalf("NewAuditService() error = %v", err)
	}
	sub := &mockAuditSubscriber{name: "mock"}
	svc.Subscribe(sub)

	svc.SendMetric("192.168.1.1:12345", model.NewCounterMetric("counter-1", 10))
	svc.Shutdown()

	events := sub.received()
	require.Len(t, events, 1)
	assert.Equal(t, "192.168.1.1", events[0].IPAddress)
	assert.Equal(t, []string{"counter-1"}, events[0].Metrics)
}

func TestAuditService_SendMetrics(t *testing.T) {
	svc, err := NewAuditService(10)
	if err != nil {
		t.Fatalf("NewAuditService() error = %v", err)
	}
	sub := &mockAuditSubscriber{name: "mock"}
	svc.Subscribe(sub)

	metrics := []model.Metric{
		model.NewCounterMetric("counter-1", 10),
		model.NewGaugeMetric("gauge-1", 1.5),
	}

	svc.SendMetrics("10.0.0.1:12345", metrics)
	svc.Shutdown()

	events := sub.received()
	require.Len(t, events, 1)
	assert.Equal(t, "10.0.0.1", events[0].IPAddress)
	assert.Equal(t, []string{"counter-1", "gauge-1"}, events[0].Metrics)
}

func TestAuditService_NotifiesAllSubscribers(t *testing.T) {
	svc, err := NewAuditService(10)
	if err != nil {
		t.Fatalf("NewAuditService() error = %v", err)
	}
	first := &mockAuditSubscriber{name: "first"}
	second := &mockAuditSubscriber{name: "second"}
	svc.Subscribe(first)
	svc.Subscribe(second)

	svc.SendMetric("192.168.1.1:12345", model.NewCounterMetric("counter-1", 1))
	svc.Shutdown()

	assert.Len(t, first.received(), 1)
	assert.Len(t, second.received(), 1)
}

func TestAuditService_NotifyWithoutSubscribers(t *testing.T) {
	svc, err := NewAuditService(10)
	if err != nil {
		t.Fatalf("NewAuditService() error = %v", err)
	}

	assert.NotPanics(t, func() {
		svc.SendMetric("192.168.1.1:12345", model.NewCounterMetric("counter-1", 1))
	})
	svc.Shutdown()
}

func TestAuditService_SendAfterShutdownIsDropped(t *testing.T) {
	svc, err := NewAuditService(10)
	if err != nil {
		t.Fatalf("NewAuditService() error = %v", err)
	}
	sub := &mockAuditSubscriber{name: "mock"}
	svc.Subscribe(sub)

	svc.SendMetric("192.168.1.1:12345", model.NewCounterMetric("counter-1", 1))
	svc.Shutdown()

	svc.SendMetric("192.168.1.1:12345", model.NewCounterMetric("counter-2", 2))
	svc.Shutdown()

	assert.Len(t, sub.received(), 1)
}

func TestAuditService_ShutdownIsIdempotent(t *testing.T) {
	service, err := NewAuditService(10)
	if err != nil {
		t.Fatalf("NewAuditService() error = %v", err)
	}

	subscriber := &mockAuditSubscriber{
		name: "mock",
	}

	service.Subscribe(subscriber)

	service.Shutdown()
	service.Shutdown()
	service.Shutdown()

	assert.Equal(t, 1, subscriber.closeCount())
}

func TestAuditService_DropsEventWhenQueueIsFull(t *testing.T) {
	service, err := NewAuditService(1)
	if err != nil {
		t.Fatalf("NewAuditService() error = %v", err)
	}

	entered := make(chan struct{})
	release := make(chan struct{})

	subscriber := &mockAuditSubscriber{
		name:    "blocking",
		entered: entered,
		block:   release,
	}

	service.Subscribe(subscriber)

	service.SendMetric(
		"127.0.0.1:8080",
		model.Metric{ID: "first"},
	)

	select {
	case <-entered:
	case <-time.After(time.Second):
		t.Fatal("subscriber did not receive first event")
	}

	service.SendMetric(
		"127.0.0.1:8080",
		model.Metric{ID: "second"},
	)

	service.SendMetric(
		"127.0.0.1:8080",
		model.Metric{ID: "dropped"},
	)

	close(release)

	service.Shutdown()

	events := subscriber.received()

	assert.Len(t, events, 2)

	if got := events[0].Metrics[0]; got != "first" {
		t.Errorf("expected first event, got %q", got)
	}

	if got := events[1].Metrics[0]; got != "second" {
		t.Errorf("expected second event, got %q", got)
	}

	for _, event := range events {
		if event.Metrics[0] == "dropped" {
			t.Fatal("expected third event to be dropped")
		}
	}
}

func TestAuditService_NotifyErrorDoesNotStopOtherSubscribers(t *testing.T) {
	service, err := NewAuditService(10)
	if err != nil {
		t.Fatalf("NewAuditService() error = %v", err)
	}

	failing := &mockAuditSubscriber{
		name:      "failing",
		notifyErr: errors.New("notify error"),
	}

	working := &mockAuditSubscriber{
		name: "working",
	}

	service.Subscribe(failing)
	service.Subscribe(working)

	service.SendMetric(
		"127.0.0.1:8080",
		model.Metric{ID: "metric"},
	)

	service.Shutdown()
	assert.Len(t, working.received(), 1)
	assert.Len(t, failing.received(), 1)
	assert.Equal(t, 1, working.closeCount())
	assert.Equal(t, 1, failing.closeCount())
}
