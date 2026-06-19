package agent

import (
	"math/rand"
	"runtime"

	"github.com/SlawaBE/go-metrics-collector/internal/model"
)

type Metrics struct {
	// метрики из runtime
	Alloc         float64
	BuckHashSys   float64
	Frees         float64
	GCCPUFraction float64
	GCSys         float64
	HeapAlloc     float64
	HeapIdle      float64
	HeapInuse     float64
	HeapObjects   float64
	HeapReleased  float64
	HeapSys       float64
	LastGC        float64
	Lookups       float64
	MCacheInuse   float64
	MCacheSys     float64
	MSpanInuse    float64
	MSpanSys      float64
	Mallocs       float64
	NextGC        float64
	NumForcedGC   float64
	NumGC         float64
	OtherSys      float64
	PauseTotalNs  float64
	StackInuse    float64
	StackSys      float64
	Sys           float64
	TotalAlloc    float64
	// дополнительные метрики
	RandomValue float64
	PollCount   int64
}

func GetRuntimeMetrics() *Metrics {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	m := Metrics{}

	m.Alloc = float64(memStats.Alloc)
	m.BuckHashSys = float64(memStats.BuckHashSys)
	m.Frees = float64(memStats.Frees)
	m.GCCPUFraction = memStats.GCCPUFraction
	m.GCSys = float64(memStats.GCSys)
	m.HeapAlloc = float64(memStats.HeapAlloc)
	m.HeapIdle = float64(memStats.HeapIdle)
	m.HeapInuse = float64(memStats.HeapInuse)
	m.HeapObjects = float64(memStats.HeapObjects)
	m.HeapReleased = float64(memStats.HeapReleased)
	m.HeapSys = float64(memStats.HeapSys)
	m.LastGC = float64(memStats.LastGC)
	m.Lookups = float64(memStats.Lookups)
	m.MCacheInuse = float64(memStats.MCacheInuse)
	m.MCacheSys = float64(memStats.MCacheSys)
	m.MSpanInuse = float64(memStats.MSpanInuse)
	m.MSpanSys = float64(memStats.MSpanSys)
	m.Mallocs = float64(memStats.Mallocs)
	m.NextGC = float64(memStats.NextGC)
	m.NumForcedGC = float64(memStats.NumForcedGC)
	m.NumGC = float64(memStats.NumGC)
	m.OtherSys = float64(memStats.OtherSys)
	m.PauseTotalNs = float64(memStats.PauseTotalNs)
	m.StackInuse = float64(memStats.StackInuse)
	m.StackSys = float64(memStats.StackSys)
	m.Sys = float64(memStats.Sys)
	m.TotalAlloc = float64(memStats.TotalAlloc)
	m.RandomValue = rand.Float64()
	m.PollCount = 1

	return &m
}

func (m *Metrics) convertToMap() map[string]model.Metric {
	return map[string]model.Metric{
		"Alloc":         model.NewGaugeMetric("Alloc", m.Alloc),
		"BuckHashSys":   model.NewGaugeMetric("BuckHashSys", m.BuckHashSys),
		"Frees":         model.NewGaugeMetric("Frees", m.Frees),
		"GCCPUFraction": model.NewGaugeMetric("GCCPUFraction", m.GCCPUFraction),
		"GCSys":         model.NewGaugeMetric("GCSys", m.GCSys),
		"HeapAlloc":     model.NewGaugeMetric("HeapAlloc", m.HeapAlloc),
		"HeapIdle":      model.NewGaugeMetric("HeapIdle", m.HeapIdle),
		"HeapInuse":     model.NewGaugeMetric("HeapInuse", m.HeapInuse),
		"HeapObjects":   model.NewGaugeMetric("HeapObjects", m.HeapObjects),
		"HeapReleased":  model.NewGaugeMetric("HeapReleased", m.HeapReleased),
		"HeapSys":       model.NewGaugeMetric("HeapSys", m.HeapSys),
		"LastGC":        model.NewGaugeMetric("LastGC", m.LastGC),
		"Lookups":       model.NewGaugeMetric("Lookups", m.Lookups),
		"MCacheInuse":   model.NewGaugeMetric("MCacheInuse", m.MCacheInuse),
		"MCacheSys":     model.NewGaugeMetric("MCacheSys", m.MCacheSys),
		"MSpanInuse":    model.NewGaugeMetric("MSpanInuse", m.MSpanInuse),
		"MSpanSys":      model.NewGaugeMetric("MSpanSys", m.MSpanSys),
		"Mallocs":       model.NewGaugeMetric("Mallocs", m.Mallocs),
		"NextGC":        model.NewGaugeMetric("NextGC", m.NextGC),
		"NumForcedGC":   model.NewGaugeMetric("NumForcedGC", m.NumForcedGC),
		"NumGC":         model.NewGaugeMetric("NumGC", m.NumGC),
		"OtherSys":      model.NewGaugeMetric("OtherSys", m.OtherSys),
		"PauseTotalNs":  model.NewGaugeMetric("PauseTotalNs", m.PauseTotalNs),
		"StackInuse":    model.NewGaugeMetric("StackInuse", m.StackInuse),
		"StackSys":      model.NewGaugeMetric("StackSys", m.StackSys),
		"Sys":           model.NewGaugeMetric("Sys", m.Sys),
		"TotalAlloc":    model.NewGaugeMetric("TotalAlloc", m.TotalAlloc),
		"RandomValue":   model.NewGaugeMetric("RandomValue", m.RandomValue),
		"PollCount":     model.NewCounterMetric("PollCount", m.PollCount),
	}
}

func (m *Metrics) convertToList() []model.Metric {
    return []model.Metric{
        model.NewGaugeMetric("Alloc", m.Alloc),
        model.NewGaugeMetric("BuckHashSys", m.BuckHashSys),
        model.NewGaugeMetric("Frees", m.Frees),
        model.NewGaugeMetric("GCCPUFraction", m.GCCPUFraction),
        model.NewGaugeMetric("GCSys", m.GCSys),
        model.NewGaugeMetric("HeapAlloc", m.HeapAlloc),
        model.NewGaugeMetric("HeapIdle", m.HeapIdle),
        model.NewGaugeMetric("HeapInuse", m.HeapInuse),
        model.NewGaugeMetric("HeapObjects", m.HeapObjects),
        model.NewGaugeMetric("HeapReleased", m.HeapReleased),
        model.NewGaugeMetric("HeapSys", m.HeapSys),
        model.NewGaugeMetric("LastGC", m.LastGC),
        model.NewGaugeMetric("Lookups", m.Lookups),
        model.NewGaugeMetric("MCacheInuse", m.MCacheInuse),
        model.NewGaugeMetric("MCacheSys", m.MCacheSys),
        model.NewGaugeMetric("MSpanInuse", m.MSpanInuse),
        model.NewGaugeMetric("MSpanSys", m.MSpanSys),
        model.NewGaugeMetric("Mallocs", m.Mallocs),
        model.NewGaugeMetric("NextGC", m.NextGC),
        model.NewGaugeMetric("NumForcedGC", m.NumForcedGC),
        model.NewGaugeMetric("NumGC", m.NumGC),
        model.NewGaugeMetric("OtherSys", m.OtherSys),
        model.NewGaugeMetric("PauseTotalNs", m.PauseTotalNs),
        model.NewGaugeMetric("StackInuse", m.StackInuse),
        model.NewGaugeMetric("StackSys", m.StackSys),
        model.NewGaugeMetric("Sys", m.Sys),
        model.NewGaugeMetric("TotalAlloc", m.TotalAlloc),
        model.NewGaugeMetric("RandomValue", m.RandomValue),
        model.NewCounterMetric("PollCount", m.PollCount),
    }
}
