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
	"maps"
	"reflect"
)

// Metadata holds arbitrary key‑value data, stored as JSONB in the database.
type Metadata struct {
	data map[string]any
}

var _ ValueObject = Metadata{}

// NewMetadata creates an empty metadata container.
func NewMetadata() Metadata {
	return Metadata{
		data: make(map[string]any),
	}
}

// MetadataFromMap creates a Metadata from an existing map.
// The map is copied to avoid external mutations.
func MetadataFromMap(raw map[string]any) Metadata {
	if raw == nil {
		return NewMetadata()
	}
	copy := make(map[string]any, len(raw))
	maps.Copy(copy, raw)
	return Metadata{data: copy}
}

// Set adds or updates a key with any JSON‑serializable value.
func (m *Metadata) Set(key string, value any) {
	if m.data == nil {
		m.data = make(map[string]any)
	}
	m.data[key] = value
}

// Get returns the value for a key and a boolean indicating presence.
func (m Metadata) Get(key string) (any, bool) {
	if m.data == nil {
		return nil, false
	}
	val, ok := m.data[key]
	return val, ok
}

// Delete removes a key.
func (m Metadata) Delete(key string) {
	if m.data != nil {
		delete(m.data, key)
	}
}

// Keys returns all keys currently set.
func (m Metadata) Keys() []string {
	if m.data == nil {
		return []string{}
	}
	keys := make([]string, 0, len(m.data))
	for k := range m.data {
		keys = append(keys, k)
	}
	return keys
}

// AsMap returns a shallow copy of the underlying map.
func (m Metadata) AsMap() map[string]any {
	if m.data == nil {
		return nil
	}
	out := make(map[string]any, len(m.data))
	maps.Copy(out, m.data)
	return out
}

// IsEmpty returns true if no metadata is stored.
func (m Metadata) IsEmpty() bool {
	return len(m.data) == 0
}

// Equals compares two Metadata objects by deep JSON equality.
// This handles nested structures correctly.
func (m Metadata) Equals(other any) bool {
	otherMeta, ok := other.(Metadata)
	if !ok {
		return false
	}
	// If both are empty (nil or len==0), they are equal.
	if len(m.data) == 0 && len(otherMeta.data) == 0 {
		return true
	}
	// Otherwise compare by JSON.
	mJSON, err1 := json.Marshal(m.data)
	oJSON, err2 := json.Marshal(otherMeta.data)
	if err1 != nil || err2 != nil {
		return reflect.DeepEqual(m.data, otherMeta.data)
	}
	return string(mJSON) == string(oJSON)
}

// String returns a JSON representation of the metadata.
func (m Metadata) String() string {
	if m.IsEmpty() {
		return "{}"
	}
	b, _ := json.Marshal(m.data)
	return string(b)
}

// IsValid always returns true because any JSON‑serializable value is allowed.
func (m Metadata) IsValid() bool {
	return true
}

// IsZero returns true for an uninitialized Metadata (nil map) or an empty one.
func (m Metadata) IsZero() bool {
	return len(m.data) == 0
}

// MarshalJSON implements json.Marshaler for seamless JSONB serialization.
func (m Metadata) MarshalJSON() ([]byte, error) {
	if m.data == nil {
		return []byte("null"), nil
	}
	return json.Marshal(m.data)
}

// UnmarshalJSON implements json.Unmarshaler.
func (m *Metadata) UnmarshalJSON(data []byte) error {
	var raw map[string]any
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	m.data = raw
	return nil
}
