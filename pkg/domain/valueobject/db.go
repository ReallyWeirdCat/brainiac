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

import "database/sql/driver"

// The value objects below are also used as sqlc/pgx wire types (see sqlc.yaml
// column overrides). Implementing driver.Valuer lets pgx encode them through
// its native driver.Valuer path instead of the reflection-based
// "underlying type" fallback on writes.
//
// Note: the decode (scan) path in pgx v5 still uses its reflection fallback for
// these named types unless a pgtype codec is registered; see ARCHITECTURE_REVIEW.md.

func (g GUID) Value() (driver.Value, error)         { return string(g), nil }
func (u Username) Value() (driver.Value, error)     { return string(u), nil }
func (e Email) Value() (driver.Value, error)        { return string(e), nil }
func (n Name) Value() (driver.Value, error)         { return string(n), nil }
func (n Nickname) Value() (driver.Value, error)     { return string(n), nil }
func (b Bio) Value() (driver.Value, error)          { return string(b), nil }
func (l LanguageCode) Value() (driver.Value, error) { return string(l), nil }
func (h HttpUrl) Value() (driver.Value, error)      { return string(h), nil }
func (m Metadata) Value() (driver.Value, error)     { return []byte(m), nil }
func (i I18nText) Value() (driver.Value, error)     { return []byte(i), nil }

var (
	_ driver.Valuer = GUID("")
	_ driver.Valuer = Username("")
	_ driver.Valuer = Email("")
	_ driver.Valuer = Name("")
	_ driver.Valuer = Nickname("")
	_ driver.Valuer = Bio("")
	_ driver.Valuer = LanguageCode("")
	_ driver.Valuer = HttpUrl("")
	_ driver.Valuer = Metadata{}
	_ driver.Valuer = I18nText{}
)
