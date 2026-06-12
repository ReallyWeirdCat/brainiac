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

func TestNewLanguageCode(t *testing.T) {
	type args struct {
		languageCode string
	}
	tests := []struct {
		name    string
		args    args
		want    LanguageCode
		wantErr bool
	}{
		// Valid cases - exactly 2 letters
		{
			name:    "valid lowercase english",
			args:    args{languageCode: "en"},
			want:    LanguageCode{value: "en"},
			wantErr: false,
		},
		{
			name:    "valid lowercase spanish",
			args:    args{languageCode: "es"},
			want:    LanguageCode{value: "es"},
			wantErr: false,
		},
		{
			name:    "valid lowercase french",
			args:    args{languageCode: "fr"},
			want:    LanguageCode{value: "fr"},
			wantErr: false,
		},
		{
			name:    "valid lowercase german",
			args:    args{languageCode: "de"},
			want:    LanguageCode{value: "de"},
			wantErr: false,
		},
		{
			name:    "valid lowercase russian",
			args:    args{languageCode: "ru"},
			want:    LanguageCode{value: "ru"},
			wantErr: false,
		},
		{
			name:    "valid lowercase chinese",
			args:    args{languageCode: "zh"},
			want:    LanguageCode{value: "zh"},
			wantErr: false,
		},
		{
			name:    "valid lowercase japanese",
			args:    args{languageCode: "ja"},
			want:    LanguageCode{value: "ja"},
			wantErr: false,
		},
		{
			name:    "valid lowercase arabic",
			args:    args{languageCode: "ar"},
			want:    LanguageCode{value: "ar"},
			wantErr: false,
		},

		// Valid cases with uppercasing (should be sanitized to lowercase)
		{
			name:    "valid uppercase becomes lowercase",
			args:    args{languageCode: "EN"},
			want:    LanguageCode{value: "en"},
			wantErr: false,
		},
		{
			name:    "valid mixed case becomes lowercase",
			args:    args{languageCode: "En"},
			want:    LanguageCode{value: "en"},
			wantErr: false,
		},
		{
			name:    "valid mixed case french",
			args:    args{languageCode: "Fr"},
			want:    LanguageCode{value: "fr"},
			wantErr: false,
		},

		// Invalid cases - wrong length
		{
			name:    "invalid empty string",
			args:    args{languageCode: ""},
			want:    LanguageCode{},
			wantErr: true,
		},
		{
			name:    "invalid too short - 1 letter",
			args:    args{languageCode: "e"},
			want:    LanguageCode{},
			wantErr: true,
		},
		{
			name:    "invalid too long - 3 letters",
			args:    args{languageCode: "eng"},
			want:    LanguageCode{},
			wantErr: true,
		},
		{
			name:    "invalid too long - 4 letters",
			args:    args{languageCode: "engl"},
			want:    LanguageCode{},
			wantErr: true,
		},

		// Invalid cases - wrong characters
		{
			name:    "invalid contains number",
			args:    args{languageCode: "e1"},
			want:    LanguageCode{},
			wantErr: true,
		},
		{
			name:    "invalid contains special char",
			args:    args{languageCode: "e-"},
			want:    LanguageCode{},
			wantErr: true,
		},
		{
			name:    "invalid contains space",
			args:    args{languageCode: "e "},
			want:    LanguageCode{},
			wantErr: true,
		},
		{
			name:    "invalid contains underscore",
			args:    args{languageCode: "e_"},
			want:    LanguageCode{},
			wantErr: true,
		},
		{
			name:    "invalid contains punctuation",
			args:    args{languageCode: "e."},
			want:    LanguageCode{},
			wantErr: true,
		},

		// Invalid cases - non-Latin characters
		{
			name:    "invalid cyrillic letters",
			args:    args{languageCode: "ен"},
			want:    LanguageCode{},
			wantErr: true,
		},
		{
			name:    "invalid chinese characters",
			args:    args{languageCode: "中文"},
			want:    LanguageCode{},
			wantErr: true,
		},
		{
			name:    "invalid arabic characters",
			args:    args{languageCode: "عع"},
			want:    LanguageCode{},
			wantErr: true,
		},

		// Invalid cases - whitespace
		{
			name:    "invalid leading space",
			args:    args{languageCode: " en"},
			want:    LanguageCode{},
			wantErr: true,
		},
		{
			name:    "invalid trailing space",
			args:    args{languageCode: "en "},
			want:    LanguageCode{},
			wantErr: true,
		},
		{
			name:    "invalid tab",
			args:    args{languageCode: "en\t"},
			want:    LanguageCode{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewLanguageCode(tt.args.languageCode)
			if (err != nil) != tt.wantErr {
				t.Fatalf("NewLanguageCode() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLanguageCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
