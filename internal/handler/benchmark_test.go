package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SlawaBE/go-metrics-collector/internal/model"
	"github.com/SlawaBE/go-metrics-collector/internal/service"
	"github.com/SlawaBE/go-metrics-collector/internal/storage"
)

func newBenchService(b *testing.B) *service.MetricsService {
	svc := service.NewMetricsService(storage.NewMemStorage(), nil)
	return svc
}

func newBenchAuditPublisher(b *testing.B) service.AuditPublisher {
	publisher, err := service.NewAuditService(1 << 20)
	if err != nil {
		b.Fatal(err)
	}
	return publisher
}

func BenchmarkJSONUpdateMetricHandler(b *testing.B) {
	svc := newBenchService(b)
	publisher := newBenchAuditPublisher(b)
	h := NewJSONUpdateMetricHandler(svc, publisher)

	body := []byte(`{"id":"bench","type":"gauge","value":1.1}`)

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest(http.MethodPost, "/update", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		h.ServeHTTP(rec, req)
	}
}

func BenchmarkJSONBatchUpdateMetricsHandler(b *testing.B) {
	svc := newBenchService(b)
	publisher := newBenchAuditPublisher(b)
	h := NewJSONBatchUpdateMetricsHandler(svc, publisher)

	body := []byte(`[{"id":"g0","type":"gauge","value":1.1},{"id":"g1","type":"gauge","value":2.2},{"id":"g2","type":"gauge","value":3.3},{"id":"g3","type":"gauge","value":4.4},{"id":"g4","type":"gauge","value":5.5}]`)

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest(http.MethodPost, "/updates", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		h.ServeHTTP(rec, req)
	}
}

func BenchmarkJSONGetMetricHandler(b *testing.B) {
	svc := newBenchService(b)
	if err := svc.UpdateMetric(b.Context(), model.NewGaugeMetric("bench", 1.1)); err != nil {
		b.Fatal(err)
	}
	h := NewJSONGetMetricHandler(svc)

	body := []byte(`{"id":"bench","type":"gauge"}`)

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest(http.MethodPost, "/value", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		h.ServeHTTP(rec, req)
	}
}
