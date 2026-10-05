package service

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/SlawaBE/go-metrics-collector/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFileAuditSubscriber_WritesEvent(t *testing.T) {
	path := filepath.Join(t.TempDir(), "audit.log")
	sub := NewFileAuditSubscriber(path)

	event := model.NewAuditEvent("192.168.1.1", []string{"counter-1"})
	require.NoError(t, sub.Notify(context.Background(), event))

	data, err := os.ReadFile(path)
	require.NoError(t, err)

	var decoded model.AuditEvent
	require.NoError(t, json.Unmarshal(data, &decoded))
	assert.Equal(t, event.IPAddress, decoded.IPAddress)
	assert.Equal(t, event.Metrics, decoded.Metrics)
}

func TestFileAuditSubscriber_Appends(t *testing.T) {
	path := filepath.Join(t.TempDir(), "audit.log")
	sub := NewFileAuditSubscriber(path)

	event := model.NewAuditEvent("192.168.1.1", []string{"a"})
	require.NoError(t, sub.Notify(context.Background(), event))
	require.NoError(t, sub.Notify(context.Background(), event))

	data, err := os.ReadFile(path)
	require.NoError(t, err)

	assert.Equal(t, 2, bytesCountNewlines(data), "каждый вызов должен дописывать отдельную JSON-строку")
}

func TestFileAuditSubscriber_InvalidPath(t *testing.T) {
	sub := NewFileAuditSubscriber(filepath.Join(t.TempDir(), "no-such-dir", "x.log"))

	err := sub.Notify(context.Background(), model.NewAuditEvent("ip", []string{"m"}))
	assert.Error(t, err)
}

func bytesCountNewlines(data []byte) int {
	n := 0
	for _, b := range data {
		if b == '\n' {
			n++
		}
	}
	return n
}
