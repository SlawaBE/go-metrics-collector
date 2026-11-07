package config

import (
	"encoding/json"
	"flag"
	"testing"
	"time"
)

func TestDuration_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    time.Duration
		wantErr bool
	}{
		{name: "seconds", input: `"5s"`, want: 5 * time.Second},
		{name: "milliseconds", input: `"150ms"`, want: 150 * time.Millisecond},
		{name: "minutes", input: `"2m"`, want: 2 * time.Minute},
		{name: "hours", input: `"1h"`, want: time.Hour},
		{name: "composite", input: `"1h30m"`, want: time.Hour + 30*time.Minute},
		{name: "composite with seconds", input: `"1h30m15s"`, want: time.Hour + 30*time.Minute + 15*time.Second},
		{name: "microseconds", input: `"500us"`, want: 500 * time.Microsecond},
		{name: "micro symbol", input: `"500µs"`, want: 500 * time.Microsecond},
		{name: "nanoseconds", input: `"100ns"`, want: 100 * time.Nanosecond},
		{name: "zero", input: `"0s"`, want: 0},
		{name: "negative", input: `"-5s"`, want: -5 * time.Second},

		{name: "null", input: `null`, want: 0},
		{name: "empty string", input: `""`, want: 0},

		{name: "invalid string", input: `"abc"`, wantErr: true},
		{name: "no unit", input: `"5"`, wantErr: true},
		{name: "bad json", input: `{`, wantErr: true},
		{name: "number", input: `123`, wantErr: true},
		{name: "bool", input: `true`, wantErr: true},
		{name: "object", input: `{}`, wantErr: true},
		{name: "array", input: `[]`, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Duration(time.Hour)
			err := json.Unmarshal([]byte(tt.input), &d)

			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error, got nil (value=%v)", time.Duration(d))
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got := time.Duration(d); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDuration_NullResetsValue(t *testing.T) {
	d := Duration(42 * time.Second)
	if err := json.Unmarshal([]byte(`null`), &d); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if time.Duration(d) != 0 {
		t.Errorf("after null: got %v, want 0", time.Duration(d))
	}
}

func TestDuration_EmptyStringResetsValue(t *testing.T) {
	d := Duration(42 * time.Second)
	if err := json.Unmarshal([]byte(`""`), &d); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if time.Duration(d) != 0 {
		t.Errorf("after empty string: got %v, want 0", time.Duration(d))
	}
}

func TestDuration_MarshalJSON(t *testing.T) {
	tests := []struct {
		name  string
		want  string
		input time.Duration
	}{
		{name: "zero", input: 0, want: `"0s"`},
		{name: "nanoseconds", input: 100 * time.Nanosecond, want: `"100ns"`},
		{name: "microseconds", input: 500 * time.Microsecond, want: `"500µs"`},
		{name: "milliseconds", input: 150 * time.Millisecond, want: `"150ms"`},
		{name: "seconds", input: 5 * time.Second, want: `"5s"`},
		{name: "minutes", input: 2 * time.Minute, want: `"2m0s"`},
		{name: "hours", input: time.Hour, want: `"1h0m0s"`},
		{name: "composite", input: time.Hour + 30*time.Minute, want: `"1h30m0s"`},
		{name: "negative", input: -5 * time.Second, want: `"-5s"`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := json.Marshal(Duration(tt.input))
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if string(b) != tt.want {
				t.Errorf("got %s, want %s", b, tt.want)
			}
		})
	}
}

func TestDuration_RoundTripMarshalUnmarshal(t *testing.T) {
	values := []time.Duration{
		0,
		time.Nanosecond,
		time.Microsecond,
		time.Millisecond,
		time.Second,
		42 * time.Second,
		150 * time.Millisecond,
		time.Hour + 15*time.Minute + 30*time.Second,
		-time.Minute,
	}

	for _, v := range values {
		t.Run(v.String(), func(t *testing.T) {
			b, err := json.Marshal(Duration(v))
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}

			var back Duration
			if err := json.Unmarshal(b, &back); err != nil {
				t.Fatalf("unmarshal %s: %v", b, err)
			}
			if time.Duration(back) != v {
				t.Errorf("round-trip mismatch: got %v, want %v", time.Duration(back), v)
			}
		})
	}
}

func TestDuration_Duration(t *testing.T) {
	tests := []struct {
		name  string
		input Duration
		want  time.Duration
	}{
		{name: "zero", input: 0, want: 0},
		{name: "positive", input: Duration(5 * time.Second), want: 5 * time.Second},
		{name: "negative", input: Duration(-time.Minute), want: -time.Minute},
		{name: "composite", input: Duration(time.Hour + 30*time.Minute), want: time.Hour + 30*time.Minute},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.input.Duration(); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDuration_InStruct(t *testing.T) {
	type Config struct {
		Timeout  Duration `json:"timeout"`
		Interval Duration `json:"interval"`
	}

	t.Run("unmarshal all strings", func(t *testing.T) {
		data := []byte(`{"timeout":"30s","interval":"1h30m"}`)
		var cfg Config
		if err := json.Unmarshal(data, &cfg); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if cfg.Timeout.Duration() != 30*time.Second {
			t.Errorf("Timeout: got %v, want 30s", cfg.Timeout.Duration())
		}
		if cfg.Interval.Duration() != time.Hour+30*time.Minute {
			t.Errorf("Interval: got %v, want 1h30m", cfg.Interval.Duration())
		}
	})

	t.Run("unmarshal with null", func(t *testing.T) {
		data := []byte(`{"timeout":null,"interval":"1s"}`)
		var cfg Config
		if err := json.Unmarshal(data, &cfg); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if cfg.Timeout.Duration() != 0 {
			t.Errorf("Timeout: got %v, want 0", cfg.Timeout.Duration())
		}
		if cfg.Interval.Duration() != time.Second {
			t.Errorf("Interval: got %v, want 1s", cfg.Interval.Duration())
		}
	})

	t.Run("unmarshal with empty string", func(t *testing.T) {
		data := []byte(`{"timeout":"","interval":"2s"}`)
		var cfg Config
		if err := json.Unmarshal(data, &cfg); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if cfg.Timeout.Duration() != 0 {
			t.Errorf("Timeout: got %v, want 0", cfg.Timeout.Duration())
		}
	})

	t.Run("unmarshal invalid value", func(t *testing.T) {
		data := []byte(`{"timeout":"abc","interval":"1s"}`)
		var cfg Config
		if err := json.Unmarshal(data, &cfg); err == nil {
			t.Fatal("expected error, got nil")
		}
	})

	t.Run("unmarshal missing field stays zero", func(t *testing.T) {
		data := []byte(`{"interval":"1s"}`)
		var cfg Config
		if err := json.Unmarshal(data, &cfg); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if cfg.Timeout.Duration() != 0 {
			t.Errorf("Timeout: got %v, want 0", cfg.Timeout.Duration())
		}
	})

	t.Run("marshal", func(t *testing.T) {
		cfg := Config{
			Timeout:  Duration(30 * time.Second),
			Interval: Duration(time.Hour + 30*time.Minute),
		}
		b, err := json.Marshal(cfg)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		want := `{"timeout":"30s","interval":"1h30m0s"}`
		if string(b) != want {
			t.Errorf("got %s, want %s", b, want)
		}
	})

	t.Run("round-trip via struct", func(t *testing.T) {
		orig := Config{
			Timeout:  Duration(5 * time.Second),
			Interval: Duration(2 * time.Minute),
		}
		b, err := json.Marshal(orig)
		if err != nil {
			t.Fatalf("marshal: %v", err)
		}

		var back Config
		if err := json.Unmarshal(b, &back); err != nil {
			t.Fatalf("unmarshal %s: %v", b, err)
		}
		if back != orig {
			t.Errorf("round-trip mismatch: got %+v, want %+v", back, orig)
		}
	})
}

func TestDuration_UnmarshalJSON_NilReceiver(t *testing.T) {
	var d *Duration
	if err := json.Unmarshal([]byte(`"1s"`), &d); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if d == nil {
		t.Fatal("expected non-nil pointer")
	}
	if time.Duration(*d) != time.Second {
		t.Errorf("got %v, want 1s", time.Duration(*d))
	}
}

func TestDuration_Set(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    time.Duration
		wantErr bool
	}{
		{name: "simple seconds", input: "5s", want: 5 * time.Second},
		{name: "minutes and seconds", input: "1m30s", want: 90 * time.Second},
		{name: "hours", input: "2h", want: 2 * time.Hour},
		{name: "zero", input: "0s", want: 0},
		{name: "negative", input: "-10s", want: -10 * time.Second},
		{name: "milliseconds", input: "150ms", want: 150 * time.Millisecond},
		{name: "empty string", input: "", want: 0},
		{name: "null-like handled via UnmarshalJSON", input: "", want: 0},
		{name: "invalid number", input: "abc", wantErr: true},
		{name: "missing unit", input: "10", wantErr: true},
		{name: "spaces", input: " 5s", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var d Duration
			err := d.Set(tt.input)

			if tt.wantErr {
				if err == nil {
					t.Fatalf("Set(%q): expected error, got nil (value=%v)", tt.input, d.Duration())
				}
				return
			}
			if err != nil {
				t.Fatalf("Set(%q): unexpected error: %v", tt.input, err)
			}
			if got := d.Duration(); got != tt.want {
				t.Errorf("Set(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestDuration_String(t *testing.T) {
	tests := []struct {
		name string
		want string
		d    Duration
	}{
		{name: "zero", d: 0, want: "0s"},
		{name: "seconds", d: Duration(5 * time.Second), want: "5s"},
		{name: "complex", d: Duration(90 * time.Second), want: "1m30s"},
		{name: "negative", d: Duration(-10 * time.Second), want: "-10s"},
		{name: "milliseconds", d: Duration(150 * time.Millisecond), want: "150ms"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.String(); got != tt.want {
				t.Errorf("String() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestDuration_RoundTripSetString(t *testing.T) {
	values := []time.Duration{
		0,
		5 * time.Second,
		90 * time.Second,
		2 * time.Hour,
		-10 * time.Second,
		150 * time.Millisecond,
	}
	for _, v := range values {
		t.Run(v.String(), func(t *testing.T) {
			orig := Duration(v)
			var parsed Duration
			if err := parsed.Set(orig.String()); err != nil {
				t.Fatalf("Set(%q): %v", orig.String(), err)
			}
			if parsed != orig {
				t.Errorf("round-trip failed: %v -> %q -> %v", orig.Duration(), orig.String(), parsed.Duration())
			}
		})
	}
}

func TestDuration_SetMatchesUnmarshalJSON(t *testing.T) {
	inputs := []string{"5s", "1m30s", "0s", "-10s", "150ms", "", "abc", "10"}

	for _, in := range inputs {
		t.Run(in, func(t *testing.T) {
			var viaSet Duration
			setErr := viaSet.Set(in)

			var viaJSON Duration
			b, _ := json.Marshal(in)
			jsonErr := viaJSON.UnmarshalJSON(b)

			// Сравниваем "ошибочность" и значение при успехе.
			if (setErr == nil) != (jsonErr == nil) {
				t.Fatalf("error mismatch: Set err=%v, UnmarshalJSON err=%v", setErr, jsonErr)
			}
			if setErr == nil && viaSet != viaJSON {
				t.Errorf("value mismatch: Set=%v, UnmarshalJSON=%v", viaSet.Duration(), viaJSON.Duration())
			}
		})
	}
}

func TestDuration_FlagIntegration(t *testing.T) {
	var d Duration
	fs := flag.NewFlagSet("test", flag.ContinueOnError)
	fs.Var(&d, "timeout", "timeout duration")

	if err := fs.Parse([]string{"-timeout=1m30s"}); err != nil {
		t.Fatalf("Parse: %v", err)
	}
	if got := d.Duration(); got != 90*time.Second {
		t.Errorf("timeout = %v, want 1m30s", got)
	}
}

func TestDuration_FlagAbsent(t *testing.T) {
	d := Duration(5 * time.Second) // дефолт
	fs := flag.NewFlagSet("test", flag.ContinueOnError)
	fs.Var(&d, "timeout", "timeout duration")

	if err := fs.Parse(nil); err != nil {
		t.Fatalf("Parse: %v", err)
	}
	if got := d.Duration(); got != 5*time.Second {
		t.Errorf("default = %v, want 5s", got)
	}
}

func TestDuration_FlagInvalid(t *testing.T) {
	var d Duration
	fs := flag.NewFlagSet("test", flag.ContinueOnError)
	fs.Var(&d, "timeout", "timeout duration")

	err := fs.Parse([]string{"-timeout=abc"})
	if err == nil {
		t.Fatalf("expected parse error for -timeout=abc, got nil")
	}
}
