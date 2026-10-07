package service

import (
	"fmt"
	"sync"

	"github.com/SlawaBE/go-metrics-collector/internal/model"
)

// memoryAuditSubscriber collects delivered audit events in a thread-safe slice
// so that an example can assert on them deterministically.
type memoryAuditSubscriber struct {
	mu     sync.Mutex
	events []model.AuditEvent
}

func (s *memoryAuditSubscriber) Name() string { return "memoryAuditSubscriber" }
func (s *memoryAuditSubscriber) Close() error { return nil }

func (s *memoryAuditSubscriber) Notify(event model.AuditEvent) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.events = append(s.events, event)
	return nil
}

func (s *memoryAuditSubscriber) collected() []model.AuditEvent {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.events
}

// ExampleAuditService demonstrates publishing audit events to a subscriber and
// draining them with Shutdown. Shutdown waits for all queued events to be
// delivered, so the example is deterministic.
func ExampleAuditService() {
	service, err := NewAuditService(10)
	if err != nil {
		panic(err)
	}

	sub := &memoryAuditSubscriber{}
	service.Subscribe(sub)

	service.SendMetric("192.168.1.10:54321", model.NewCounterMetric("PollCount", 1))
	service.SendMetrics("10.0.0.5:80", []model.Metric{
		model.NewGaugeMetric("HeapAlloc", 1024.5),
		model.NewGaugeMetric("CPUUtilization1", 37.5),
	})

	service.Shutdown()

	fmt.Printf("delivered events: %d\n", len(sub.collected()))
	for _, e := range sub.collected() {
		fmt.Printf("ip=%s metrics=%v\n", e.IPAddress, e.Metrics)
	}

	// Output:
	// delivered events: 2
	// ip=192.168.1.10 metrics=[PollCount]
	// ip=10.0.0.5 metrics=[HeapAlloc CPUUtilization1]
}
