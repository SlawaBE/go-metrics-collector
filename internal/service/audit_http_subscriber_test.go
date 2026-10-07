package service

import (
	"encoding/json"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/SlawaBE/go-metrics-collector/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHTTPAuditSubscriber_Notify(t *testing.T) {
	var received string
	var receivedHeader string

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		receivedHeader = r.Header.Get("Content-Type")
		body, _ := io.ReadAll(r.Body)
		received = string(body)
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	sub := NewHTTPAuditSubscriber(srv.URL)
	event := model.NewAuditEvent("192.168.1.1", []string{"metric-1"})

	err := sub.Notify(event)
	require.NoError(t, err)

	assert.Equal(t, "application/json", receivedHeader)
	var decoded model.AuditEvent
	require.NoError(t, json.Unmarshal([]byte(received), &decoded))
	assert.Equal(t, event.IPAddress, decoded.IPAddress)
	assert.Equal(t, event.Metrics, decoded.Metrics)
}

func TestHTTPAuditSubscriber_ErrorStatus(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer srv.Close()

	sub := NewHTTPAuditSubscriber(srv.URL)
	event := model.NewAuditEvent("192.168.1.1", []string{"metric-1"})

	err := sub.Notify(event)
	require.Error(t, err)
}

func TestHTTPAuditSubscriber_ConnectionReuse(t *testing.T) {
	var mu sync.Mutex
	newConns := 0

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	srv := httptest.NewUnstartedServer(handler)
	oldConnState := srv.Config.ConnState
	srv.Config.ConnState = func(c net.Conn, cs http.ConnState) {
		if cs == http.StateNew {
			mu.Lock()
			newConns++
			mu.Unlock()
		}
		if oldConnState != nil {
			oldConnState(c, cs)
		}
	}
	srv.Start()
	defer srv.Close()

	sub := NewHTTPAuditSubscriber(srv.URL)

	for i := 0; i < 5; i++ {
		err := sub.Notify(model.NewAuditEvent("192.168.1.1", []string{"m"}))
		require.NoError(t, err)
	}

	mu.Lock()
	defer mu.Unlock()
	assert.Equal(t, 1, newConns, "keep-alive должен переиспользовать одно соединение для нескольких событий")
}
