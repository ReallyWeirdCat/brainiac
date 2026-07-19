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

package ports

import "time"

const (
	ScopeAccess  = "a"
	ScopeRefresh = "r"
	ScopeTOTP    = "t"
)

type TokenGenerator interface {
	IssueAccessToken(userGUID, sessionGUID string) (string, error)
	IssueRefreshToken(userGUID, sessionGUID string) (string, error)
	IssueTOTPToken(userGUID string) (string, error)
}

type TokenValidator interface {
	ValidateAccessToken(tokenString string) (*TokenClaims, error)
	ValidateRefreshToken(tokenString string) (*TokenClaims, error)
	ValidateTOTPToken(tokenString string) (*TokenClaims, error)
}

type TokenProvider interface {
	TokenGenerator
	TokenValidator
}

// TokenClaims represents claims of a JWT token.
type TokenClaims struct {
	TokenGUID   string    `json:"jti"`
	SessionGUID string    `json:"sid"`
	UserGUID    string    `json:"sub"`
	Scope       string    `json:"scope"`
	IssuedAt    time.Time `json:"iat"`
	ExpiresAt   time.Time `json:"exp"`
}
