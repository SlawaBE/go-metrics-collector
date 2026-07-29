package handler

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SlawaBE/go-metrics-collector/internal/model"
	"github.com/SlawaBE/go-metrics-collector/internal/service"
	"github.com/SlawaBE/go-metrics-collector/internal/storage"
	"github.com/go-chi/chi/v5"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestListMetricHandler_ServeHTTP(t *testing.T) {
	type input struct {
		method  string
		metrics []model.Metric
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
			name: "Get all metrics invalid method",
			data: input{
				method: http.MethodPost,
			},
			want: output{
				statusCode: 405,
			},
		},
		{
			name: "Success counter metric get",
			data: input{
				method: http.MethodGet,
				metrics: []model.Metric{
					model.NewCounterMetric("A", 1),
					model.NewGaugeMetric("B", 1.1),
				},
			},
			want: output{
				statusCode: 200,
				value:      "A: 1\nB: 1.1\n",
			},
		},
	}
	stor := storage.NewMemStorage()
	saver := service.NewJsonFileMetricSaver(-1, "", stor)
	h := NewListMetricHandler(service.NewMetricsService(stor, saver))

	r := chi.NewRouter()
	r.Handle("GET /", h)
	srv := httptest.NewServer(r)
	defer srv.Close()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.data.metrics != nil {
				for _, m := range tt.data.metrics {
					stor.UpdateMetric(context.Background(), m)
				}
			}
			url := "/"

			req := resty.New().R()
			req.Method = tt.data.method
			req.URL = srv.URL + url

			resp, err := req.Send()
			assert.NoError(t, err, "error making HTTP request")

			assert.Equal(t, tt.want.statusCode, resp.StatusCode(), "Response code didn't match expected")
			if tt.want.value != "" {
				assert.Equal(t, tt.want.value, string(resp.Body()))
			}
		})
	}
}
