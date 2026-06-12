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
	"reflect"
	"testing"
)

func TestNewNickname(t *testing.T) {
	type args struct {
		nickname string
	}
	tests := []struct {
		name    string
		args    args
		want    Nickname
		wantErr bool
	}{
		// Valid cases
		{
			name:    "valid simple latin nickname",
			args:    args{nickname: "john_doe123"},
			want:    Nickname{value: "john_doe123"},
			wantErr: false,
		},
		{
			name:    "valid with spaces",
			args:    args{nickname: "John Doe"},
			want:    Nickname{value: "John Doe"},
			wantErr: false,
		},
		{
			name:    "valid with capital letters",
			args:    args{nickname: "JohnDoe"},
			want:    Nickname{value: "JohnDoe"},
			wantErr: false,
		},
		{
			name:    "valid multilingual (Cyrillic)",
			args:    args{nickname: "Иван123"},
			want:    Nickname{value: "Иван123"},
			wantErr: false,
		},
		{
			name:    "valid multilingual (Chinese)",
			args:    args{nickname: "张三_123"},
			want:    Nickname{value: "张三_123"},
			wantErr: false,
		},
		{
			name:    "valid with accents",
			args:    args{nickname: "José_Pérez"},
			want:    Nickname{value: "José_Pérez"},
			wantErr: false,
		},
		{
			name:    "valid exactly 3 characters",
			args:    args{nickname: "abc"},
			want:    Nickname{value: "abc"},
			wantErr: false,
		},
		{
			name:    "valid exactly 30 characters",
			args:    args{nickname: "a23456789012345678901234567890"},
			want:    Nickname{value: "a23456789012345678901234567890"},
			wantErr: false,
		},
		{
			name:    "valid with only underscores",
			args:    args{nickname: "___"},
			want:    Nickname{value: "___"},
			wantErr: false,
		},
		{
			name:    "valid with numbers only",
			args:    args{nickname: "12345678"},
			want:    Nickname{value: "12345678"},
			wantErr: false,
		},
		{
			name:    "valid Arabic script",
			args:    args{nickname: "محمد_123"},
			want:    Nickname{value: "محمد_123"},
			wantErr: false,
		},
		{
			name:    "valid Japanese",
			args:    args{nickname: "田中_太郎"},
			want:    Nickname{value: "田中_太郎"},
			wantErr: false,
		},
		{
			name:    "valid with space trimming",
			args:    args{nickname: " bob   "},
			want:    Nickname{value: "bob"},
			wantErr: false,
		},

		// Invalid cases - too short/long
		{
			name:    "invalid too short (2 chars)",
			args:    args{nickname: "ab"},
			want:    Nickname{value: ""},
			wantErr: true,
		},
		{
			name:    "invalid empty string",
			args:    args{nickname: ""},
			want:    Nickname{value: ""},
			wantErr: true,
		},
		{
			name:    "invalid too long (31 chars)",
			args:    args{nickname: "a234567890123456789012345678901"},
			want:    Nickname{value: ""},
			wantErr: true,
		},

		// Invalid cases - disallowed characters
		{
			name:    "invalid with @ symbol",
			args:    args{nickname: "john@doe"},
			want:    Nickname{value: ""},
			wantErr: true,
		},
		{
			name:    "invalid with # symbol",
			args:    args{nickname: "john#doe"},
			want:    Nickname{value: ""},
			wantErr: true,
		},
		{
			name:    "invalid with hyphen",
			args:    args{nickname: "john-doe"},
			want:    Nickname{value: ""},
			wantErr: true,
		},
		{
			name:    "invalid with exclamation",
			args:    args{nickname: "john!doe"},
			want:    Nickname{value: ""},
			wantErr: true,
		},
		{
			name:    "invalid with newline",
			args:    args{nickname: "john\ndoe"},
			want:    Nickname{value: ""},
			wantErr: true,
		},
		{
			name:    "invalid with tab",
			args:    args{nickname: "john\tdoe"},
			want:    Nickname{value: ""},
			wantErr: true,
		},
		{
			name:    "invalid with emoji",
			args:    args{nickname: "john😀doe"},
			want:    Nickname{value: ""},
			wantErr: true,
		},
		{
			name:    "invalid with special symbol",
			args:    args{nickname: "john$doe"},
			want:    Nickname{value: ""},
			wantErr: true,
		},
		{
			name:    "invalid with percent",
			args:    args{nickname: "john%doe"},
			want:    Nickname{value: ""},
			wantErr: true,
		},
		{
			name:    "invalid with asterisk",
			args:    args{nickname: "john*doe"},
			want:    Nickname{value: ""},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewNickname(tt.args.nickname)
			if (err != nil) != tt.wantErr {
				t.Fatalf("NewNickname() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNickname() = %v, want %v", got, tt.want)
			}
		})
	}
}
