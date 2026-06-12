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

func TestNewEmail(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		args    args
		want    *Email
		wantErr bool
	}{
		{
			name:    "valid email - simple",
			args:    args{email: "user@example.com"},
			want:    &Email{value: "user@example.com"},
			wantErr: false,
		},
		{
			name:    "valid email - cyrillic characters",
			args:    args{email: "юзер@почта.рф"},
			want:    &Email{value: "юзер@почта.рф"},
			wantErr: false,
		},
		{
			name:    "valid email - with plus",
			args:    args{email: "user+tag@example.com"},
			want:    &Email{value: "user+tag@example.com"},
			wantErr: false,
		},
		{
			name:    "valid email - with dot",
			args:    args{email: "first.last@example.co.uk"},
			want:    &Email{value: "first.last@example.co.uk"},
			wantErr: false,
		},
		{
			name:    "valid email - numbers",
			args:    args{email: "user123@example.com"},
			want:    &Email{value: "user123@example.com"},
			wantErr: false,
		},
		{
			name:    "valid email - untrimmed",
			args:    args{email: "  user@example.com   "},
			want:    &Email{value: "user@example.com"},
			wantErr: false,
		},
		{
			name:    "invalid email - missing @",
			args:    args{email: "userexample.com"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "invalid email - missing domain",
			args:    args{email: "user@"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "invalid email - missing local part",
			args:    args{email: "@example.com"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "invalid email - empty string",
			args:    args{email: ""},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "invalid email - spaces only",
			args:    args{email: "   "},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "invalid email - invalid characters",
			args:    args{email: "user name@example.com"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "invalid email - double dots",
			args:    args{email: "user..name@example.com"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "invalid email - no dot in domain",
			args:    args{email: "user@example"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "invalid email - multiple dots in domain",
			args:    args{email: "user@example..com"},
			want:    nil,
			wantErr: true,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewEmail(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Fatalf("NewEmail() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}
