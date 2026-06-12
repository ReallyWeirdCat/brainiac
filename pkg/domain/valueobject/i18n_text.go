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
	"sort"

	"github.com/ReallyWeirdCat/brainiac/pkg/domain/errors"
)

// I18nText holds translations for multiple languages.
type I18nText struct {
	translations map[string]string
}

var _ ValueObject = I18nText{}

// NewI18nText creates an empty I18nText.
func NewI18nText() I18nText {
	return I18nText{
		translations: make(map[string]string),
	}
}

// I18nTextFromMap creates an I18nText from a raw map, validating all language codes.
func I18nTextFromMap(raw map[string]string) (I18nText, error) {
	t := NewI18nText()
	for langCode, text := range raw {
		lang, err := NewLanguageCode(langCode)
		if err != nil {
			return I18nText{}, errors.ErrInvalidLanguageCode
		}
		if text != "" {
			t.translations[lang.String()] = text
		}
	}
	return t, nil
}

// Set adds or updates a translation for the given language.
func (t *I18nText) Set(lang LanguageCode, text string) error {
	if lang.IsZero() {
		return errors.ErrInvalidLanguageCode
	}
	if text == "" {
		delete(t.translations, lang.String())
		return nil
	}
	t.translations[lang.String()] = text
	return nil
}

// Get returns the translation for a language and a boolean indicating presence.
func (t I18nText) Get(lang LanguageCode) (string, bool) {
	text, ok := t.translations[lang.String()]
	return text, ok
}

// All returns a copy of all translations as a map with string keys.
func (t I18nText) All() map[string]string {
	out := make(map[string]string, len(t.translations))
	maps.Copy(out, t.translations)
	return out
}

// SupportedLanguages returns all language codes for which a translation exists.
func (t I18nText) SupportedLanguages() []LanguageCode {
	codes := make([]LanguageCode, 0, len(t.translations))
	for codeStr := range t.translations {
		// Code is guaranteed valid because it came from NewLanguageCode in constructor/setters.
		lang, _ := NewLanguageCode(codeStr)
		codes = append(codes, lang)
	}
	// Sort for deterministic output.
	sort.Slice(codes, func(i, j int) bool {
		return codes[i].String() < codes[j].String()
	})
	return codes
}

// IsEmpty returns true if there are no translations.
func (t I18nText) IsEmpty() bool {
	return len(t.translations) == 0
}

// Equals compares two I18nText objects by their translations.
func (t I18nText) Equals(other any) bool {
	otherText, ok := other.(I18nText)
	if !ok {
		return false
	}
	if len(t.translations) != len(otherText.translations) {
		return false
	}
	for k, v := range t.translations {
		if ov, ok := otherText.translations[k]; !ok || ov != v {
			return false
		}
	}
	return true
}

// String returns a compact JSON representation of the translations.
func (t I18nText) String() string {
	if t.IsEmpty() {
		return "{}"
	}
	// Use JSON for readability; errors are ignored because map is always marshalable.
	b, _ := json.Marshal(t.translations)
	return string(b)
}

// IsValid returns true if all language codes are valid and no empty keys exist.
// Since construction and setters enforce validity, this always returns true.
func (t I18nText) IsValid() bool {
	for codeStr := range t.translations {
		if _, err := NewLanguageCode(codeStr); err != nil {
			return false
		}
	}
	return true
}

// IsZero returns true for an uninitialized I18nText (nil map) or an empty one.
// Zero value (I18nText{}) has nil map, which is considered empty.
func (t I18nText) IsZero() bool {
	return len(t.translations) == 0
}

// MarshalJSON implements json.Marshaler for seamless JSONB serialization.
func (t I18nText) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.translations)
}

// UnmarshalJSON implements json.Unmarshaler, validating language codes.
func (t *I18nText) UnmarshalJSON(data []byte) error {
	var raw map[string]string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	validated := make(map[string]string, len(raw))
	for langCode, text := range raw {
		lang, err := NewLanguageCode(langCode)
		if err != nil {
			return errors.ErrInvalidLanguageCode
		}
		if text != "" {
			validated[lang.String()] = text
		}
	}
	t.translations = validated
	return nil
}
