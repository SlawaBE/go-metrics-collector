package agent

import "github.com/SlawaBE/go-metrics-collector/internal/storage"

func Run(address string, pollInterval int, reportInterval int) {
	storage := storage.NewMemStorage()
	poller := NewPoller(storage, pollInterval)
	reporter := NewReporter(storage, address, reportInterval)

	go poller.Run()
	reporter.Run()
}
