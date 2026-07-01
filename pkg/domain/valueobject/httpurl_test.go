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
	"net/url"
	"reflect"
	"testing"
)

func TestNewHttpUrl(t *testing.T) {
	type args struct {
		rawUrl string
	}
	tests := []struct {
		name    string
		args    args
		want    HttpUrl
		wantErr bool
	}{
		{
			name:    "valid HTTP URL",
			args:    args{rawUrl: "http://example.com"},
			want:    HttpUrl("http://example.com"),
			wantErr: false,
		},
		{
			name:    "valid HTTPS URL",
			args:    args{rawUrl: "https://example.com/path?query=1"},
			want:    HttpUrl("https://example.com/path?query=1"),
			wantErr: false,
		},
		{
			name:    "URL with whitespace trimmed",
			args:    args{rawUrl: "  https://example.com  "},
			want:    HttpUrl("https://example.com"),
			wantErr: false,
		},
		{
			name:    "Automatic HTTPS for URL with missing schema",
			args:    args{rawUrl: "example.com"},
			want:    HttpUrl("https://example.com"),
			wantErr: false,
		},
		{
			name:    "empty string",
			args:    args{rawUrl: ""},
			wantErr: true,
		},
		{
			name:    "whitespace only",
			args:    args{rawUrl: "   "},
			wantErr: true,
		},
		{
			name:    "invalid scheme - ftp",
			args:    args{rawUrl: "ftp://example.com"},
			wantErr: true,
		},
		{
			name:    "missing host",
			args:    args{rawUrl: "https:///path"},
			wantErr: true,
		},
		{
			name:    "malformed URL",
			args:    args{rawUrl: "http://exam ple.com"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewHttpUrl(tt.args.rawUrl)
			if (err != nil) != tt.wantErr {
				t.Fatalf("NewHttpUrl() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}

			_, err = mustParseURL(tt.want.String())
			if err != nil {
				t.Fatalf("url.Parse() error = %v, expected %v to be parsed", err, tt.want)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHttpUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Helper function to create URLs for test expectations
func mustParseURL(rawURL string) (url.URL, error) {
	parsed, err := url.Parse(rawURL)
	if err != nil {
		return url.URL{}, err
	}
	return *parsed, nil
}
