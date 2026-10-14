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

// Package storage contains object-store adapters. DiskStore writes to a local
// directory behind the ports.ObjectStore contract so an S3/Garage backend can
// be dropped in later.
package storage

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/ReallyWeirdCat/brainiac/pkg/domain/app/ports"
	"github.com/ReallyWeirdCat/brainiac/pkg/domain/config"
)

// DiskStore stores files under a local root directory.
type DiskStore struct {
	root    string
	baseURL string
}

var _ ports.ObjectStore = (*DiskStore)(nil)

func NewDiskStore(cfg config.AppConfigProvider) ports.ObjectStore {
	c := cfg.Get().Storage
	return &DiskStore{root: c.Root, baseURL: c.BaseURL}
}

// Put writes the reader to key atomically (temp file + fsync + rename).
func (d *DiskStore) Put(ctx context.Context, key string, r io.Reader, _ int64, _ string) (string, error) {
	if err := ctx.Err(); err != nil {
		return "", err
	}
	clean, err := safeKey(key)
	if err != nil {
		return "", err
	}

	dst := filepath.Join(d.root, filepath.FromSlash(clean))
	dir := filepath.Dir(dst)
	if err := os.MkdirAll(dir, 0o750); err != nil {
		return "", err
	}

	// Create the temp file in the destination directory so the final rename is
	// atomic and stays on the same filesystem.
	tmp, err := os.CreateTemp(dir, ".brainiac-put-*")
	if err != nil {
		return "", err
	}
	tmpName := tmp.Name()
	defer os.Remove(tmpName)

	if _, err := io.Copy(tmp, r); err != nil {
		tmp.Close()
		return "", err
	}
	if err := tmp.Sync(); err != nil {
		tmp.Close()
		return "", err
	}
	if err := tmp.Close(); err != nil {
		return "", err
	}
	if err := os.Rename(tmpName, dst); err != nil {
		return "", err
	}
	return clean, nil
}

func (d *DiskStore) Get(ctx context.Context, key string) (io.ReadCloser, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	clean, err := safeKey(key)
	if err != nil {
		return nil, err
	}
	return os.Open(filepath.Join(d.root, filepath.FromSlash(clean)))
}

func (d *DiskStore) Delete(ctx context.Context, key string) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	clean, err := safeKey(key)
	if err != nil {
		return err
	}
	err = os.Remove(filepath.Join(d.root, filepath.FromSlash(clean)))
	if os.IsNotExist(err) {
		return nil
	}
	return err
}

// safeKey rejects empty, NUL, absolute, and path-traversal keys, then returns a
// normalized relative key.
func safeKey(key string) (string, error) {
	if key == "" || strings.ContainsRune(key, 0) {
		return "", errors.New("invalid storage key")
	}
	if path.IsAbs(key) || strings.Contains(key, "..") {
		return "", fmt.Errorf("invalid storage key %q", key)
	}
	clean := path.Clean(strings.ReplaceAll(key, "\\", "/"))
	if clean == "." || clean == "/" {
		return "", fmt.Errorf("invalid storage key %q", key)
	}
	return clean, nil
}
