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
	"testing"

	domerr "github.com/ReallyWeirdCat/brainiac/pkg/domain/errors"
)

func Test_generateCode(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "successful code generation",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := generateCode()
			if (err != nil) != tt.wantErr {
				t.Fatalf("generateCode() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if len(got) != 6 {
				t.Errorf("generateCode() length = %d, want 6", len(got))
			}
			if !validateCode(got) {
				t.Errorf("generateCode() produced invalid code: %s", got)
			}
		})
	}
}

func Test_validateCode(t *testing.T) {
	tests := []struct {
		name string
		code string
		want bool
	}{
		{"valid code - mix of digits", "123457", true},
		{"valid code - starting with zeros", "000001", true},
		{"valid code - random digits", "482930", true},
		{"invalid - too short (5 digits)", "12345", false},
		{"invalid - too long (7 digits)", "1234567", false},
		{"invalid - empty string", "", false},
		{"invalid - letters", "abcdef", false},
		{"invalid - alphanumeric", "12345a", false},
		{"invalid - special characters", "12@456", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateCode(tt.code); got != tt.want {
				t.Errorf("validateCode(%q) = %v, want %v", tt.code, got, tt.want)
			}
		})
	}
}

func TestNewConfirmationCode(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "creates valid code",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewConfirmationCode()
			if (err != nil) != tt.wantErr {
				t.Fatalf("NewConfirmationCode() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if !got.IsValid() {
				t.Errorf("NewConfirmationCode() produced invalid code: %s", got.String())
			}
		})
	}
}

func TestConfirmationCodeFromString(t *testing.T) {
	tests := []struct {
		name    string
		code    string
		want    ConfirmationCode
		wantErr bool
	}{
		{
			name:    "valid code - acceptable string",
			code:    "123457",
			want:    ConfirmationCode("123457"),
			wantErr: false,
		},
		{
			name:    "valid code with zeros",
			code:    "000001",
			want:    ConfirmationCode("000001"),
			wantErr: false,
		},
		{
			name:    "invalid - empty string",
			code:    "",
			want:    ConfirmationCode(""),
			wantErr: true,
		},
		{
			name:    "invalid - too short",
			code:    "12345",
			want:    ConfirmationCode(""),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConfirmationCodeFromString(tt.code)
			if (err != nil) != tt.wantErr {
				t.Fatalf("ConfirmationCodeFromString() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				if err != domerr.ErrInvalidConfirmationCode {
					t.Errorf("ConfirmationCodeFromString() error = %v, want %v", err, domerr.ErrInvalidConfirmationCode)
				}
				return
			}
			if got != tt.want {
				t.Errorf("ConfirmationCodeFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfirmationCode_Equals(t *testing.T) {
	valid := ConfirmationCode("123457")
	otherValid := ConfirmationCode("000001")

	tests := []struct {
		name  string
		c     ConfirmationCode
		other any
		want  bool
	}{
		{"same value", valid, ConfirmationCode("123457"), true},
		{"different value", valid, otherValid, false},
		{"other type - string", valid, "123457", false},
		{"other type - int", valid, 123457, false},
		{"nil", valid, nil, false},
		{"zero value equal to itself", ConfirmationCode(""), ConfirmationCode(""), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Equals(tt.other); got != tt.want {
				t.Errorf("ConfirmationCode.Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfirmationCode_IsValid(t *testing.T) {
	tests := []struct {
		name string
		c    ConfirmationCode
		want bool
	}{
		{"valid code", ConfirmationCode("123457"), true},
		{"invalid - empty", ConfirmationCode(""), false},
		{"invalid - too short", ConfirmationCode("123"), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.IsValid(); got != tt.want {
				t.Errorf("ConfirmationCode.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfirmationCode_IsZero(t *testing.T) {
	tests := []struct {
		name string
		c    ConfirmationCode
		want bool
	}{
		{"zero - empty", ConfirmationCode(""), true},
		{"non-zero - valid code", ConfirmationCode("123457"), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.IsZero(); got != tt.want {
				t.Errorf("ConfirmationCode.IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfirmationCode_String(t *testing.T) {
	tests := []struct {
		name string
		c    ConfirmationCode
		want string
	}{
		{"valid code", ConfirmationCode("123457"), "123457"},
		{"empty", ConfirmationCode(""), ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.String(); got != tt.want {
				t.Errorf("ConfirmationCode.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
