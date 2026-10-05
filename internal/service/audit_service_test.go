package service

import (
	"context"
	"sync"
	"testing"

	"github.com/SlawaBE/go-metrics-collector/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type mockAuditSubscriber struct {
	mu     sync.Mutex
	events []model.AuditEvent
	name   string
}

func (m *mockAuditSubscriber) Notify(_ context.Context, event model.AuditEvent) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.events = append(m.events, event)
	return nil
}

func (m *mockAuditSubscriber) Name() string { return m.name }

func (m *mockAuditSubscriber) received() []model.AuditEvent {
	m.mu.Lock()
	defer m.mu.Unlock()
	return append([]model.AuditEvent(nil), m.events...)
}

func TestAuditService_SendMetric(t *testing.T) {
	svc := NewAuditService()
	sub := &mockAuditSubscriber{name: "mock"}
	svc.Subscribe(sub)

	svc.SendMetric(context.Background(), "192.168.1.1", model.NewCounterMetric("counter-1", 10))

	events := sub.received()
	require.Len(t, events, 1)
	assert.Equal(t, "192.168.1.1", events[0].IPAddress)
	assert.Equal(t, []string{"counter-1"}, events[0].Metrics)
}

func TestAuditService_SendMetrics(t *testing.T) {
	svc := NewAuditService()
	sub := &mockAuditSubscriber{name: "mock"}
	svc.Subscribe(sub)

	metrics := []model.Metric{
		model.NewCounterMetric("counter-1", 10),
		model.NewGaugeMetric("gauge-1", 1.5),
	}

	svc.SendMetrics(context.Background(), "10.0.0.1", metrics)

	events := sub.received()
	require.Len(t, events, 1)
	assert.Equal(t, "10.0.0.1", events[0].IPAddress)
	assert.Equal(t, []string{"counter-1", "gauge-1"}, events[0].Metrics)
}

func TestAuditService_SubscribeBroadcastsToAll(t *testing.T) {
	svc := NewAuditService()
	first := &mockAuditSubscriber{name: "first"}
	second := &mockAuditSubscriber{name: "second"}
	svc.Subscribe(first)
	svc.Subscribe(second)

	svc.SendMetric(context.Background(), "192.168.1.1", model.NewCounterMetric("c", 1))

	assert.Len(t, first.received(), 1)
	assert.Len(t, second.received(), 1)
}

func TestAuditService_NotifyWithoutSubscribers(t *testing.T) {
	svc := NewAuditService()

	assert.NotPanics(t, func() {
		svc.SendMetric(context.Background(), "192.168.1.1", model.NewCounterMetric("c", 1))
	})
}
