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
	"strings"
	"testing"
)

func TestNewBio(t *testing.T) {
	type args struct {
		bio string
	}
	tests := []struct {
		name    string
		args    args
		want    Bio
		wantErr bool
	}{
		{
			name:    "valid bio",
			args:    args{bio: "Hello, I am a software engineer."},
			want:    Bio("Hello, I am a software engineer."),
			wantErr: false,
		},
		{
			name:    "valid bio with minimum length",
			args:    args{bio: "A"},
			want:    Bio("A"),
			wantErr: false,
		},
		{
			name:    "valid bio with maximum length",
			args:    args{bio: strings.Repeat("A", 175)},
			want:    Bio(strings.Repeat("A", 175)),
			wantErr: false,
		},
		{
			name:    "valid bio with special characters",
			args:    args{bio: "Hello! How are you? #gopher"},
			want:    Bio("Hello! How are you? #gopher"),
			wantErr: false,
		},
		{
			name:    "valid bio with newlines",
			args:    args{bio: "Line 1\nLine 2\nLine 3"},
			want:    Bio("Line 1\nLine 2\nLine 3"),
			wantErr: false,
		},
		{
			name:    "valid bio with space trimming",
			args:    args{bio: "  Spaces are not real, they can't hurt you   "},
			want:    Bio("Spaces are not real, they can't hurt you"),
			wantErr: false,
		},
		{
			name:    "empty bio",
			args:    args{bio: ""},
			want:    Bio(""),
			wantErr: true,
		},
		{
			name:    "bio too long",
			args:    args{bio: strings.Repeat("A", 176)},
			want:    Bio(""),
			wantErr: true,
		},
		{
			name:    "contains @ symbol",
			args:    args{bio: "Contact me at user@example.com"},
			want:    Bio(""),
			wantErr: true,
		},
		{
			name:    "contains http URL",
			args:    args{bio: "Check out http://example.com"},
			want:    Bio(""),
			wantErr: true,
		},
		{
			name:    "contains https URL",
			args:    args{bio: "Visit https://golang.org for docs"},
			want:    Bio(""),
			wantErr: true,
		},
		{
			name:    "contains ftp URL",
			args:    args{bio: "Download from ftp://files.example.com"},
			want:    Bio(""),
			wantErr: true,
		},
		{
			name:    "contains file URL",
			args:    args{bio: "Access file:///home/user/doc.pdf"},
			want:    Bio(""),
			wantErr: true,
		},
		{
			name:    "contains ws URL",
			args:    args{bio: "Connect to ws://socket.example.com"},
			want:    Bio(""),
			wantErr: true,
		},
		{
			name:    "contains wss URL",
			args:    args{bio: "Secure connection wss://example.com/socket"},
			want:    Bio(""),
			wantErr: true,
		},
		{
			name:    "contains sftp URL",
			args:    args{bio: "Transfer files via sftp://server.com"},
			want:    Bio(""),
			wantErr: true,
		},
		{
			name:    "contains URL without www",
			args:    args{bio: "Go to https://go.dev"},
			want:    Bio(""),
			wantErr: true,
		},
		{
			name:    "contains text similar to URL but not URL",
			args:    args{bio: "http is a protocol and https is secure"},
			want:    Bio("http is a protocol and https is secure"),
			wantErr: false,
		},
		{
			name:    "only @ symbol",
			args:    args{bio: "@"},
			want:    Bio(""),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewBio(tt.args.bio)
			if (err != nil) != tt.wantErr {
				t.Fatalf("NewBio() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBio() = %v, want %v", got, tt.want)
			}
		})
	}
}
