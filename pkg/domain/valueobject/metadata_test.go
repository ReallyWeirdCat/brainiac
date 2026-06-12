// Copyright (c) 2026 Nikolai Papin
//
// This file is part of Brainiac gamification and education platform
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package valueobject

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestNewMetadata(t *testing.T) {
	got := NewMetadata()
	if got.data == nil {
		t.Error("NewMetadata() returned nil internal map")
	}
	if len(got.data) != 0 {
		t.Errorf("NewMetadata() map length = %d, want 0", len(got.data))
	}
	if !got.IsEmpty() {
		t.Error("NewMetadata() should be empty")
	}
}

func TestMetadataFromMap(t *testing.T) {
	tests := []struct {
		name string
		raw  map[string]any
		want map[string]any
	}{
		{
			name: "nil input",
			raw:  nil,
			want: map[string]any{},
		},
		{
			name: "empty map",
			raw:  map[string]any{},
			want: map[string]any{},
		},
		{
			name: "non-empty map",
			raw:  map[string]any{"key": "value", "num": 42},
			want: map[string]any{"key": "value", "num": 42},
		},
		{
			name: "nested structure",
			raw:  map[string]any{"nested": map[string]any{"inner": true}},
			want: map[string]any{"nested": map[string]any{"inner": true}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MetadataFromMap(tt.raw)
			if !reflect.DeepEqual(got.data, tt.want) {
				t.Errorf("MetadataFromMap().data = %v, want %v", got.data, tt.want)
			}
			// Verify it's a copy: modify original map and ensure metadata unchanged.
			if len(tt.raw) > 0 {
				originalKey := ""
				for k := range tt.raw {
					originalKey = k
					break
				}
				delete(tt.raw, originalKey)
				if _, ok := got.data[originalKey]; !ok {
					t.Error("MetadataFromMap() did not copy the map; modification affected metadata")
				}
			}
		})
	}
}

func TestMetadata_Set(t *testing.T) {
	tests := []struct {
		name      string
		initial   map[string]any
		key       string
		value     any
		wantData  map[string]any
		wantEmpty bool
	}{
		{
			name:      "set on nil map",
			initial:   nil,
			key:       "a",
			value:     1,
			wantData:  map[string]any{"a": 1},
			wantEmpty: false,
		},
		{
			name:      "set new key on existing map",
			initial:   map[string]any{"b": 2},
			key:       "c",
			value:     3,
			wantData:  map[string]any{"b": 2, "c": 3},
			wantEmpty: false,
		},
		{
			name:      "overwrite existing key",
			initial:   map[string]any{"x": "old"},
			key:       "x",
			value:     "new",
			wantData:  map[string]any{"x": "new"},
			wantEmpty: false,
		},
		{
			name:      "set nil value",
			initial:   map[string]any{},
			key:       "nilVal",
			value:     nil,
			wantData:  map[string]any{"nilVal": nil},
			wantEmpty: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Metadata{data: tt.initial}
			m.Set(tt.key, tt.value)
			if !reflect.DeepEqual(m.data, tt.wantData) {
				t.Errorf("After Set, data = %v, want %v", m.data, tt.wantData)
			}
			if gotEmpty := m.IsEmpty(); gotEmpty != tt.wantEmpty {
				t.Errorf("IsEmpty() = %v, want %v", gotEmpty, tt.wantEmpty)
			}
		})
	}
}

func TestMetadata_Get(t *testing.T) {
	tests := []struct {
		name      string
		data      map[string]any
		key       string
		wantValue any
		wantOk    bool
	}{
		{
			name:      "nil map",
			data:      nil,
			key:       "any",
			wantValue: nil,
			wantOk:    false,
		},
		{
			name:      "key exists",
			data:      map[string]any{"foo": "bar"},
			key:       "foo",
			wantValue: "bar",
			wantOk:    true,
		},
		{
			name:      "key does not exist",
			data:      map[string]any{"foo": "bar"},
			key:       "baz",
			wantValue: nil,
			wantOk:    false,
		},
		{
			name:      "key with nil value",
			data:      map[string]any{"nilKey": nil},
			key:       "nilKey",
			wantValue: nil,
			wantOk:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Metadata{data: tt.data}
			got, ok := m.Get(tt.key)
			if !reflect.DeepEqual(got, tt.wantValue) {
				t.Errorf("Get() value = %v, want %v", got, tt.wantValue)
			}
			if ok != tt.wantOk {
				t.Errorf("Get() ok = %v, want %v", ok, tt.wantOk)
			}
		})
	}
}

func TestMetadata_Delete(t *testing.T) {
	tests := []struct {
		name     string
		initial  map[string]any
		key      string
		wantData map[string]any
	}{
		{
			name:     "delete from nil map",
			initial:  nil,
			key:      "key",
			wantData: nil, // remains nil
		},
		{
			name:     "delete existing key",
			initial:  map[string]any{"a": 1, "b": 2},
			key:      "a",
			wantData: map[string]any{"b": 2},
		},
		{
			name:     "delete non-existing key",
			initial:  map[string]any{"a": 1},
			key:      "missing",
			wantData: map[string]any{"a": 1},
		},
		{
			name:     "delete last key",
			initial:  map[string]any{"only": "value"},
			key:      "only",
			wantData: map[string]any{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Metadata{data: tt.initial}
			m.Delete(tt.key)
			if !reflect.DeepEqual(m.data, tt.wantData) {
				t.Errorf("After Delete, data = %v, want %v", m.data, tt.wantData)
			}
		})
	}
}

func TestMetadata_Keys(t *testing.T) {
	tests := []struct {
		name string
		data map[string]any
		want []string
	}{
		{
			name: "nil map",
			data: nil,
			want: []string{},
		},
		{
			name: "empty map",
			data: map[string]any{},
			want: []string{},
		},
		{
			name: "single key",
			data: map[string]any{"foo": 1},
			want: []string{"foo"},
		},
		{
			name: "multiple keys",
			data: map[string]any{"a": 1, "b": 2, "c": 3},
			want: []string{"a", "b", "c"}, // order will be checked via set equivalence
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Metadata{data: tt.data}
			got := m.Keys()
			if len(got) != len(tt.want) {
				t.Errorf("Keys() length = %d, want %d", len(got), len(tt.want))
			}
			if tt.want != nil {
				// Check that all wanted keys are present.
				wantSet := make(map[string]bool, len(tt.want))
				for _, k := range tt.want {
					wantSet[k] = true
				}
				for _, k := range got {
					if !wantSet[k] {
						t.Errorf("Keys() returned unexpected key %q", k)
					}
				}
			}
		})
	}
}

func TestMetadata_AsMap(t *testing.T) {
	tests := []struct {
		name string
		data map[string]any
		want map[string]any
	}{
		{
			name: "nil map",
			data: nil,
			want: nil,
		},
		{
			name: "empty map",
			data: map[string]any{},
			want: map[string]any{},
		},
		{
			name: "non-empty map",
			data: map[string]any{"k1": "v1", "k2": 2},
			want: map[string]any{"k1": "v1", "k2": 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Metadata{data: tt.data}
			got := m.AsMap()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AsMap() = %v, want %v", got, tt.want)
			}
			// Verify it's a copy by modifying the returned map.
			if len(got) > 0 {
				got["newKey"] = "shouldNotAppearInOriginal"
				if _, exists := m.data["newKey"]; exists {
					t.Error("AsMap() did not return a copy; modification affected original metadata")
				}
			}
		})
	}
}

func TestMetadata_IsEmpty(t *testing.T) {
	tests := []struct {
		name string
		data map[string]any
		want bool
	}{
		{"nil map", nil, true},
		{"empty map", map[string]any{}, true},
		{"non-empty map", map[string]any{"x": 1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Metadata{data: tt.data}
			if got := m.IsEmpty(); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMetadata_IsZero(t *testing.T) {
	// IsZero returns true for empty or nil map (same as IsEmpty).
	tests := []struct {
		name string
		data map[string]any
		want bool
	}{
		{"nil map", nil, true},
		{"empty map", map[string]any{}, true},
		{"non-empty map", map[string]any{"a": 1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Metadata{data: tt.data}
			if got := m.IsZero(); got != tt.want {
				t.Errorf("IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMetadata_Equals(t *testing.T) {
	tests := []struct {
		name  string
		m     Metadata
		other any
		want  bool
	}{
		{
			name:  "equal empty metadata",
			m:     Metadata{data: map[string]any{}},
			other: Metadata{data: map[string]any{}},
			want:  true,
		},
		{
			name:  "equal with values",
			m:     Metadata{data: map[string]any{"a": 1, "b": "text"}},
			other: Metadata{data: map[string]any{"a": 1, "b": "text"}},
			want:  true,
		},
		{
			name:  "different values",
			m:     Metadata{data: map[string]any{"a": 1}},
			other: Metadata{data: map[string]any{"a": 2}},
			want:  false,
		},
		{
			name:  "different keys",
			m:     Metadata{data: map[string]any{"a": 1}},
			other: Metadata{data: map[string]any{"b": 1}},
			want:  false,
		},
		{
			name:  "nil vs empty map",
			m:     Metadata{data: nil},
			other: Metadata{data: map[string]any{}},
			want:  true, // both are empty, Equals should treat them equal
		},
		{
			name:  "nil vs non-empty",
			m:     Metadata{data: nil},
			other: Metadata{data: map[string]any{"x": 1}},
			want:  false,
		},
		{
			name:  "different type",
			m:     Metadata{data: map[string]any{}},
			other: "not a metadata",
			want:  false,
		},
		{
			name:  "deep equality nested",
			m:     Metadata{data: map[string]any{"nested": map[string]any{"inner": true}}},
			other: Metadata{data: map[string]any{"nested": map[string]any{"inner": true}}},
			want:  true,
		},
		{
			name:  "deep equality different nested",
			m:     Metadata{data: map[string]any{"nested": map[string]any{"inner": true}}},
			other: Metadata{data: map[string]any{"nested": map[string]any{"inner": false}}},
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Equals(tt.other); got != tt.want {
				t.Errorf("Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMetadata_String(t *testing.T) {
	tests := []struct {
		name string
		data map[string]any
		want string
	}{
		{"nil map", nil, "{}"},
		{"empty map", map[string]any{}, "{}"},
		{"single key", map[string]any{"a": 1}, `{"a":1}`},
		{"multiple keys", map[string]any{"a": 1, "b": "two"}, `{"a":1,"b":"two"}`}, // order may vary
		{"nested", map[string]any{"n": map[string]any{"x": 2}}, `{"n":{"x":2}}`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Metadata{data: tt.data}
			got := m.String()
			// For maps with multiple keys, JSON order is not guaranteed; we can unmarshal and compare.
			var gotMap, wantMap map[string]any
			if err := json.Unmarshal([]byte(got), &gotMap); err != nil {
				t.Fatalf("String() returned invalid JSON: %v", err)
			}
			if err := json.Unmarshal([]byte(tt.want), &wantMap); err != nil {
				t.Fatalf("want string is invalid JSON: %v", err)
			}
			if !reflect.DeepEqual(gotMap, wantMap) {
				t.Errorf("String() = %v, want JSON representing %v", got, tt.want)
			}
		})
	}
}

func TestMetadata_IsValid(t *testing.T) {
	tests := []struct {
		name string
		data map[string]any
	}{
		{"nil", nil},
		{"empty", map[string]any{}},
		{"valid data", map[string]any{"a": 1, "b": nil, "c": []int{1, 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Metadata{data: tt.data}
			if !m.IsValid() {
				t.Error("IsValid() returned false, should always be true")
			}
		})
	}
}

func TestMetadata_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		data    map[string]any
		want    []byte
		wantErr bool
	}{
		{
			name:    "nil map",
			data:    nil,
			want:    []byte("null"),
			wantErr: false,
		},
		{
			name:    "empty map",
			data:    map[string]any{},
			want:    []byte("{}"),
			wantErr: false,
		},
		{
			name:    "non-empty map",
			data:    map[string]any{"key": "value"},
			want:    []byte(`{"key":"value"}`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Metadata{data: tt.data}
			got, err := m.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Fatalf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalJSON() = %s, want %s", got, tt.want)
			}
		})
	}
}

func TestMetadata_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		wantData map[string]any
		wantErr  bool
	}{
		{
			name:     "null",
			input:    []byte("null"),
			wantData: nil,
			wantErr:  false,
		},
		{
			name:     "empty object",
			input:    []byte("{}"),
			wantData: map[string]any{},
			wantErr:  false,
		},
		{
			name:     "simple object",
			input:    []byte(`{"a":1,"b":"text"}`),
			wantData: map[string]any{"a": float64(1), "b": "text"}, // numbers become float64 in unmarshaled interface{}
			wantErr:  false,
		},
		{
			name:     "nested object",
			input:    []byte(`{"nested":{"x":true}}`),
			wantData: map[string]any{"nested": map[string]any{"x": true}},
			wantErr:  false,
		},
		{
			name:    "invalid JSON",
			input:   []byte(`{not json}`),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Metadata{}
			err := m.UnmarshalJSON(tt.input)
			if (err != nil) != tt.wantErr {
				t.Fatalf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && !reflect.DeepEqual(m.data, tt.wantData) {
				t.Errorf("UnmarshalJSON() data = %v, want %v", m.data, tt.wantData)
			}
		})
	}
}
