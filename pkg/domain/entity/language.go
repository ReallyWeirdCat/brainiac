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

package entity

import (
	"time"

	"github.com/ReallyWeirdCat/brainiac/pkg/domain/valueobject"
)

type Language struct {
	GUID         valueobject.GUID         `json:"guid"`
	LanguageCode valueobject.LanguageCode `json:"language_code"`
	EnglishTitle string                   `json:"english_title"`
	LocalTitle   string                   `json:"local_title"`
	Emoji        *string                  `json:"emoji,omitempty"`
	IsSupported  bool                     `json:"is_supported"`
	IsBeta       bool                     `json:"is_beta"`
	CreatedAt    time.Time                `json:"created_at"`
	DeletedAt    *time.Time               `json:"deleted_at,omitempty"`
}

var _ Entity = Language{}

func (l Language) IsValid() bool {

	if len(l.EnglishTitle) > 25 {
		return false
	}
	if len(l.LocalTitle) > 50 {
		return false
	}
	if l.GUID == nil {
		return false
	}
	return l.GUID.IsValid() && l.LanguageCode.IsValid()
}
