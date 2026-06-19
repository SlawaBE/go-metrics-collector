package model

import "fmt"

const (
	Counter = "counter"
	Gauge   = "gauge"
)

// NOTE: Не усложняем пример, вводя иерархическую вложенность структур.
// Органичиваясь плоской моделью.
// Delta и Value объявлены через указатели,
// что бы отличать значение "0", от не заданного значения
// и соответственно не кодировать в структуру.
type Metric struct {
	ID    string   `json:"id"`
	MType string   `json:"type"`
	Delta *int64   `json:"delta,omitempty"`
	Value *float64 `json:"value,omitempty"`
	Hash  string   `json:"hash,omitempty"`
}

func NewCounterMetric(name string, value int64) Metric {
	return Metric{
		ID:    name,
		MType: Counter,
		Delta: &value,
	}
}

func NewGaugeMetric(name string, value float64) Metric {
	return Metric{
		ID:    name,
		MType: Gauge,
		Value: &value,
	}
}

func (m *Metric) String() string {
	if m.MType == Counter {
		return fmt.Sprintf("Counter: %s %d", m.ID, *m.Delta)
	}
    if m.MType == Gauge {
        return fmt.Sprintf("Gauge: %s %f", m.ID, *m.Value)
    }
    return ""
}
