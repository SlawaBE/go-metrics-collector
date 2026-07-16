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

func TestJsonUpdateMetricHandler_ServeHTTP(t *testing.T) {
	type input struct {
		method string
		json   string
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
				method: http.MethodPost,
				json:   "{\"id\":\"name\",\"type\":\"unknown\",\"value\":1.1}",
			},
			want: output{
				statusCode: 400,
			},
		},
		{
			name: "Invalid counter metric value",
			data: input{
				method: http.MethodPost,
				json:   "{\"id\":\"name\",\"type\":\"counter\",\"value\":\"a\"}",
			},
			want: output{
				statusCode: 400,
			},
		},
		{
			name: "Invalid gauge metric value",
			data: input{
				method: http.MethodPost,
				json:   "{\"id\":\"name\",\"type\":\"gauge\",\"value\":\"b\"}",
			},
			want: output{
				statusCode: 400,
			},
		},
		{
			name: "Invalid method",
			data: input{
				method: http.MethodPut,
				json:   "{\"id\":\"name\",\"type\":\"gauge\",\"value\":1.1}",
			},
			want: output{
				statusCode: 405,
			},
		},
		{
			name: "Success counter metric update",
			data: input{
				method: http.MethodPost,
				json:   "{\"id\":\"name\",\"type\":\"counter\",\"delta\":1}",
			},
			want: output{
				statusCode: 200,
			},
		},
		{
			name: "Success gauge metric update",
			data: input{
				method: http.MethodPost,
				json:   "{\"id\":\"name\",\"type\":\"gauge\",\"value\":1.1}",
			},
			want: output{
				statusCode: 200,
			},
		},
		{
			name: "Update counter metric without name (is \"\")",
			data: input{
				method: http.MethodPost,
				json:   "{\"id\":\"\",\"type\":\"counter\",\"delta\":1}",
			},
			want: output{
				statusCode: 404,
			},
		},
		{
			name: "Update counter metric without name (is null)",
			data: input{
				method: http.MethodPost,
				json:   "{\"type\":\"counter\",\"delta\":1}",
			},
			want: output{
				statusCode: 404,
			},
		},
		{
			name: "Update gauge metric without name",
			data: input{
				method: http.MethodPost,
				json:   "{\"id\":\"\",\"type\":\"unknown\",\"value\":1.1}",
			},
			want: output{
				statusCode: 404,
			},
		},
		{
			name: "Update gauge metric without name (is null)",
			data: input{
				method: http.MethodPost,
				json:   "{\"type\":\"gauge\",\"value\":1.1}",
			},
			want: output{
				statusCode: 404,
			},
		},
	}
	service := service.NewMetricsService(storage.NewMemStorage())
	handler := NewJsonUpdateMetricHandler(service)

	r := chi.NewRouter()
	r.Handle("POST /update", handler)
	srv := httptest.NewServer(r)
	defer srv.Close()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/update"

			req := resty.New().R()
			req.Method = tt.data.method
			req.URL = srv.URL + url
			req.Header.Add("Content-Type", "application/json")
			req.Body = tt.data.json

			resp, err := req.Send()
			assert.NoError(t, err, "error making HTTP request")

			assert.Equal(t, tt.want.statusCode, resp.StatusCode(), "Response code didn't match expected")
		})
	}
}
