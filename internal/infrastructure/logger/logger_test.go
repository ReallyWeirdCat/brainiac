package logger

import (
	"bytes"
	"context"
	"encoding/json"
	"testing"

	"github.com/ReallyWeirdCat/brainiac/pkg/domain/app/ports"
	"github.com/rs/zerolog"
)

var _ ports.Logger = (*ZerologLogger)(nil)

func newTestLogger() (ports.Logger, *bytes.Buffer) {
	buf := &bytes.Buffer{}
	zl := zerolog.New(buf).Level(zerolog.DebugLevel) // capture everything
	return NewZerologLogger(zl), buf
}

// parseLogEntry unmarshals the first JSON line from the buffer.
func parseLogEntry(t *testing.T, buf *bytes.Buffer) map[string]any {
	t.Helper()
	line := bytes.TrimSpace(buf.Bytes())
	if len(line) == 0 {
		t.Fatal("buffer is empty, expected a log entry")
	}
	var entry map[string]any
	if err := json.Unmarshal(line, &entry); err != nil {
		t.Fatalf("failed to parse log entry: %v", err)
	}
	buf.Reset()
	return entry
}

func TestArgsToFields(t *testing.T) {
	tests := []struct {
		name     string
		args     []any
		expected map[string]any
	}{
		{
			name:     "no arguments",
			args:     nil,
			expected: nil,
		},
		{
			name:     "empty slice",
			args:     []any{},
			expected: nil,
		},
		{
			name: "even pairs",
			args: []any{"key1", "value1", "key2", 42},
			expected: map[string]any{
				"key1": "value1",
				"key2": 42,
			},
		},
		{
			name: "odd number appends '?'",
			args: []any{"odd_key", "odd_value", "dangling"},
			expected: map[string]any{
				"odd_key":  "odd_value",
				"dangling": "?",
			},
		},
		{
			name: "non-string key converted with fmt.Sprint",
			args: []any{123, "val", true, false},
			expected: map[string]any{
				"123":  "val",
				"true": false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := argsToFields(tt.args...)
			if len(result) != len(tt.expected) {
				t.Errorf("got %d fields, want %d", len(result), len(tt.expected))
			}
			for k, v := range tt.expected {
				rv, ok := result[k]
				if !ok {
					t.Errorf("missing key %q", k)
					continue
				}
				if rv != v {
					t.Errorf("key %q: got %v, want %v", k, rv, v)
				}
			}
		})
	}
}

func TestZerologLogger_Debug(t *testing.T) {
	log, buf := newTestLogger()
	log.Debug("test message", "key", "value")
	entry := parseLogEntry(t, buf)

	if level := entry["level"]; level != "debug" {
		t.Errorf("level = %q, want debug", level)
	}
	if msg := entry["message"]; msg != "test message" {
		t.Errorf("message = %q, want test message", msg)
	}
	if v, ok := entry["key"]; !ok || v != "value" {
		t.Errorf("missing or wrong field 'key': %v", v)
	}
}

func TestZerologLogger_Info(t *testing.T) {
	log, buf := newTestLogger()
	log.Info("info msg", "num", 7)
	entry := parseLogEntry(t, buf)

	if level := entry["level"]; level != "info" {
		t.Errorf("level = %q, want info", level)
	}
	if msg := entry["message"]; msg != "info msg" {
		t.Errorf("message = %q, want info msg", msg)
	}
	if v, ok := entry["num"]; !ok {
		t.Error("missing field 'num'")
	} else if n, ok := v.(float64); !ok || int(n) != 7 {
		t.Errorf("field 'num' = %v, want 7", v)
	}
}

func TestZerologLogger_Warn(t *testing.T) {
	log, buf := newTestLogger()
	log.Warn("warning", "active", true)
	entry := parseLogEntry(t, buf)

	if level := entry["level"]; level != "warn" {
		t.Errorf("level = %q, want warn", level)
	}
	if msg := entry["message"]; msg != "warning" {
		t.Errorf("message = %q, want warning", msg)
	}
	if v, ok := entry["active"]; !ok || v != true {
		t.Errorf("field 'active' = %v, want true", v)
	}
}

func TestZerologLogger_Error(t *testing.T) {
	log, buf := newTestLogger()
	log.Error("error occurred", "code", 500)
	entry := parseLogEntry(t, buf)

	if level := entry["level"]; level != "error" {
		t.Errorf("level = %q, want error", level)
	}
	if msg := entry["message"]; msg != "error occurred" {
		t.Errorf("message = %q, want error occurred", msg)
	}
	if v, ok := entry["code"]; !ok {
		t.Error("missing field 'code'")
	} else if n, ok := v.(float64); !ok || int(n) != 500 {
		t.Errorf("field 'code' = %v, want 500", v)
	}
}

func TestZerologLogger_With(t *testing.T) {
	baseLog, buf := newTestLogger()
	childLog := baseLog.With("component", "test", "pid", 42)

	// Log via child – should include added fields
	childLog.Info("inside component")
	entry := parseLogEntry(t, buf)

	if msg := entry["message"]; msg != "inside component" {
		t.Errorf("message = %q, want inside component", msg)
	}
	if v, ok := entry["component"]; !ok || v != "test" {
		t.Errorf("field 'component' = %v, want test", v)
	}
	if v, ok := entry["pid"]; !ok {
		t.Error("missing field 'pid'")
	} else if n, ok := v.(float64); !ok || int(n) != 42 {
		t.Errorf("field 'pid' = %v, want 42", v)
	}

	// Log via base – should not contain child's fields
	baseLog.Info("base message")
	entry = parseLogEntry(t, buf)
	if msg := entry["message"]; msg != "base message" {
		t.Errorf("message = %q, want base message", msg)
	}
	if _, ok := entry["component"]; ok {
		t.Error("base logger should not have 'component' field")
	}
}

func TestZerologLogger_WithContext(t *testing.T) {
	baseLog, buf := newTestLogger()

	// Context without any logger
	ctx := context.Background()
	ctxLog := baseLog.WithContext(ctx)
	if ctxLog != baseLog {
		t.Error("expected the same logger instance when no logger is in context")
	}
	ctxLog.Info("no context logger")
	entry := parseLogEntry(t, buf)
	if msg := entry["message"]; msg != "no context logger" {
		t.Errorf("message = %q, want no context logger", msg)
	}

	zl := zerolog.New(buf).Level(zerolog.DebugLevel).With().Str("from", "context").Logger()
	ctx = zl.WithContext(context.Background())

	ctxLog = baseLog.WithContext(ctx)
	ctxLog.Info("context logger used")
	entry = parseLogEntry(t, buf)
	if msg := entry["message"]; msg != "context logger used" {
		t.Errorf("message = %q, want context logger used", msg)
	}
	if v, ok := entry["from"]; !ok || v != "context" {
		t.Errorf("expected field 'from' = 'context', got %v", v)
	}
}
