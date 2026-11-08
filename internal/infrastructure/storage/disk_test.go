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

package storage

import (
	"context"
	"io"
	"strings"
	"testing"
)

func TestDiskStorePutGetDelete(t *testing.T) {
	store := &DiskStore{root: t.TempDir()}
	ctx := context.Background()

	key, err := store.Put(ctx, "dir/file.txt", strings.NewReader("hello"), 5, "text/plain")
	if err != nil {
		t.Fatalf("Put: %v", err)
	}
	if key != "dir/file.txt" {
		t.Errorf("key = %q, want %q", key, "dir/file.txt")
	}

	rc, err := store.Get(ctx, "dir/file.txt")
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	defer rc.Close()
	data, err := io.ReadAll(rc)
	if err != nil {
		t.Fatalf("ReadAll: %v", err)
	}
	if string(data) != "hello" {
		t.Errorf("data = %q, want %q", data, "hello")
	}

	if err := store.Delete(ctx, "dir/file.txt"); err != nil {
		t.Fatalf("Delete: %v", err)
	}
	if err := store.Delete(ctx, "dir/file.txt"); err != nil {
		t.Fatalf("second Delete should be a no-op: %v", err)
	}
}

func TestSafeKey(t *testing.T) {
	valid := map[string]string{
		"a/b.txt":     "a/b.txt",
		"dir//file":   "dir/file",
		"back\\slash": "back/slash",
		"./a.txt":     "a.txt",
	}
	for in, want := range valid {
		got, err := safeKey(in)
		if err != nil {
			t.Errorf("safeKey(%q) unexpected error: %v", in, err)
			continue
		}
		if got != want {
			t.Errorf("safeKey(%q) = %q, want %q", in, got, want)
		}
	}

	for _, in := range []string{"", "..", "../etc/passwd", "a/../../b", "/etc/passwd", "/a/b.txt", "a\x00b"} {
		if _, err := safeKey(in); err == nil {
			t.Errorf("safeKey(%q) expected error, got nil", in)
		}
	}
}
