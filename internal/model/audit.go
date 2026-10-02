package model

import "time"

type AuditEvent struct {
	Timestamp int64    `json:"ts"`
	Metrics   []string `json:"metrics"`
	IPAddress string   `json:"ip_address"`
}

func NewAuditEvent(ip string, metrics []string) AuditEvent {
	return AuditEvent{
		Timestamp: time.Now().Unix(),
		IPAddress: ip,
		Metrics:   metrics,
	}
}
