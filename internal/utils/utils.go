package utils

import (
	"strconv"
)

func ConvertCounter(value int64) string {
	return strconv.FormatInt(value, 10)
}

func ConvertGauge(value float64) string {
    return strconv.FormatFloat(value, 'f', -1, 64)
}