// Package utils contains helper utilities for converting metric values and
// working with pointers.
package utils

import (
	"net"
	"strconv"

	"github.com/SlawaBE/go-metrics-collector/internal/model"
)

// ConvertCounter converts a counter value to a string.
func ConvertCounter(value int64) string {
	return strconv.FormatInt(value, 10)
}

// ConvertGauge converts a gauge value to a string in its shortest form.
func ConvertGauge(value float64) string {
	return strconv.FormatFloat(value, 'f', -1, 64)
}

// Ptr returns a pointer to the given metric.
func Ptr(m model.Metric) *model.Metric {
	return &m
}

// GetIPFromAddress extracts ip from address string
func GetIPFromAddress(address string) (string, error) {
	host, _, err := net.SplitHostPort(address)
	return host, err
}
