package agent

import "github.com/SlawaBE/go-metrics-collector/internal/storage"

func Run() {
    storage := storage.NewMemStorage()
    poller := NewPoller(storage, 2)
    reporter := NewReporter(storage, 10)

    go poller.Run()
    reporter.Run()
}