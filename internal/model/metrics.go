package model

import "fmt"

// Metric type values
const (
	Counter = "counter"
	Gauge   = "gauge"
)

// Metric is a data model for storing and working with metrics.
type Metric struct {
	ID    string   `json:"id"`              // metric id
	MType string   `json:"type"`            // metric type
	Delta *int64   `json:"delta,omitempty"` // metric value for type 'counter'
	Value *float64 `json:"value,omitempty"` // metric value for type 'gauge'
}

// NewCounterMetric build metric with type 'counter'.
func NewCounterMetric(name string, value int64) Metric {
	return Metric{
		ID:    name,
		MType: Counter,
		Delta: &value,
	}
}

// NewGaugeMetric build metric with type 'gauge'.
func NewGaugeMetric(name string, value float64) Metric {
	return Metric{
		ID:    name,
		MType: Gauge,
		Value: &value,
	}
}

// String generates a string representation of the metric.
func (m *Metric) String() string {
	if m.MType == Counter {
		return fmt.Sprintf("Counter: %s %d", m.ID, *m.Delta)
	}
	if m.MType == Gauge {
		return fmt.Sprintf("Gauge: %s %f", m.ID, *m.Value)
	}
	return ""
}

// MetricRequest is a data model for a request to retrieve the value of a specific metric.
type MetricRequest struct {
	ID    string `json:"id"`
	MType string `json:"type"`
}
