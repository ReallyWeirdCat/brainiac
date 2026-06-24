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
	"os"
	"sync"
	"testing"
)

func TestPwnedPasswordChecker_load(t *testing.T) {
	// Create a temporary file with a few passwords and empty lines.
	tmpFile, err := os.CreateTemp("", "rockyou_test_*.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	content := "password123\nqwertyuiop\nadmin123\n\n   \n"
	if _, err := tmpFile.WriteString(content); err != nil {
		t.Fatal(err)
	}
	tmpFile.Close()
	validPath := tmpFile.Name()

	nonExistentPath := "/non/existent/file"

	type fields struct {
		compromised map[string]struct{}
		filePath    string
		loadErr     error
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "successful load",
			fields: fields{
				filePath: validPath,
			},
		},
		{
			name: "file not found",
			fields: fields{
				filePath: nonExistentPath,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new PwnedPasswordChecker with a fresh sync.Once (do not copy from tt.fields)
			p := &PwnedPasswordChecker{
				compromised:       tt.fields.compromised,
				minPasswordLength: 8,
				once:              sync.Once{}, // fresh, no copy
				filePath:          tt.fields.filePath,
				loadErr:           tt.fields.loadErr,
			}
			p.load()

			switch tt.name {
			case "successful load":
				if p.loadErr != nil {
					t.Errorf("load() error = %v, want nil", p.loadErr)
				}
				expected := map[string]struct{}{
					"password123": {},
					"qwertyuiop":  {},
					// admin123 is skipped: length < 8
				}
				if len(p.compromised) != len(expected) {
					t.Errorf("compromised map length = %d, want %d", len(p.compromised), len(expected))
				}
				for k := range expected {
					if _, ok := p.compromised[k]; !ok {
						t.Errorf("missing expected key %q in compromised map", k)
					}
				}
				if _, ok := p.compromised[""]; ok {
					t.Error("empty string should not be in compromised map")
				}
				if _, ok := p.compromised["   "]; ok {
					t.Error("whitespace-only string should not be in compromised map")
				}
			case "file not found":
				if p.loadErr == nil {
					t.Error("load() error = nil, want error")
				}
			}
		})
	}
}

func TestPwnedPasswordChecker_IsCompromised(t *testing.T) {
	// Create a temporary file with test passwords.
	tmpFile, err := os.CreateTemp("", "rockyou_test_*.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	content := "password123\nqwertyuiop\nadmin123\n"
	if _, err := tmpFile.WriteString(content); err != nil {
		t.Fatal(err)
	}
	tmpFile.Close()
	validPath := tmpFile.Name()

	nonExistentPath := "/non/existent/file"

	type fields struct {
		compromised map[string]struct{}
		filePath    string
		loadErr     error
	}
	type args struct {
		password string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "compromised password found",
			fields: fields{
				filePath: validPath,
			},
			args:    args{password: "password123"},
			want:    true,
			wantErr: false,
		},
		{
			name: "unknown password not compromised",
			fields: fields{
				filePath: validPath,
			},
			args:    args{password: "unknown"},
			want:    false,
			wantErr: false,
		},
		{
			name: "file not found error",
			fields: fields{
				filePath: nonExistentPath,
			},
			args:    args{password: "anything"},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PwnedPasswordChecker{
				compromised: tt.fields.compromised,
				once:        sync.Once{},
				filePath:    tt.fields.filePath,
				loadErr:     tt.fields.loadErr,
			}
			got, err := p.IsCompromised(tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Fatalf("PwnedPasswordChecker.IsCompromised() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if got != tt.want {
				t.Errorf("PwnedPasswordChecker.IsCompromised() = %v, want %v", got, tt.want)
			}
		})
	}
}
