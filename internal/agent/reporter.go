package agent

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/SlawaBE/go-metrics-collector/internal/model"
	"github.com/SlawaBE/go-metrics-collector/internal/storage"
)

type Reporter struct {
	storage        Storage
	reportInterval int
	baseUrl        string
}

type Storage interface {
    storage.Storage
    GetValuesAndClear() []model.Metric
}

func NewReporter(storage Storage, reportInterval int) *Reporter {
	return &Reporter{
		storage:        storage,
		reportInterval: reportInterval,
		baseUrl:        "http://localhost:8080/update/",
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
	url := r.baseUrl + metric.MType + "/" + metric.ID + "/"
	if metric.MType == model.Counter {
		url = url + strconv.FormatInt(*metric.Delta, 10)
	}
	if metric.MType == model.Gauge {
		url = url + strconv.FormatFloat(*metric.Value, 'e', -1, 64)
	}
	res, err := http.Post(url, "text/plain", nil)
	if err != nil {
		return fmt.Errorf("error sending metric: %v", err)
	}
    if  res.StatusCode != 200 {
        return fmt.Errorf("error in server response: %v", err)
    }
	res.Body.Close()
	return nil
}
