package agent

import (
	"fmt"
	"net/http"
	"time"

	"github.com/SlawaBE/go-metrics-collector/internal/model"
	"github.com/SlawaBE/go-metrics-collector/internal/storage"
	"github.com/SlawaBE/go-metrics-collector/internal/utils"
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
		baseURL:        "http://" + reportAddress + "/update/",
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
	url := r.baseURL + metric.MType + "/" + metric.ID + "/"
	if metric.MType == model.Counter {
		url = url + utils.ConvertCounter(*metric.Delta)
	}
	if metric.MType == model.Gauge {
		url = url + utils.ConvertGauge(*metric.Value)
	}
	res, err := http.Post(url, "text/plain", nil)
	if err != nil {
		return fmt.Errorf("error sending metric: %v", err)
	}
	if res.StatusCode != 200 {
		return fmt.Errorf("error in server response: %v", err)
	}
	res.Body.Close()
	return nil
}
