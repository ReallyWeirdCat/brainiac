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

func TestNewI18nText(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    I18nText
		wantErr bool
	}{
		{
			name:    "valid single language",
			args:    args{data: []byte(`{"en":"Hello"}`)},
			want:    I18nText([]byte(`{"en":"Hello"}`)),
			wantErr: false,
		},
		{
			name:    "valid uppercase language code",
			args:    args{data: []byte(`{"EN":"Hello"}`)},
			want:    I18nText([]byte(`{"EN":"Hello"}`)),
			wantErr: false,
		},
		{
			name:    "valid multiple languages",
			args:    args{data: []byte(`{"en":"Hello","fr":"Bonjour"}`)},
			want:    I18nText([]byte(`{"en":"Hello","fr":"Bonjour"}`)),
			wantErr: false,
		},
		{
			name:    "invalid language code - numeric",
			args:    args{data: []byte(`{"e1":"Hello"}`)},
			want:    I18nText{},
			wantErr: true,
		},
		{
			name:    "invalid language code - too long",
			args:    args{data: []byte(`{"eng":"Hello"}`)},
			want:    I18nText{},
			wantErr: true,
		},
		{
			name:    "invalid JSON",
			args:    args{data: []byte(`invalid`)},
			want:    I18nText{},
			wantErr: true,
		},
		{
			name:    "empty object",
			args:    args{data: []byte(`{}`)},
			want:    I18nText([]byte(`{}`)),
			wantErr: false,
		},
		{
			name:    "array instead of object",
			args:    args{data: []byte(`["en"]`)},
			want:    I18nText{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewI18nText(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Fatalf("NewI18nText() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewI18nText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewI18nTextFromMap(t *testing.T) {
	type args struct {
		data map[string]any
	}
	tests := []struct {
		name    string
		args    args
		want    I18nText
		wantErr bool
	}{
		{
			name:    "valid single entry",
			args:    args{data: map[string]any{"en": "Hello"}},
			want:    I18nText([]byte(`{"en":"Hello"}`)),
			wantErr: false,
		},
		{
			name:    "valid multiple entries",
			args:    args{data: map[string]any{"en": "Hello", "fr": "Bonjour"}},
			want:    I18nText([]byte(`{"en":"Hello","fr":"Bonjour"}`)),
			wantErr: false,
		},
		{
			name:    "invalid language code",
			args:    args{data: map[string]any{"123": "test"}},
			want:    I18nText{},
			wantErr: true,
		},
		{
			name:    "empty map",
			args:    args{data: map[string]any{}},
			want:    I18nText([]byte(`{}`)),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewI18nTextFromMap(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Fatalf("NewI18nTextFromMap() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewI18nTextFromMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestI18nText_Equals(t *testing.T) {
	type args struct {
		other any
	}
	tests := []struct {
		name string
		i    I18nText
		args args
		want bool
	}{
		{
			name: "equal identical",
			i:    I18nText([]byte(`{"en":"Hello"}`)),
			args: args{other: I18nText([]byte(`{"en":"Hello"}`))},
			want: true,
		},
		{
			name: "equal different whitespace",
			i:    I18nText([]byte(`{"en": "Hello"}`)),
			args: args{other: I18nText([]byte(`{"en":"Hello"}`))},
			want: true,
		},
		{
			name: "equal key order",
			i:    I18nText([]byte(`{"en":"Hello","fr":"Bonjour"}`)),
			args: args{other: I18nText([]byte(`{"fr":"Bonjour","en":"Hello"}`))},
			want: true,
		},
		{
			name: "different values",
			i:    I18nText([]byte(`{"en":"Hello"}`)),
			args: args{other: I18nText([]byte(`{"en":"Hi"}`))},
			want: false,
		},
		{
			name: "other nil",
			i:    I18nText([]byte(`{}`)),
			args: args{other: nil},
			want: false,
		},
		{
			name: "other wrong type",
			i:    I18nText([]byte(`{}`)),
			args: args{other: "not I18nText"},
			want: false,
		},
		{
			name: "both invalid JSON",
			i:    I18nText([]byte(`invalid`)),
			args: args{other: I18nText([]byte(`invalid`))},
			want: true,
		},
		{
			name: "one invalid JSON",
			i:    I18nText([]byte(`invalid`)),
			args: args{other: I18nText([]byte(`{}`))},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Equals(tt.args.other); got != tt.want {
				t.Errorf("I18nText.Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestI18nText_IsValid(t *testing.T) {
	tests := []struct {
		name string
		i    I18nText
		want bool
	}{
		{
			name: "valid object with codes",
			i:    I18nText([]byte(`{"en":"Hello"}`)),
			want: true,
		},
		{
			name: "valid empty object",
			i:    I18nText([]byte(`{}`)),
			want: true,
		},
		{
			name: "invalid language code",
			i:    I18nText([]byte(`{"eng":"Hello"}`)),
			want: false,
		},
		{
			name: "invalid JSON",
			i:    I18nText([]byte(`{bad`)),
			want: false,
		},
		{
			name: "empty byte slice",
			i:    I18nText{},
			want: false,
		},
		{
			name: "array instead of object",
			i:    I18nText([]byte(`["en"]`)),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.IsValid(); got != tt.want {
				t.Errorf("I18nText.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestI18nText_IsZero(t *testing.T) {
	tests := []struct {
		name string
		i    I18nText
		want bool
	}{
		{
			name: "zero length",
			i:    I18nText{},
			want: true,
		},
		{
			name: "empty object",
			i:    I18nText([]byte(`{}`)),
			want: false,
		},
		{
			name: "non-empty",
			i:    I18nText([]byte(`{"en":"Hello"}`)),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.IsZero(); got != tt.want {
				t.Errorf("I18nText.IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestI18nText_String(t *testing.T) {
	tests := []struct {
		name string
		i    I18nText
		want string
	}{
		{
			name: "empty",
			i:    I18nText{},
			want: "",
		},
		{
			name: "simple",
			i:    I18nText([]byte(`{"en":"Hello"}`)),
			want: `{"en":"Hello"}`,
		},
		{
			name: "with spaces",
			i:    I18nText([]byte(`{  }`)),
			want: `{  }`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.String(); got != tt.want {
				t.Errorf("I18nText.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_canonicalizeI18n(t *testing.T) {
	type args struct {
		i I18nText
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name:    "compact and sorted keys",
			args:    args{i: I18nText([]byte(`{"fr": "Bonjour", "en": "Hello" }`))},
			want:    []byte(`{"en":"Hello","fr":"Bonjour"}`),
			wantErr: false,
		},
		{
			name:    "invalid JSON",
			args:    args{i: I18nText([]byte(`not json`))},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "empty byte slice",
			args:    args{i: I18nText{}},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := canonicalizeI18n(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Fatalf("canonicalizeI18n() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("canonicalizeI18n() = %v, want %v", got, tt.want)
			}
		})
	}
}
