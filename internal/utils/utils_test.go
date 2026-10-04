package utils

import (
	"testing"

	"github.com/SlawaBE/go-metrics-collector/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestConvertCounter(t *testing.T) {
	assert.Equal(t, "0", ConvertCounter(0))
	assert.Equal(t, "42", ConvertCounter(42))
	assert.Equal(t, "-7", ConvertCounter(-7))
	assert.Equal(t, "9223372036854775807", ConvertCounter(int64(9223372036854775807)))
}

func TestConvertGauge(t *testing.T) {
	assert.Equal(t, "0", ConvertGauge(0))
	assert.Equal(t, "1.5", ConvertGauge(1.5))
	assert.Equal(t, "-2.25", ConvertGauge(-2.25))
	assert.Equal(t, "3.141592653589793", ConvertGauge(3.141592653589793))
}

func TestPtr(t *testing.T) {
	m := model.NewCounterMetric("counter-1", 10)

	p := Ptr(m)

	assert.NotNil(t, p)
	assert.Equal(t, m, *p)
}
