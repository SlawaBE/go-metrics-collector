package agent

import (
	"context"
	"fmt"
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
	    if err := p.storage.UpdateMetric(context.Background(), m); err != nil {
            fmt.Println("error update metric:", m.ID)
	    }
	}
}
