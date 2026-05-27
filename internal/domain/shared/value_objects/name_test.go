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
)

func TestNewName(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantError bool
	}{
		// Valid names
		{"valid latin name", "John Doe", false},
		{"valid cyrillic name", "Иван Петров", false},
		{"valid name with hyphen", "Jean-Pierre", false},
		{"valid single letter", "A", false},
		{"valid name with spaces", "Mary Ann Smith", false},
		{"valid name with diacritics", "José María", false},
		{"valid greek name", "Γιώργος Παπαδόπουλος", false},
		{"valid chinese name", "王小明", false},
		{"valid japanese name", "山田太郎", false},
		{"valid korean name", "김철수", false},
		{"valid name with multiple spaces", "John  Michael  Doe", false},
		{"valid name with trailing space", "John ", false},
		{"valid name with leading space", " John", false},
		{"valid length name", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", false},  // 70 characters
		// Invalid names
		{"empty string", "", true},
		{"too long string", string(make([]byte, 71)), true},
		{"invalid characters numbers", "John123", true},
		{"invalid characters symbols", "John@Doe", true},
		{"invalid characters punctuation", "John,Doe", true},
		{"invalid characters quotes", "John\"Doe", true},
		{"invalid characters apostrophe", "John'Doe", true},
		{"invalid characters asterisk", "John*Doe", true},
		{"invalid characters line break", "John\nDoe", true},
		{"invalid characters tab", "John\tDoe", true},
		{"invalid special characters", "!@#$%", true},
		{"emojis", "Nikita💀", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := NewName(tt.input)
			
			if tt.wantError {
				if err == nil {
					t.Errorf("NewName(%q) expected error but got none", tt.input)
				}
			} else {
				if err != nil {
					t.Errorf("NewName(%q) unexpected error: %v", tt.input, err)
				}
				if result.String() != tt.input {
					t.Errorf("NewName(%q) = %q, want %q", tt.input, result.String(), tt.input)
				}
			}
		})
	}
}

func TestNameString(t *testing.T) {
	name := "Test User"
	n, err := NewName(name)
	if err != nil {
		t.Fatalf("Failed to create name: %v", err)
	}
	
	if n.String() != name {
		t.Errorf("String() = %q, want %q", n.String(), name)
	}
}

func TestNameImmutability(t *testing.T) {
	original := "John Doe"
	n, err := NewName(original)
	if err != nil {
		t.Fatalf("Failed to create name: %v", err)
	}
	
	// Try to modify the returned value (should not affect the struct)
	modified := n.String()
	modified = "Jane Smith"
	
	if n.String() == modified {
		t.Errorf("Name was modified, got %q, want %q", n.String(), original)
	}
}

func TestNameEquality(t *testing.T) {
	name1, _ := NewName("John Doe")
	name2, _ := NewName("John Doe")
	name3, _ := NewName("Jane Doe")
	
	if name1.String() != name2.String() {
		t.Errorf("Names with same value should be equal, got %q and %q", name1.String(), name2.String())
	}
	
	if name1.String() == name3.String() {
		t.Errorf("Names with different values should not be equal, got %q and %q", name1.String(), name3.String())
	}
}
