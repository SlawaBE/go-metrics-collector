package agent

import (
	"fmt"
	"math/rand"
	"runtime"

	"github.com/SlawaBE/go-metrics-collector/internal/model"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/mem"
)

type Metrics interface {
	convertToList() []model.Metric
}

type RuntimeMetrics struct {
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

func GetRuntimeMetrics() *RuntimeMetrics {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	m := RuntimeMetrics{}

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

func (m *RuntimeMetrics) convertToMap() map[string]model.Metric {
	res := make(map[string]model.Metric)
	for _, v := range m.convertToList() {
		res[v.ID] = v
	}
	return res
}

func (m *RuntimeMetrics) convertToList() []model.Metric {
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

type GopsutilsMetrics struct {
	TotalMemory    *float64
	FreeMemory     *float64
	CPUutilization []float64
}

func GetGopsutilsMetrics() *GopsutilsMetrics {
	m := GopsutilsMetrics{}

	memoryStat, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println("failed read virtual memory stat")
	} else {
		m.TotalMemory = new(float64(memoryStat.Total))
		m.FreeMemory = new(float64(memoryStat.Free))
	}

	m.CPUutilization, err = cpu.Percent(0, true)
	if err != nil {
		fmt.Println("failed read CPU utilization")
	}

	return &m
}

func (m *GopsutilsMetrics) convertToList() []model.Metric {
	list := []model.Metric{}
	if m.FreeMemory != nil {
		list = append(list, model.NewGaugeMetric("FreeMemory", *m.FreeMemory))
	}
	if m.TotalMemory != nil {
		list = append(list, model.NewGaugeMetric("TotalMemory", *m.TotalMemory))
	}

	if m.CPUutilization != nil {
		for numCpu, percentage := range m.CPUutilization {
			list = append(list, model.NewGaugeMetric(fmt.Sprintf("CPUutilization%d", numCpu+1), percentage))
		}
	}	

	return list
}
