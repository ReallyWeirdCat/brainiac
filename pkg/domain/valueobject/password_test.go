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
	errs "errors"
	"reflect"
	"testing"

	"github.com/ReallyWeirdCat/brainiac/pkg/domain/errors"
)

func TestNewPassword(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name    string
		args    args
		want    Password
		wantErr error
	}{
		{
			name: "valid password with all requirements",
			args: args{
				password: "Abcdef1!",
			},
			want:    Password("Abcdef1!"),
			wantErr: nil,
		},
		{
			name: "valid password exactly 8 chars",
			args: args{
				password: "Aa1!aaaa",
			},
			want:    Password("Aa1!aaaa"),
			wantErr: nil,
		},
		{
			name: "too short - less than 8 chars",
			args: args{
				password: "Abc1fa!",
			},
			want:    Password(""),
			wantErr: errors.ErrWeakPassword,
		},
		{
			name: "too long - more than 72 bytes",
			args: args{
				password: string(make([]byte, 73)),
			},
			want:    Password(""),
			wantErr: errors.ErrPasswordTooLong,
		},
		{
			name: "missing digit",
			args: args{
				password: "Abcdefgh!",
			},
			want:    Password(""),
			wantErr: errors.ErrWeakPassword,
		},
		{
			name: "missing lowercase",
			args: args{
				password: "ABCDEF1!",
			},
			want:    Password(""),
			wantErr: errors.ErrWeakPassword,
		},
		{
			name: "missing uppercase",
			args: args{
				password: "abcdef1!",
			},
			want:    Password(""),
			wantErr: errors.ErrWeakPassword,
		},
		{
			name: "missing symbol",
			args: args{
				password: "Abcdefg1",
			},
			want:    Password(""),
			wantErr: errors.ErrWeakPassword,
		},
		{
			name: "invalid characters (tab)",
			args: args{
				password: "Valid1!\t",
			},
			want:    Password(""),
			wantErr: errors.ErrWeakPassword,
		},
		{
			name: "invalid characters (Cyrillic)",
			args: args{
				password: "Валид1!\t",
			},
			want:    Password(""),
			wantErr: errors.ErrWeakPassword,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPassword(tt.args.password)
			if tt.wantErr != nil {
				if err == nil {
					t.Fatalf("NewPassword() error = nil, wantErr %v", tt.wantErr)
				}
				if !errs.Is(err, tt.wantErr) {
					t.Fatalf("NewPassword() error = %v, want %v", err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Fatalf("NewPassword() unexpected error = %v", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}
