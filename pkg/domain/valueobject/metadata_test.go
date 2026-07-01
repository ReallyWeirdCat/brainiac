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
	"reflect"
	"testing"
)

func TestNewMetadata(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    Metadata
		wantErr bool
	}{
		{
			name:    "valid empty object",
			args:    args{data: []byte(`{}`)},
			want:    Metadata([]byte(`{}`)),
			wantErr: false,
		},
		{
			name:    "valid object with fields",
			args:    args{data: []byte(`{"key":"value"}`)},
			want:    Metadata([]byte(`{"key":"value"}`)),
			wantErr: false,
		},
		{
			name:    "invalid json",
			args:    args{data: []byte(`{invalid}`)},
			want:    Metadata{},
			wantErr: true,
		},
		{
			name:    "empty byte slice",
			args:    args{data: []byte{}},
			want:    Metadata{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewMetadata(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Fatalf("NewMetadata() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMetadata() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewMetadataFromMap(t *testing.T) {
	type args struct {
		data map[string]any
	}
	tests := []struct {
		name    string
		args    args
		want    Metadata
		wantErr bool
	}{
		{
			name:    "empty map",
			args:    args{data: map[string]any{}},
			want:    Metadata([]byte(`{}`)),
			wantErr: false,
		},
		{
			name:    "simple map",
			args:    args{data: map[string]any{"a": 1, "b": "text"}},
			want:    Metadata([]byte(`{"a":1,"b":"text"}`)),
			wantErr: false,
		},
		{
			name:    "map with nested data",
			args:    args{data: map[string]any{"nested": map[string]any{"x": true}}},
			want:    Metadata([]byte(`{"nested":{"x":true}}`)),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewMetadataFromMap(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Fatalf("NewMetadataFromMap() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMetadataFromMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMetadata_Equals(t *testing.T) {
	type args struct {
		other any
	}
	tests := []struct {
		name string
		m    Metadata
		args args
		want bool
	}{
		{
			name: "equal identical",
			m:    Metadata([]byte(`{"a":1}`)),
			args: args{other: Metadata([]byte(`{"a":1}`))},
			want: true,
		},
		{
			name: "equal different whitespace",
			m:    Metadata([]byte(`{"a": 1}`)),
			args: args{other: Metadata([]byte(`{"a":1}`))},
			want: true,
		},
		{
			name: "equal key order",
			m:    Metadata([]byte(`{"a":1,"b":2}`)),
			args: args{other: Metadata([]byte(`{"b":2,"a":1}`))},
			want: true,
		},
		{
			name: "different values",
			m:    Metadata([]byte(`{"a":1}`)),
			args: args{other: Metadata([]byte(`{"a":2}`))},
			want: false,
		},
		{
			name: "other nil",
			m:    Metadata([]byte(`{}`)),
			args: args{other: nil},
			want: false,
		},
		{
			name: "other wrong type",
			m:    Metadata([]byte(`{}`)),
			args: args{other: "not metadata"},
			want: false,
		},
		{
			name: "both invalid json",
			m:    Metadata([]byte(`invalid`)),
			args: args{other: Metadata([]byte(`invalid`))},
			want: true,
		},
		{
			name: "one invalid json",
			m:    Metadata([]byte(`invalid`)),
			args: args{other: Metadata([]byte(`{}`))},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Equals(tt.args.other); got != tt.want {
				t.Errorf("Metadata.Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMetadata_IsValid(t *testing.T) {
	tests := []struct {
		name string
		m    Metadata
		want bool
	}{
		{
			name: "valid object",
			m:    Metadata([]byte(`{"a":1}`)),
			want: true,
		},
		{
			name: "invalid json",
			m:    Metadata([]byte(`{bad}`)),
			want: false,
		},
		{
			name: "empty byte slice",
			m:    Metadata{},
			want: false,
		},
		{
			name: "empty object",
			m:    Metadata([]byte(`{}`)),
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.IsValid(); got != tt.want {
				t.Errorf("Metadata.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMetadata_IsZero(t *testing.T) {
	tests := []struct {
		name string
		m    Metadata
		want bool
	}{
		{
			name: "zero length",
			m:    Metadata{},
			want: true,
		},
		{
			name: "empty object",
			m:    Metadata([]byte(`{}`)),
			want: false,
		},
		{
			name: "non-empty",
			m:    Metadata([]byte(`{"a":1}`)),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.IsZero(); got != tt.want {
				t.Errorf("Metadata.IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMetadata_String(t *testing.T) {
	tests := []struct {
		name string
		m    Metadata
		want string
	}{
		{
			name: "empty",
			m:    Metadata{},
			want: "",
		},
		{
			name: "simple",
			m:    Metadata([]byte(`{"key":"value"}`)),
			want: `{"key":"value"}`,
		},
		{
			name: "with spaces",
			m:    Metadata([]byte(`{  }`)),
			want: `{  }`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.String(); got != tt.want {
				t.Errorf("Metadata.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_canonicalize(t *testing.T) {
	type args struct {
		m Metadata
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name:    "valid compact",
			args:    args{m: Metadata([]byte(`{"a": 1 }`))},
			want:    []byte(`{"a":1}`),
			wantErr: false,
		},
		{
			name:    "invalid json",
			args:    args{m: Metadata([]byte(`not json`))},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "empty json not allowed",
			args:    args{m: Metadata{}},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := canonicalize(tt.args.m)
			if (err != nil) != tt.wantErr {
				t.Fatalf("canonicalize() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("canonicalize() = %v, want %v", got, tt.want)
			}
		})
	}
}
