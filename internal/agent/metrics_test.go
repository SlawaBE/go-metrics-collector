package agent

import (
	"testing"

	"github.com/SlawaBE/go-metrics-collector/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestMetrics_GetRuntimeMetrics(t *testing.T) {
	tests := []struct {
		mType string
		name  string
	}{
		{
			mType: model.Gauge,
			name:  "Alloc",
		},
		{
			mType: model.Gauge,
			name:  "BuckHashSys",
		},
		{
			mType: model.Gauge,
			name:  "Frees",
		},
		{
			mType: model.Gauge,
			name:  "GCCPUFraction",
		},
		{
			mType: model.Gauge,
			name:  "GCSys",
		},
		{
			mType: model.Gauge,
			name:  "HeapAlloc",
		},
		{
			mType: model.Gauge,
			name:  "HeapIdle",
		},
		{
			mType: model.Gauge,
			name:  "HeapInuse",
		},
		{
			mType: model.Gauge,
			name:  "HeapObjects",
		},
		{
			mType: model.Gauge,
			name:  "HeapReleased",
		},
		{
			mType: model.Gauge,
			name:  "HeapSys",
		},
		{
			mType: model.Gauge,
			name:  "LastGC",
		},
		{
			mType: model.Gauge,
			name:  "Lookups",
		},
		{
			mType: model.Gauge,
			name:  "MCacheInuse",
		},
		{
			mType: model.Gauge,
			name:  "MCacheSys",
		},
		{
			mType: model.Gauge,
			name:  "MSpanInuse",
		},
		{
			mType: model.Gauge,
			name:  "MSpanSys",
		},
		{
			mType: model.Gauge,
			name:  "Mallocs",
		},
		{
			mType: model.Gauge,
			name:  "NextGC",
		},
		{
			mType: model.Gauge,
			name:  "NumForcedGC",
		},
		{
			mType: model.Gauge,
			name:  "NumGC",
		},
		{
			mType: model.Gauge,
			name:  "OtherSys",
		},
		{
			mType: model.Gauge,
			name:  "PauseTotalNs",
		},
		{
			mType: model.Gauge,
			name:  "StackInuse",
		},
		{
			mType: model.Gauge,
			name:  "StackSys",
		},
		{
			mType: model.Gauge,
			name:  "Sys",
		},
		{
			mType: model.Gauge,
			name:  "TotalAlloc",
		},
		{
			mType: model.Gauge,
			name:  "RandomValue",
		},
		{
			mType: model.Counter,
			name:  "PollCount",
		},
	}

	metrics := GetRuntimeMetrics()
	mMap := metrics.convertToMap()

	assert.Len(t, mMap, 29)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, ok := mMap[tt.name]
			assert.True(t, true, ok)
			assert.Equal(t, tt.name, m.ID)
			assert.Equal(t, tt.mType, m.MType)
			if tt.mType == model.Gauge {
				assert.Nil(t, m.Delta)
				assert.NotNil(t, m.Value)
			} else if tt.mType == model.Counter {
				assert.NotNil(t, m.Delta)
				assert.Nil(t, m.Value)
			}
		})
	}
}
