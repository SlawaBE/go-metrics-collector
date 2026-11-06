// Package config holds custom struct with UnmarshalJSON for time.Duration
package config

import (
	"encoding/json"
	"time"
)

type Duration time.Duration

func (d *Duration) UnmarshalJSON(b []byte) error {
	s := string(b)
	if s == "null" || s == `""` {
		*d = 0
		return nil
	}
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	duration, err := time.ParseDuration(s)
	if err != nil {
		return err
	}
	*d = Duration(duration)
	return nil
}

func (d Duration) Duration() time.Duration {
	return time.Duration(d)
}

func (d Duration) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Duration(d).String())
}

// String возвращает строковое представление для справки flag.
func (d Duration) String() string {
	return time.Duration(d).String()
}

// Set вызывается пакетом flag при разборе аргумента командной строки.
// Переиспользуем UnmarshalJSON: оборачиваем входную строку в кавычки,
// чтобы получить валидный JSON-литерал.
func (d *Duration) Set(s string) error {
	b, err := json.Marshal(s) // превращаем "5s" в `"5s"`
	if err != nil {
		return err
	}
	return d.UnmarshalJSON(b)
}