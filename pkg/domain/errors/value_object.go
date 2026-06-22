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

package errors

var (
	ErrInvalidUsername     = NewDomainError("invalid username format", nil).WithType(Validation)
	ErrInvalidEmail        = NewDomainError("invalid email format", nil).WithType(Validation)
	ErrInvalidGUID         = NewDomainError("invalid GUID format", nil).WithType(Validation)
	ErrInvalidName         = NewDomainError("invalid name format", nil).WithType(Validation)
	ErrInvalidNickname     = NewDomainError("invalid nickname format", nil).WithType(Validation)
	ErrInvalidBio          = NewDomainError("invalid bio format", nil).WithType(Validation)
	ErrInvalidLanguageCode = NewDomainError("invalid language code format", nil).WithType(Validation)
	ErrInvalidHttpUrl      = NewDomainError("invalid HTTP URL format", nil).WithType(Validation)
)
