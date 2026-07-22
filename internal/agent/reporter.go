package agent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/SlawaBE/go-metrics-collector/internal/gzip"
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
		return fmt.Errorf("failed to marshal JSON: %v", err)
	}

	data, err := gzip.Compress(jsonData)
	if err != nil {
		return fmt.Errorf("error compress metric: %v", err)
	}

	req, err := http.NewRequest("POST", r.baseURL, bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Content-Encoding", "gzip")
	req.Header.Set("Accept-Encoding", "gzip")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending metric: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(res.Body)
		return fmt.Errorf("error in server response: %d %s", res.StatusCode, string(body))
	}

	_, err = io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	return nil
}
