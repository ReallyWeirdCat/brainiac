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

import "fmt"

type DomainErrorType int16

const (
	Unspecified DomainErrorType = iota
	Validation
	NotFound
	Conflict
	Forbidden
	Unauthorized
	InvalidState
	LimitExceeded
	PreconditionFailed
)

func (d DomainErrorType) String() string {
	switch d {
	case Unspecified:
		return "unspecified"
	case Validation:
		return "validation"
	case NotFound:
		return "not_found"
	case Conflict:
		return "conflict"
	case Forbidden:
		return "forbidden"
	case Unauthorized:
		return "unauthorized"
	case InvalidState:
		return "invalid_state"
	case LimitExceeded:
		return "limit_exceeded"
	case PreconditionFailed:
		return "precondition_failed"
	default:
		return "unknown"
	}
}

type DomainError struct {
	errtype DomainErrorType
	message string
	err     error
}

var _ error = DomainError{}

func NewDomainError(message string, err error) DomainError {
	return DomainError{errtype: Unspecified, message: message, err: err}
}

func (d DomainError) WithType(error_type DomainErrorType) DomainError {
	clone := d
	clone.errtype = error_type
	return clone
}

func (d DomainError) DomainErrorType() DomainErrorType {
	return d.errtype
}

func (d DomainError) Error() string {
	if d.err != nil {
		return fmt.Sprintf("domain %s error: %s: %v", d.errtype.String(), d.message, d.err)
	}
	return fmt.Sprintf("domain %s error: %s", d.errtype.String(), d.message)
}

func (e DomainError) Is(target error) bool {
	if t, ok := target.(DomainError); ok {
		return e.errtype == t.errtype
	}
	return false
}

func (d DomainError) Unwrap() error {
	return d.err
}
