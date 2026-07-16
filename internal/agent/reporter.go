package agent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/SlawaBE/go-metrics-collector/internal/model"
	"github.com/SlawaBE/go-metrics-collector/internal/storage"
)

type Reporter struct {
	storage        Storage
	reportInterval int
	baseURL        string
}

type Storage interface {
	storage.Storage
	GetValuesAndClear() []model.Metric
}

func NewReporter(storage Storage, reportAddress string, reportInterval int) *Reporter {
	return &Reporter{
		storage:        storage,
		reportInterval: reportInterval,
		baseURL:        "http://" + reportAddress + "/update",
	}
}

func (r *Reporter) Run() {
	for {
		time.Sleep(time.Duration(r.reportInterval) * time.Second)
		r.Report()
	}
}

func (r *Reporter) Report() {
	metrics := r.storage.GetValuesAndClear()
	for _, m := range metrics {
		if err := r.sendMetric(m); err != nil {
			fmt.Println("Error sending metric:", m.ID)
		}
	}
}

func (r *Reporter) sendMetric(metric model.Metric) error {
	jsonData, err := json.Marshal(metric)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}
	res, err := http.Post(r.baseURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error sending metric: %v", err)
	}
	if res.StatusCode != 200 {
		return fmt.Errorf("error in server response: %v", err)
	}
	res.Body.Close()
	return nil
}
