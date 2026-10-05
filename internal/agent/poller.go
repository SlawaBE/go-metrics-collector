package agent

import (
	"context"
	"fmt"
	"time"

	"github.com/SlawaBE/go-metrics-collector/internal/model"
	"github.com/SlawaBE/go-metrics-collector/internal/storage"
)

// Poller periodically collects metrics through the given function and saves
// them to the storage at the configured polling interval.
type Poller struct {
	storage      storage.Storage
	pollInterval int
	getMetrics   func() []model.Metric
}

// NewPoller creates a poller with the given storage, polling interval and
// metric collection function.
func NewPoller(storage storage.Storage, pollInterval int, getMetrics func() []model.Metric) *Poller {
	return &Poller{
		storage:      storage,
		pollInterval: pollInterval,
		getMetrics:   getMetrics,
	}
}

// NewRuntimeMetricsPoller creates a poller for metrics obtained from the
// runtime.
func NewRuntimeMetricsPoller(storage storage.Storage, pollInterval int) *Poller {
	return NewPoller(storage, pollInterval, func() []model.Metric { return GetRuntimeMetrics().convertToList() })
}

// NewGopsutilMetricsPoller creates a poller for metrics obtained via gopsutil
// (memory and CPU).
func NewGopsutilMetricsPoller(storage storage.Storage, pollInterval int) *Poller {
	return NewPoller(storage, pollInterval, func() []model.Metric { return GetGopsutilsMetrics().convertToList() })
}

// Run starts the cyclic metric collection at the polling interval until the
// context is cancelled.
func (p *Poller) Run(ctx context.Context) {
	ticker := time.NewTicker(time.Duration(p.pollInterval) * time.Second)
	for {
		select {
		case <-ctx.Done():
			ticker.Stop()
			return
		case <-ticker.C:
			p.Poll()
		}
	}
}

// Poll collects metrics and saves them to the storage.
func (p *Poller) Poll() {
	metrics := p.getMetrics() //GetRuntimeMetrics().convertToList()
	if err := p.storage.UpdateAll(context.Background(), metrics); err != nil {
		fmt.Println("error update metrics")
	}
}
