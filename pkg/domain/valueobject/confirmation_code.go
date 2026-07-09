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
	"crypto/rand"
	"fmt"
	"math/big"
	"regexp"

	domerr "github.com/ReallyWeirdCat/brainiac/pkg/domain/errors"
)

var confirmationCodeRegex = regexp.MustCompile(`^[0-9]{6}$`)
var ErrInvalidConfirmationCode = domerr.NewDomainError("invalid confirmation code", nil).WithType(domerr.Validation)

func generateCode() (string, error) {
	max := big.NewInt(1_000_000)
	for {
		n, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "", err
		}
		code := fmt.Sprintf("%06d", n.Int64())
		if validateCode(code) {
			return code, nil
		}
	}
}

func validateCode(code string) bool {
	return confirmationCodeRegex.MatchString(code)
}

type ConfirmationCode string

func NewConfirmationCode() (ConfirmationCode, error) {
	code, err := generateCode()
	return ConfirmationCode(code), err
}

func ConfirmationCodeFromString(code string) (ConfirmationCode, error) {
	if !validateCode(code) {
		return ConfirmationCode(""), ErrInvalidConfirmationCode
	}
	return ConfirmationCode(code), nil
}

func (c ConfirmationCode) Equals(other any) bool {
	o, ok := other.(ConfirmationCode)
	return ok && c == o
}

func (c ConfirmationCode) IsValid() bool {
	s := string(c)

	return validateCode(s)
}

func (c ConfirmationCode) IsZero() bool {
	return string(c) == ""
}

func (c ConfirmationCode) String() string {
	return string(c)
}

var _ ValueObject = ConfirmationCode("")
