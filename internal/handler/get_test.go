package handler

import (
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

func TestJsonGetMetricHandler_ServeHTTP(t *testing.T) {
	type input struct {
		method   string
		pathType string
		pathName string
		metric   *model.Metric
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
				method:   http.MethodGet,
				pathType: "unknown",
				pathName: "name",
			},
			want: output{
				statusCode: 404,
				value:      "",
			},
		},
		{
			name: "Invalid method",
			data: input{
				method:   http.MethodPut,
				pathType: "gauge",
				pathName: "name",
			},
			want: output{
				statusCode: 405,
				value:      "",
			},
		},
		{
			name: "Success counter metric get",
			data: input{
				method:   http.MethodGet,
				pathType: "counter",
				pathName: "name",
				metric:   utils.Ptr(model.NewCounterMetric("name", 1)),
			},
			want: output{
				statusCode: 200,
				value:      "1",
			},
		},
		{
			name: "Success gauge metric get",
			data: input{
				method:   http.MethodGet,
				pathType: "gauge",
				pathName: "name",
				metric:   utils.Ptr(model.NewGaugeMetric("name", 1.1)),
			},
			want: output{
				statusCode: 200,
				value:      "1.1",
			},
		},
		{
			name: "Get counter metric without name",
			data: input{
				method:   http.MethodGet,
				pathType: "counter",
				pathName: "",
			},
			want: output{
				statusCode: 404,
				value:      "",
			},
		},
		{
			name: "Get gauge metric without name",
			data: input{
				method:   http.MethodGet,
				pathType: "gauge",
				pathName: "",
			},
			want: output{
				statusCode: 404,
				value:      "",
			},
		},
	}
	stor := storage.NewMemStorage()
	saver := service.NewJsonFileMetricSaver(-1, "", stor)
	h := NewGetMetricHandler(service.NewMetricsService(stor, saver))

	r := chi.NewRouter()
	r.Handle("GET /value/{type}/{name}", h)
	srv := httptest.NewServer(r)
	defer srv.Close()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.data.metric != nil {
				stor.UpdateMetric(*tt.data.metric)
			}
			url := "/value/" + tt.data.pathType + "/" + tt.data.pathName

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
