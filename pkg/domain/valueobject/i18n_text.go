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
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	domerr "github.com/ReallyWeirdCat/brainiac/pkg/domain/errors"
)

var ErrInvalidI18nText = domerr.NewDomainError("invalid I18nText", nil).WithType(domerr.Validation)

type I18nText json.RawMessage

var _ ValueObject = I18nText([]byte("{}"))

func NewI18nText(data []byte) (I18nText, error) {
	if !json.Valid(data) {
		return I18nText{}, ErrInvalidI18nText.FromError(errors.New("invalid json"))
	}
	var m map[string]json.RawMessage
	if err := json.Unmarshal(data, &m); err != nil {
		return I18nText{}, ErrInvalidI18nText.FromError(errors.New("I18nText must be a JSON object"))
	}
	for key := range m {
		if _, err := NewLanguageCode(key); err != nil {
			return I18nText{}, ErrInvalidI18nText.FromError(fmt.Errorf("invalid language code %q: %w", key, err))
		}
	}
	return I18nText(data), nil
}

func NewI18nTextFromMap(data map[string]any) (I18nText, error) {
	for key := range data {
		if _, err := NewLanguageCode(key); err != nil {
			return I18nText{}, ErrInvalidI18nText.FromError(fmt.Errorf("invalid language code %q: %w", key, err))
		}
	}
	b, err := json.Marshal(data)
	if err != nil {
		return I18nText{}, err
	}
	return I18nText(b), nil
}

func (i I18nText) Equals(other any) bool {
	otherI, ok := other.(I18nText)
	if !ok {
		return false
	}

	canon1, err1 := canonicalizeI18n(i)
	canon2, err2 := canonicalizeI18n(otherI)
	if err1 != nil || err2 != nil {
		return bytes.Equal(i, otherI)
	}
	return bytes.Equal(canon1, canon2)
}

func (i I18nText) IsValid() bool {
	if !json.Valid(i) {
		return false
	}
	if len(i) == 0 {
		return false
	}
	var m map[string]json.RawMessage
	if err := json.Unmarshal(i, &m); err != nil {
		return false
	}
	for key := range m {
		if _, err := NewLanguageCode(key); err != nil {
			return false
		}
	}
	return true
}

func (i I18nText) IsZero() bool {
	return len(i) == 0
}

func (i I18nText) String() string {
	return string(i)
}

func canonicalizeI18n(i I18nText) ([]byte, error) {
	if len(i) == 0 {
		return nil, ErrInvalidI18nText.FromError(errors.New("nil or empty I18nText"))
	}
	var obj any
	dec := json.NewDecoder(bytes.NewReader(i))
	dec.UseNumber()
	if err := dec.Decode(&obj); err != nil {
		return nil, err
	}
	return json.Marshal(obj)
}
