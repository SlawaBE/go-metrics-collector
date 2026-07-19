package utils

import (
	"strconv"

	"github.com/SlawaBE/go-metrics-collector/internal/model"
)

func ConvertCounter(value int64) string {
	return strconv.FormatInt(value, 10)
}

func ConvertGauge(value float64) string {
	return strconv.FormatFloat(value, 'f', -1, 64)
}

func Ptr(m model.Metric) *model.Metric {
	return &m
}
