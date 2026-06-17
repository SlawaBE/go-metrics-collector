package storage

type MemStorage struct {
	gauges   map[string]float64
	counters map[string]int64
}

func NewMemStorage() *MemStorage {
	return &MemStorage{
		gauges:   make(map[string]float64),
		counters: make(map[string]int64),
	}
}

func (storage *MemStorage) SaveCounter(key string, value int64) error {
	storage.counters[key] += value
	return nil
}

func (storage *MemStorage) SaveGauge(key string, value float64) error {
	storage.gauges[key] = value
	return nil
}
