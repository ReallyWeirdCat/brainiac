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

var registrationCodeActiveWindowDuration = time.Minute * 10

type RegistrationCode struct {
	Username  valueobject.Username         `json:"username"`
	Code      valueobject.ConfirmationCode `json:"code"`
	CreatedAt time.Time                    `json:"created_at"`
	ExpireAt  time.Time                    `json:"expire_at"`
}

var _ Entity = RegistrationCode{}

func NewRegistrationCode(username valueobject.Username) (RegistrationCode, error) {
	code, err := valueobject.NewConfirmationCode()
	if err != nil {
		return RegistrationCode{}, err
	}
	expireAt := time.Now().Add(registrationCodeActiveWindowDuration)

	return RegistrationCode{
		Username:  username,
		Code:      code,
		CreatedAt: time.Now(),
		ExpireAt:  expireAt,
	}, nil
}

func (r RegistrationCode) IsValid() bool {
	now := time.Now()
	if now.After(r.ExpireAt) {
		return false
	}

	// Registraion code that exceeds max allowed duration is invalid
	if r.ExpireAt.Sub(r.CreatedAt) > registrationCodeActiveWindowDuration {
		return false
	}

	return r.Username.IsValid() && r.Code.IsValid()
}
