package model

import "time"

// AuditEvent describes an audit event raised when metrics are received.
type AuditEvent struct {
	// Timestamp - Unix-время возникновения события.
	Timestamp int64 `json:"ts"`
	// Metrics is the list of metric names affected by the event.
	Metrics []string `json:"metrics"`
	// IPAddress is the IP address of the client that sent the metrics.
	IPAddress string `json:"ip_address"`
}

// NewAuditEvent creates an audit event with the current time.
func NewAuditEvent(ip string, metrics []string) AuditEvent {
	return AuditEvent{
		Timestamp: time.Now().Unix(),
		IPAddress: ip,
		Metrics:   metrics,
	}
}
