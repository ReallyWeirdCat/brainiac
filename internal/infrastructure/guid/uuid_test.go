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

package guid_test

import (
	"testing"

	"github.com/ReallyWeirdCat/brainiac/internal/infrastructure/guid"
)

func TestUuidGuidProvider_New(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{name: "Generated UUID", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := guid.UuidGuidProvider{}
			_, gotErr := u.New()
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("New() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("New() succeeded unexpectedly")
			}
		})
	}
}
