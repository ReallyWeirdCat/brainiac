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

func TestNewUsername(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantError bool
	}{
		// Valid usernames
		{"valid lowercase", "john_doe", false},
		{"valid uppercase", "JOHN_DOE", false},
		{"valid mixed case", "John_Doe", false},
		{"valid with numbers", "user123", false},
		{"valid only numbers", "123456", false},
		{"valid only underscores", "_____", false},
		{"valid min length", "abc", false},
		{"valid max length", "abcdefghijklmnopqr", false}, // 18 characters
		{"valid alphanumeric", "abc123def456", false},
		{"valid complex", "A1_B2_C3", false},

		// Invalid usernames
		{"empty string", "", true},
		{"too short", "ab", true},
		{"too long", "abcdefghijklmnopqrs", true}, // 19 characters
		{"contains space", "john doe", true},
		{"contains hyphen", "john-doe", true},
		{"contains dot", "john.doe", true},
		{"contains @", "john@doe", true},
		{"contains #", "john#doe", true},
		{"contains $", "john$doe", true},
		{"contains %", "john%doe", true},
		{"contains ^", "john^doe", true},
		{"contains &", "john&doe", true},
		{"contains *", "john*doe", true},
		{"contains (", "john(doe", true},
		{"contains )", "john)doe", true},
		{"contains +", "john+doe", true},
		{"contains =", "john=doe", true},
		{"contains [", "john[doe", true},
		{"contains ]", "john]doe", true},
		{"contains {", "john{doe", true},
		{"contains }", "john}doe", true},
		{"contains |", "john|doe", true},
		{"contains \\", "john\\doe", true},
		{"contains ;", "john;doe", true},
		{"contains :", "john:doe", true},
		{"contains '", "john'doe", true},
		{"contains \"", "john\"doe", true},
		{"contains <", "john<doe", true},
		{"contains >", "john>doe", true},
		{"contains ,", "john,doe", true},
		{"contains ?", "john?doe", true},
		{"contains /", "john/doe", true},
		{"contains newline", "john\ndoe", true},
		{"contains tab", "john\tdoe", true},
		{"contains unicode", "jöhndoe", true},
		{"contains cyrillic", "джон", true},
		{"starts with underscore", "_john", false}, // underscore allowed anywhere
		{"ends with underscore", "john_", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := NewUsername(tt.input)

			if tt.wantError {
				if err == nil {
					t.Errorf("NewUsername(%q) expected error but got none", tt.input)
				}
			} else {
				if err != nil {
					t.Errorf("NewUsername(%q) unexpected error: %v", tt.input, err)
				}
				if result.String() != tt.input {
					t.Errorf("NewUsername(%q) = %q, want %q", tt.input, result.String(), tt.input)
				}
			}
		})
	}
}

func TestUsernameString(t *testing.T) {
	username := "test_user_123"
	u, err := NewUsername(username)
	if err != nil {
		t.Fatalf("Failed to create username: %v", err)
	}

	if u.String() != username {
		t.Errorf("String() = %q, want %q", u.String(), username)
	}
}

func TestUsernameImmutability(t *testing.T) {
	original := "immutable_user"
	u, err := NewUsername(original)
	if err != nil {
		t.Fatalf("Failed to create username: %v", err)
	}

	// Try to modify the returned value
	modified := u.String()
	modified = "hacked_user"

	if u.String() == modified {
		t.Errorf("Username was modified, got %q, want %q", u.String(), original)
	}
}

func TestUsernameEquality(t *testing.T) {
	user1, _ := NewUsername("john_doe")
	user2, _ := NewUsername("john_doe")
	user3, _ := NewUsername("jane_doe")

	if user1.String() != user2.String() {
		t.Errorf("Usernames with same value should be equal, got %q and %q", user1.String(), user2.String())
	}

	if user1.String() == user3.String() {
		t.Errorf("Usernames with different values should not be equal, got %q and %q", user1.String(), user3.String())
	}
}

func TestUsernameCaseSensitivity(t *testing.T) {
	lowerCase, _ := NewUsername("john_doe")
	upperCase, _ := NewUsername("JOHN_DOE")
	mixedCase, _ := NewUsername("John_Doe")

	// Usernames should be case-sensitive
	if lowerCase.String() == upperCase.String() {
		t.Error("Usernames should be case-sensitive, but lowercase and uppercase are considered equal")
	}

	if lowerCase.String() == mixedCase.String() {
		t.Error("Usernames should be case-sensitive, but lowercase and mixed-case are considered equal")
	}

	if upperCase.String() == mixedCase.String() {
		t.Error("Usernames should be case-sensitive, but uppercase and mixed-case are considered equal")
	}
}

func TestUsernameBoundaryConditions(t *testing.T) {
	// Test minimum length (3 characters)
	minUsername := "abc"
	_, err := NewUsername(minUsername)
	if err != nil {
		t.Errorf("Should accept 3-character username, got error: %v", err)
	}

	// Test below minimum length (2 characters)
	_, err = NewUsername("ab")
	if err == nil {
		t.Error("Should reject 2-character username")
	}

	// Test maximum length (18 characters)
	maxUsername := "abcdefghijklmnopqr" // 18 characters
	_, err = NewUsername(maxUsername)
	if err != nil {
		t.Errorf("Should accept 18-character username, got error: %v", err)
	}

	// Test above maximum length (19 characters)
	tooLong := maxUsername + "s" // 19 characters
	_, err = NewUsername(tooLong)
	if err == nil {
		t.Error("Should reject 19-character username")
	}
}

func TestUsernameErrorFormat(t *testing.T) {
	invalidUsername := "invalid username!"
	_, err := NewUsername(invalidUsername)

	if err == nil {
		t.Fatal("Expected error for invalid username")
	}

	expectedErrorSubstring := "invalid username format"
	if !contains(err.Error(), expectedErrorSubstring) {
		t.Errorf("Error message does not contain expected text.\nGot: %q\nWant to contain: %q", err.Error(), expectedErrorSubstring)
	}
}

// Helper function
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(substr) == 0 ||
		(len(s) > 0 && len(substr) > 0 && findSubstring(s, substr)))
}

func findSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
