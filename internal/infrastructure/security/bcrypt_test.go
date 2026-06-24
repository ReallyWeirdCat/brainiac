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

package security

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestBcryptPasswordHasher_Compare(t *testing.T) {
	// Pre‑hash a known password for valid comparison tests
	validPassword := "securePass123"
	hashed, err := bcrypt.GenerateFromPassword([]byte(validPassword), bcrypt.DefaultCost)
	if err != nil {
		t.Fatalf("failed to generate test hash: %v", err)
	}
	validHash := string(hashed)

	type args struct {
		hashedPassword string
		password       string
	}
	tests := []struct {
		name    string
		b       BcryptPasswordHasher
		args    args
		wantErr bool
	}{
		{
			name: "correct password",
			b:    BcryptPasswordHasher{},
			args: args{
				hashedPassword: validHash,
				password:       validPassword,
			},
			wantErr: false,
		},
		{
			name: "wrong password",
			b:    BcryptPasswordHasher{},
			args: args{
				hashedPassword: validHash,
				password:       "wrongPassword",
			},
			wantErr: true,
		},
		{
			name: "invalid hash (malformed)",
			b:    BcryptPasswordHasher{},
			args: args{
				hashedPassword: "not-a-bcrypt-hash",
				password:       validPassword,
			},
			wantErr: true,
		},
		{
			name: "empty password (bcrypt accepts empty)",
			b:    BcryptPasswordHasher{},
			args: args{
				hashedPassword: validHash,
				password:       "",
			},
			wantErr: true, // empty password does not match the hashed one
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := BcryptPasswordHasher{}
			if err := b.Compare(tt.args.hashedPassword, tt.args.password); (err != nil) != tt.wantErr {
				t.Errorf("BcryptPasswordHasher.Compare() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBcryptPasswordHasher_Hash(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name    string
		b       BcryptPasswordHasher
		args    args
		want    string // not used for exact equality; we check non‑empty and compare later
		wantErr bool
	}{
		{
			name:    "hash valid password",
			b:       BcryptPasswordHasher{},
			args:    args{password: "mySecret"},
			want:    "", // we don't assert exact value
			wantErr: false,
		},
		{
			name:    "hash empty password (bcrypt handles it)",
			b:       BcryptPasswordHasher{},
			args:    args{password: ""},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := BcryptPasswordHasher{}
			got, err := b.Hash(tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Fatalf("BcryptPasswordHasher.Hash() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			// Ensure the hash is not empty and is a valid bcrypt hash
			if got == "" {
				t.Error("Hash() returned empty string")
			}
			// Optionally verify that the hash can be compared with the original password
			if err := bcrypt.CompareHashAndPassword([]byte(got), []byte(tt.args.password)); err != nil {
				t.Errorf("Hash() produced invalid hash: %v", err)
			}
			// The 'want' field is ignored; we keep the signature unchanged.
			if tt.want != "" && got != tt.want {
				t.Errorf("BcryptPasswordHasher.Hash() = %v, want %v", got, tt.want)
			}
		})
	}
}
