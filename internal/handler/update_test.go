package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SlawaBE/go-metrics-collector/internal/service"
	"github.com/SlawaBE/go-metrics-collector/internal/storage"
	"github.com/go-chi/chi/v5"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestUpdateMetricHandler_ServeHTTP(t *testing.T) {
	type input struct {
		method    string
		pathType  string
		pathName  string
		pathValue string
	}
	type output struct {
		statusCode int
	}
	tests := []struct {
		name string
		data input
		want output
	}{
		{
			name: "Unknown metric type",
			data: input{
				method:    http.MethodPost,
				pathType:  "unknown",
				pathName:  "name",
				pathValue: "1",
			},
			want: output{
				statusCode: 400,
			},
		},
		{
			name: "Invalid counter metric value",
			data: input{
				method:    http.MethodPost,
				pathType:  "counter",
				pathName:  "name",
				pathValue: "a",
			},
			want: output{
				statusCode: 400,
			},
		},
		{
			name: "Invalid gauge metric value",
			data: input{
				method:    http.MethodPost,
				pathType:  "gauge",
				pathName:  "name",
				pathValue: "b",
			},
			want: output{
				statusCode: 400,
			},
		},
		{
			name: "Invalid method",
			data: input{
				method:    http.MethodPut,
				pathType:  "gauge",
				pathName:  "name",
				pathValue: "1.0",
			},
			want: output{
				statusCode: 405,
			},
		},
		{
			name: "Success counter metric update",
			data: input{
				method:    http.MethodPost,
				pathType:  "counter",
				pathName:  "name",
				pathValue: "1",
			},
			want: output{
				statusCode: 200,
			},
		},
		{
			name: "Success gauge metric update",
			data: input{
				method:    http.MethodPost,
				pathType:  "gauge",
				pathName:  "name",
				pathValue: "1.0",
			},
			want: output{
				statusCode: 200,
			},
		},
		{
			name: "Update counter metric without name",
			data: input{
				method:    http.MethodPost,
				pathType:  "counter",
				pathName:  "",
				pathValue: "1",
			},
			want: output{
				statusCode: 404,
			},
		},
		{
			name: "Update gauge metric without name",
			data: input{
				method:    http.MethodPost,
				pathType:  "gauge",
				pathName:  "",
				pathValue: "1.0",
			},
			want: output{
				statusCode: 404,
			},
		},
	}
	service := service.NewMetricsService(storage.NewMemStorage())
	handler := NewUpdateMetricHandler(service)

	r := chi.NewRouter()
	r.Handle("POST /update/{type}/{name}/{value}", handler)
	srv := httptest.NewServer(r)
	defer srv.Close()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/update/" + tt.data.pathType + "/" + tt.data.pathName + "/" + tt.data.pathValue

			req := resty.New().R()
			req.Method = tt.data.method
			req.URL = srv.URL + url

			resp, err := req.Send()
			assert.NoError(t, err, "error making HTTP request")

			assert.Equal(t, tt.want.statusCode, resp.StatusCode(), "Response code didn't match expected")
		})
	}
}
