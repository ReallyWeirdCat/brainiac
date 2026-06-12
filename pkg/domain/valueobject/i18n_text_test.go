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

func TestNewI18nText(t *testing.T) {
	got := NewI18nText()
	if got.translations == nil {
		t.Error("NewI18nText() returned nil internal map")
	}
	if len(got.translations) != 0 {
		t.Errorf("NewI18nText() map length = %d, want 0", len(got.translations))
	}
	if !got.IsEmpty() {
		t.Error("NewI18nText() should be empty")
	}
}

func TestI18nTextFromMap(t *testing.T) {
	validEn, _ := NewLanguageCode("en")
	validFr, _ := NewLanguageCode("fr")
	tests := []struct {
		name    string
		raw     map[string]string
		want    I18nText
		wantErr bool
	}{
		{
			name:    "nil input",
			raw:     nil,
			want:    I18nText{translations: map[string]string{}},
			wantErr: false,
		},
		{
			name:    "empty map",
			raw:     map[string]string{},
			want:    I18nText{translations: map[string]string{}},
			wantErr: false,
		},
		{
			name: "valid map",
			raw:  map[string]string{"en": "Hello", "fr": "Bonjour"},
			want: func() I18nText {
				t := NewI18nText()
				t.Set(validEn, "Hello")
				t.Set(validFr, "Bonjour")
				return t
			}(),
			wantErr: false,
		},
		{
			name: "map with empty string value (should be omitted)",
			raw:  map[string]string{"en": "Hello", "fr": ""},
			want: func() I18nText {
				t := NewI18nText()
				t.Set(validEn, "Hello")
				return t
			}(),
			wantErr: false,
		},
		{
			name:    "invalid language code",
			raw:     map[string]string{"en": "Hello", "invalid": "Oops"},
			want:    I18nText{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := I18nTextFromMap(tt.raw)
			if (err != nil) != tt.wantErr {
				t.Fatalf("I18nTextFromMap() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if !reflect.DeepEqual(got.translations, tt.want.translations) {
				t.Errorf("I18nTextFromMap().translations = %v, want %v", got.translations, tt.want.translations)
			}
		})
	}
}

func TestI18nText_Set(t *testing.T) {
	validEn, _ := NewLanguageCode("en")
	zeroLang := LanguageCode{}
	tests := []struct {
		name    string
		initial map[string]string
		lang    LanguageCode
		text    string
		wantMap map[string]string
		wantErr bool
	}{
		{
			name:    "set new key on existing map",
			initial: map[string]string{"fr": "Bonjour"},
			lang:    validEn,
			text:    "Hello",
			wantMap: map[string]string{"fr": "Bonjour", "en": "Hello"},
			wantErr: false,
		},
		{
			name:    "overwrite existing key",
			initial: map[string]string{"en": "Hi"},
			lang:    validEn,
			text:    "Hello",
			wantMap: map[string]string{"en": "Hello"},
			wantErr: false,
		},
		{
			name:    "delete key with empty string",
			initial: map[string]string{"en": "Hello", "fr": "Bonjour"},
			lang:    validEn,
			text:    "",
			wantMap: map[string]string{"fr": "Bonjour"},
			wantErr: false,
		},
		{
			name:    "delete non-existing key does nothing",
			initial: map[string]string{"fr": "Bonjour"},
			lang:    validEn,
			text:    "",
			wantMap: map[string]string{"fr": "Bonjour"},
			wantErr: false,
		},
		{
			name:    "zero language code error",
			initial: map[string]string{},
			lang:    zeroLang,
			text:    "text",
			wantMap: map[string]string{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := I18nText{translations: tt.initial}
			err := tr.Set(tt.lang, tt.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("Set() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tr.translations, tt.wantMap) {
				t.Errorf("After Set, translations = %v, want %v", tr.translations, tt.wantMap)
			}
		})
	}
}

func TestI18nText_Get(t *testing.T) {
	validEn, _ := NewLanguageCode("en")
	validFr, _ := NewLanguageCode("fr")
	tests := []struct {
		name         string
		translations map[string]string
		lang         LanguageCode
		wantText     string
		wantOk       bool
	}{
		{
			name:         "nil map",
			translations: nil,
			lang:         validEn,
			wantText:     "",
			wantOk:       false,
		},
		{
			name:         "key exists",
			translations: map[string]string{"en": "Hello", "fr": "Bonjour"},
			lang:         validEn,
			wantText:     "Hello",
			wantOk:       true,
		},
		{
			name:         "key does not exist",
			translations: map[string]string{"en": "Hello"},
			lang:         validFr,
			wantText:     "",
			wantOk:       false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := I18nText{translations: tt.translations}
			got, ok := tr.Get(tt.lang)
			if got != tt.wantText {
				t.Errorf("Get() text = %v, want %v", got, tt.wantText)
			}
			if ok != tt.wantOk {
				t.Errorf("Get() ok = %v, want %v", ok, tt.wantOk)
			}
		})
	}
}

func TestI18nText_All(t *testing.T) {
	tests := []struct {
		name         string
		translations map[string]string
		want         map[string]string
	}{
		{
			name:         "nil map",
			translations: nil,
			want:         map[string]string{},
		},
		{
			name:         "empty map",
			translations: map[string]string{},
			want:         map[string]string{},
		},
		{
			name:         "non-empty map",
			translations: map[string]string{"en": "Hello", "fr": "Bonjour"},
			want:         map[string]string{"en": "Hello", "fr": "Bonjour"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := I18nText{translations: tt.translations}
			got := tr.All()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("All() = %v, want %v", got, tt.want)
			}
			// Verify copy: modify returned map, original unchanged.
			if len(got) > 0 {
				got["new"] = "should not appear"
				if _, exists := tr.translations["new"]; exists {
					t.Error("All() did not return a copy; modification affected original")
				}
			}
		})
	}
}

func TestI18nText_SupportedLanguages(t *testing.T) {
	en, _ := NewLanguageCode("en")
	fr, _ := NewLanguageCode("fr")
	de, _ := NewLanguageCode("de")
	tests := []struct {
		name         string
		translations map[string]string
		want         []LanguageCode
	}{
		{
			name:         "nil map",
			translations: nil,
			want:         []LanguageCode{},
		},
		{
			name:         "empty map",
			translations: map[string]string{},
			want:         []LanguageCode{},
		},
		{
			name:         "single language",
			translations: map[string]string{"en": "Hello"},
			want:         []LanguageCode{en},
		},
		{
			name:         "multiple languages (sorted)",
			translations: map[string]string{"fr": "Bonjour", "de": "Hallo", "en": "Hello"},
			want:         []LanguageCode{de, en, fr}, // sorted by code string
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := I18nText{translations: tt.translations}
			got := tr.SupportedLanguages()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SupportedLanguages() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestI18nText_IsEmpty(t *testing.T) {
	tests := []struct {
		name         string
		translations map[string]string
		want         bool
	}{
		{"nil map", nil, true},
		{"empty map", map[string]string{}, true},
		{"non-empty map", map[string]string{"en": "Hello"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := I18nText{translations: tt.translations}
			if got := tr.IsEmpty(); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestI18nText_Equals(t *testing.T) {
	tests := []struct {
		name  string
		tr    I18nText
		other any
		want  bool
	}{
		{
			name:  "equal empty maps (nil vs empty)",
			tr:    I18nText{translations: nil},
			other: I18nText{translations: map[string]string{}},
			want:  true,
		},
		{
			name:  "both nil",
			tr:    I18nText{translations: nil},
			other: I18nText{translations: nil},
			want:  true,
		},
		{
			name:  "equal non-empty",
			tr:    I18nText{translations: map[string]string{"en": "Hello", "fr": "Bonjour"}},
			other: I18nText{translations: map[string]string{"fr": "Bonjour", "en": "Hello"}},
			want:  true,
		},
		{
			name:  "different values",
			tr:    I18nText{translations: map[string]string{"en": "Hello"}},
			other: I18nText{translations: map[string]string{"en": "Hi"}},
			want:  false,
		},
		{
			name:  "different keys",
			tr:    I18nText{translations: map[string]string{"en": "Hello"}},
			other: I18nText{translations: map[string]string{"fr": "Bonjour"}},
			want:  false,
		},
		{
			name:  "different type",
			tr:    I18nText{translations: map[string]string{}},
			other: "not an I18nText",
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.Equals(tt.other); got != tt.want {
				t.Errorf("Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestI18nText_String(t *testing.T) {
	tests := []struct {
		name         string
		translations map[string]string
		want         string
	}{
		{"nil map", nil, "{}"},
		{"empty map", map[string]string{}, "{}"},
		{"single key", map[string]string{"en": "Hello"}, `{"en":"Hello"}`},
		{"multiple keys", map[string]string{"en": "Hello", "fr": "Bonjour"}, `{"en":"Hello","fr":"Bonjour"}`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := I18nText{translations: tt.translations}
			got := tr.String()
			// Compare by unmarshalling to avoid ordering issues.
			var gotMap, wantMap map[string]string
			if err := json.Unmarshal([]byte(got), &gotMap); err != nil {
				t.Fatalf("String() returned invalid JSON: %v", err)
			}
			if err := json.Unmarshal([]byte(tt.want), &wantMap); err != nil {
				t.Fatalf("want string is invalid JSON: %v", err)
			}
			if !reflect.DeepEqual(gotMap, wantMap) {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestI18nText_IsValid(t *testing.T) {
	// Valid case: all keys are valid language codes.
	valid := I18nText{translations: map[string]string{"en": "Hello", "fr": "Bonjour"}}
	if !valid.IsValid() {
		t.Error("IsValid() should return true for valid language codes")
	}
	// Invalid case: manually construct with invalid key (bypassing constructor).
	invalid := I18nText{translations: map[string]string{"invalid": "Oops"}}
	if invalid.IsValid() {
		t.Error("IsValid() should return false when language code is invalid")
	}
	// Nil or empty map is considered valid (no entries to validate).
	emptyNil := I18nText{translations: nil}
	if !emptyNil.IsValid() {
		t.Error("IsValid() should return true for nil map")
	}
}

func TestI18nText_IsZero(t *testing.T) {
	tests := []struct {
		name         string
		translations map[string]string
		want         bool
	}{
		{"nil map", nil, true},
		{"empty map", map[string]string{}, true},
		{"non-empty map", map[string]string{"en": "Hello"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := I18nText{translations: tt.translations}
			if got := tr.IsZero(); got != tt.want {
				t.Errorf("IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestI18nText_MarshalJSON(t *testing.T) {
	tests := []struct {
		name         string
		translations map[string]string
		want         []byte
		wantErr      bool
	}{
		{
			name:         "empty map",
			translations: map[string]string{},
			want:         []byte("{}"),
			wantErr:      false,
		},
		{
			name:         "non-empty map",
			translations: map[string]string{"en": "Hello"},
			want:         []byte(`{"en":"Hello"}`),
			wantErr:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := I18nText{translations: tt.translations}
			got, err := tr.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Fatalf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalJSON() = %s, want %s", got, tt.want)
			}
		})
	}
}

func TestI18nText_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantMap map[string]string
		wantErr bool
	}{
		{
			name:    "empty object",
			input:   []byte("{}"),
			wantMap: map[string]string{},
			wantErr: false,
		},
		{
			name:    "valid object",
			input:   []byte(`{"en":"Hello","fr":"Bonjour"}`),
			wantMap: map[string]string{"en": "Hello", "fr": "Bonjour"},
			wantErr: false,
		},
		{
			name:    "invalid language code key",
			input:   []byte(`{"en":"Hello","invalid":"Oops"}`),
			wantErr: true,
		},
		{
			name:    "invalid JSON",
			input:   []byte(`{not json}`),
			wantErr: true,
		},
		{
			name:    "non-string value (should fail because map[string]string)",
			input:   []byte(`{"en":123}`),
			wantErr: true, // JSON unmarshal will fail because 123 is not string
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &I18nText{}
			err := tr.UnmarshalJSON(tt.input)
			if (err != nil) != tt.wantErr {
				t.Fatalf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && !reflect.DeepEqual(tr.translations, tt.wantMap) {
				t.Errorf("UnmarshalJSON() translations = %v, want %v", tr.translations, tt.wantMap)
			}
		})
	}
}
