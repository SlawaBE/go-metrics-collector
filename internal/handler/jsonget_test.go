package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SlawaBE/go-metrics-collector/internal/model"
	"github.com/SlawaBE/go-metrics-collector/internal/service"
	"github.com/SlawaBE/go-metrics-collector/internal/storage"
	"github.com/SlawaBE/go-metrics-collector/internal/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetMetricHandler_ServeHTTP(t *testing.T) {
	type input struct {
		method string
		mType  string
		id     string
		metric *model.Metric
		json   string
	}
	type output struct {
		statusCode int
		value      string
	}
	tests := []struct {
		name string
		data input
		want output
	}{
		{
			name: "Unknown metric type",
			data: input{
				method: http.MethodPost,
				mType:  "unknown",
				id:     "name",
				json:   "{\"id\":\"name\",\"type\":\"unknown\"}",
			},
			want: output{
				statusCode: 404,
				value:      "",
			},
		},
		{
			name: "Invalid method",
			data: input{
				method: http.MethodPut,
				mType:  "gauge",
				id:     "name",
				json:   "{\"id\":\"name\",\"type\":\"gauge\"}",
			},
			want: output{
				statusCode: 405,
				value:      "",
			},
		},
		{
			name: "Success counter metric get",
			data: input{
				method: http.MethodPost,
				mType:  "counter",
				id:     "name",
				metric: utils.Ptr(model.NewCounterMetric("name", 1)),
				json:   "{\"id\":\"name\",\"type\":\"counter\"}",
			},
			want: output{
				statusCode: 200,
				value:      "1",
			},
		},
		{
			name: "Success gauge metric get",
			data: input{
				method: http.MethodPost,
				mType:  "gauge",
				id:     "name",
				metric: utils.Ptr(model.NewGaugeMetric("name", 1.1)),
				json:   "{\"id\":\"name\",\"type\":\"gauge\"}",
			},
			want: output{
				statusCode: 200,
				value:      "1.1",
			},
		},
		{
			name: "Get counter metric without name (is \"\")",
			data: input{
				method: http.MethodPost,
				mType:  "counter",
				id:     "",
				json:   "{\"id\":\"\",\"type\":\"counter\"}",
			},
			want: output{
				statusCode: 404,
				value:      "",
			},
		},
		{
			name: "Get gauge metric without name (is \"\")",
			data: input{
				method: http.MethodPost,
				mType:  "gauge",
				id:     "",
				json:   "{\"id\":\"\",\"type\":\"gauge\"}",
			},
			want: output{
				statusCode: 404,
				value:      "",
			},
		},
		{
			name: "Get counter metric without name (is null)",
			data: input{
				method: http.MethodPost,
				mType:  "counter",
				id:     "",
				json:   "{\"type\":\"counter\"}",
			},
			want: output{
				statusCode: 404,
				value:      "",
			},
		},
		{
			name: "Get gauge metric without name (is null)",
			data: input{
				method: http.MethodPost,
				mType:  "gauge",
				id:     "",
				json:   "{\"type\":\"gauge\"}",
			},
			want: output{
				statusCode: 404,
				value:      "",
			},
		},
	}
	stor := storage.NewMemStorage()
	saver := service.NewJsonFileMetricSaver(-1, "", stor)
	h := NewJsonGetMetricHandler(service.NewMetricsService(stor, saver))

	r := chi.NewRouter()
	r.Handle("POST /value", h)
	srv := httptest.NewServer(r)
	defer srv.Close()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.data.metric != nil {
				stor.UpdateMetric(*tt.data.metric)
			}
			url := "/value"

			req := resty.New().R()
			req.Method = tt.data.method
			req.URL = srv.URL + url
			req.Header.Add("Content-Type", "application/json")
			req.SetBody(tt.data.json)

			resp, err := req.Send()
			assert.NoError(t, err, "error making HTTP request")

			assert.Equal(t, tt.want.statusCode, resp.StatusCode(), "Response code didn't match expected")
			var metric model.Metric
			if tt.want.value == "" {
				return
			}

			json.Unmarshal(resp.Body(), &metric)
			assert.Equal(t, tt.data.id, metric.ID)
			assert.Equal(t, tt.data.mType, metric.MType)
			if metric.MType == "counter" {
				assert.Equal(t, tt.want.value, utils.ConvertCounter(*metric.Delta))
			}
			if metric.MType == "gauge" {
				assert.Equal(t, tt.want.value, utils.ConvertGauge(*metric.Value))
			}
		})
	}
}
