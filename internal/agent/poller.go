package agent

import (
	"context"
	"fmt"
	"time"

	"github.com/SlawaBE/go-metrics-collector/internal/model"
	"github.com/SlawaBE/go-metrics-collector/internal/storage"
)

type Poller struct {
	storage      storage.Storage
	pollInterval int
	getMetrics   func() []model.Metric
}

func NewPoller(storage storage.Storage, pollInterval int, getMetrics func() []model.Metric) *Poller {
	return &Poller{
		storage:      storage,
		pollInterval: pollInterval,
		getMetrics:   getMetrics,
	}
}

func NewRuntimeMetricsPoller(storage storage.Storage, pollInterval int) *Poller {
	return NewPoller(storage, pollInterval, func() []model.Metric { return GetRuntimeMetrics().convertToList() })
}

func NewGopsutilMetricsPoller(storage storage.Storage, pollInterval int) *Poller {
	return NewPoller(storage, pollInterval, func() []model.Metric { return GetGopsutilsMetrics().convertToList() })
}

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

func (p *Poller) Poll() {
	metrics := p.getMetrics() //GetRuntimeMetrics().convertToList()
	if err := p.storage.UpdateAll(context.Background(), metrics); err != nil {
		fmt.Println("error update metrics")
	}
}
