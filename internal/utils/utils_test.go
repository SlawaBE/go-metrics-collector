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

func TestGetIPFromAddress(t *testing.T) {
	tests := []struct {
		name    string
		address string
		want    string
		wantErr bool
	}{
		{name: "host and port", address: "localhost:8080", want: "localhost"},
		{name: "ipv4 and port", address: "192.168.1.10:9090", want: "192.168.1.10"},
		{name: "ipv6 and port", address: "[::1]:8080", want: "::1"},
		{name: "hostname without port", address: "localhost", wantErr: true},
		{name: "empty string", address: "", wantErr: true},
		{name: "port only", address: ":8080", want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetIPFromAddress(tt.address)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
