package agent

import (
	"time"

	"github.com/SlawaBE/go-metrics-collector/internal/storage"
)

type Poller struct {
	storage      storage.Storage
	pollInterval int
}

func NewPoller(storage storage.Storage, pollInterval int) *Poller {
	return &Poller{
		storage:      storage,
		pollInterval: pollInterval,
	}
}

func (p *Poller) Run() {
	for {
		time.Sleep(time.Duration(p.pollInterval) * time.Second)
		p.Poll()
	}
}

func (p *Poller) Poll() {
	metrics := GetRuntimeMetrics().convertToList()
	for _, m := range metrics {
	    p.storage.UpdateMetric(m)
	}
}
