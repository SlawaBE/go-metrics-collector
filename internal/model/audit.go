package model

import "time"

// Audit Event is a data model for sending audit events.
type AuditEvent struct {
	Timestamp int64    `json:"ts"`         // event time
	Metrics   []string `json:"metrics"`    // set of obtained metrics
	IPAddress string   `json:"ip_address"` // ip address of the metric source
}

// NewAuditEvent generates an audit event from provided ip address and metric names.
func NewAuditEvent(ip string, metrics []string) AuditEvent {
	return AuditEvent{
		Timestamp: time.Now().Unix(),
		IPAddress: ip,
		Metrics:   metrics,
	}
}
