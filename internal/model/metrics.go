package model

import "fmt"

// Supported metric types.
const (
	// Counter is an accumulating counter (int64). On repeated updates its
	// value is added to the previous one.
	Counter = "counter"
	// Gauge is an arbitrary floating-point number (float64). On repeated
	// updates its value is overwritten.
	Gauge = "gauge"
)

// Metric describes a single metric in a flat (non-nested) model.
//
// Delta and Value are declared as pointers to distinguish the value 0 from an
// unset value and, consequently, to avoid encoding it into the JSON structure
// (the fields are marked with omitempty).
type Metric struct {
	// ID is the unique metric name.
	ID string `json:"id"`
	// MType is the metric type (Counter or Gauge).
	MType string `json:"type"`
	// Delta is the accumulating counter value (for Counter).
	Delta *int64 `json:"delta,omitempty"`
	// Value is the floating-point metric value (for Gauge).
	Value *float64 `json:"value,omitempty"`
	// Hash is an optional HMAC signature of the metric.
	Hash string `json:"hash,omitempty"`
}

// NewCounterMetric creates a Counter metric with the given name and value.
func NewCounterMetric(name string, value int64) Metric {
	return Metric{
		ID:    name,
		MType: Counter,
		Delta: &value,
	}
}

// NewGaugeMetric creates a Gauge metric with the given name and value.
func NewGaugeMetric(name string, value float64) Metric {
	return Metric{
		ID:    name,
		MType: Gauge,
		Value: &value,
	}
}

// String returns the string representation of the metric depending on its type.
func (m *Metric) String() string {
	if m.MType == Counter {
		return fmt.Sprintf("Counter: %s %d", m.ID, *m.Delta)
	}
	if m.MType == Gauge {
		return fmt.Sprintf("Gauge: %s %f", m.ID, *m.Value)
	}
	return ""
}

// MetricRequest describes a request to fetch a metric by its identifier.
type MetricRequest struct {
	// ID is the unique metric name.
	ID string `json:"id"`
	// MType is the metric type (Counter or Gauge).
	MType string `json:"type"`
}
