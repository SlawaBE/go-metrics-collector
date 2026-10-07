package server

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/SlawaBE/go-metrics-collector/internal/model"
	"github.com/SlawaBE/go-metrics-collector/internal/service"
	"github.com/SlawaBE/go-metrics-collector/internal/storage"
	"github.com/go-chi/chi/v5"
)

// newExampleRouter builds a working router on an in-memory storage without a
// file saver and without a database.
func newExampleRouter() chi.Router {
	stor := storage.NewMemStorage()
	metricsService := service.NewMetricsService(stor, nil)
	audit, err := service.NewAuditService(100)
	if err != nil {
		panic(fmt.Sprintf("cannot create audit service: %v", err))
	}
	return InitRouter(metricsService, nil, audit)
}

// request performs an HTTP request against the test server and returns the
// status code and the response body.
func request(srv *httptest.Server, method, url string, body string) (int, string) {
	var reader io.Reader
	if body != "" {
		reader = strings.NewReader(body)
	}

	req, err := http.NewRequest(method, srv.URL+url, reader)
	if err != nil {
		panic(fmt.Sprintf("cannot create request: %v", err))
	}
	if body != "" {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := srv.Client().Do(req)
	if err != nil {
		panic(fmt.Sprintf("cannot do request: %v", err))
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(fmt.Sprintf("cannot read response body: %v", err))
	}
	return resp.StatusCode, string(data)
}

// ExampleInitRouter_urlUpdateAndGet demonstrates updating a metric via the URL
// path (POST /update/{type}/{name}/{value}) and fetching its value
// (GET /value/{type}/{name}).
func ExampleInitRouter_urlUpdateAndGet() {
	srv := httptest.NewServer(newExampleRouter())
	defer srv.Close()

	status, _ := request(srv, http.MethodPost, "/update/gauge/CPUUtilization1/73.5", "")
	fmt.Printf("update status: %d\n", status)

	status, body := request(srv, http.MethodGet, "/value/gauge/CPUUtilization1", "")
	fmt.Printf("get status: %d, value: %s\n", status, body)

	// Output:
	// update status: 200
	// get status: 200, value: 73.5
}

// ExampleInitRouter_jsonUpdate demonstrates updating a metric via a JSON
// request (POST /update).
func ExampleInitRouter_jsonUpdate() {
	srv := httptest.NewServer(newExampleRouter())
	defer srv.Close()

	body := `{"id":"HeapAlloc","type":"gauge","value":1024.5}`
	status, respBody := request(srv, http.MethodPost, "/update", body)
	fmt.Printf("status: %d, body: %s\n", status, respBody)

	// Output:
	// status: 200, body: {}
}

// ExampleInitRouter_jsonGet demonstrates fetching a metric via a JSON request
// (POST /value).
func ExampleInitRouter_jsonGet() {
	stor := storage.NewMemStorage()
	metricsService := service.NewMetricsService(stor, nil)
	audit, err := service.NewAuditService(100)
	if err != nil {
		panic(fmt.Sprintf("cannot create audit service: %v", err))
	}

	if err := stor.UpdateMetric(context.TODO(), model.NewGaugeMetric("HeapAlloc", 1024.5)); err != nil {
		panic(err)
	}

	srv := httptest.NewServer(InitRouter(metricsService, nil, audit))
	defer srv.Close()

	body := `{"id":"HeapAlloc","type":"gauge"}`
	status, respBody := request(srv, http.MethodPost, "/value", body)
	fmt.Printf("status: %d\nvalue found: %t\n", status, strings.Contains(respBody, "1024.5"))

	// Output:
	// status: 200
	// value found: true
}

// ExampleInitRouter_batchUpdate demonstrates batch updating of several metrics
// via a JSON request (POST /updates).
func ExampleInitRouter_batchUpdate() {
	srv := httptest.NewServer(newExampleRouter())
	defer srv.Close()

	body := `[
		{"id":"PollCount","type":"counter","delta":1},
		{"id":"RandomValue","type":"gauge","value":0.5}
	]`
	status, _ := request(srv, http.MethodPost, "/updates", body)
	fmt.Printf("status: %d\n", status)

	// Output:
	// status: 200
}

// ExampleInitRouter_list demonstrates rendering the list of all metrics as
// HTML (GET /).
func ExampleInitRouter_list() {
	stor := storage.NewMemStorage()
	metricsService := service.NewMetricsService(stor, nil)
	audit, err := service.NewAuditService(100)
	if err != nil {
		panic(fmt.Sprintf("cannot create audit service: %v", err))
	}

	if err := stor.UpdateMetric(context.TODO(), model.NewCounterMetric("PollCount", 10)); err != nil {
		panic(err)
	}

	srv := httptest.NewServer(InitRouter(metricsService, nil, audit))
	defer srv.Close()

	status, body := request(srv, http.MethodGet, "/", "")
	fmt.Printf("status: %d, contains PollCount: %t\n", status, strings.Contains(body, "PollCount"))

	// Output:
	// status: 200, contains PollCount: true
}

// ExampleInitRouter_ping demonstrates the database connectivity check
// (GET /ping). Without a configured DB the endpoint returns 200.
func ExampleInitRouter_ping() {
	srv := httptest.NewServer(newExampleRouter())
	defer srv.Close()

	status, _ := request(srv, http.MethodGet, "/ping", "")
	fmt.Printf("status: %d\n", status)

	// Output:
	// status: 200
}
