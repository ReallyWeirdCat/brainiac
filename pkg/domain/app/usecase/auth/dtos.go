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

package auth

// RegisterRequest contains the data needed to register a new user.
type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// RegisterResponse returns the created user's identifier and basic info.
type RegisterResponse struct {
	GUID     string `json:"guid"`
	Username string `json:"username"`
}

// LoginRequest contains credentials for logging in.
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginSuccessResponse returns the authenticated user's info and a session token.
type LoginSuccessResponse struct {
	GUID     string `json:"guid"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

// LoginTOTPRequiredResponse returns a temporary token needed for TOTP confirmation
type LoginTOTPRequiredResponse struct {
	TempTOTPToken string `json:"temp_totp_token"`
}

// ConfirmTOTPRequest is a struct for requesting TOTP confirmation.
type ConfirmTOTPRequest struct {
	TOTPCode string `json:"totp_code"`
}

// ConfirmTOTPResponse is a struct representing the response to confirming TOTP.
type ConfirmTOTPResponse struct {
	GUID     string `json:"guid"`
	Username string `json:"username"`
	Token    string `json:"token"`
}
