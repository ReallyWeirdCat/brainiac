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
	"bytes"
	"encoding/json"
	"errors"
)

type Metadata json.RawMessage

var _ ValueObject = Metadata([]byte("{}"))

func NewMetadata(data []byte) (Metadata, error) {
	if !json.Valid(data) {
		return Metadata{}, errors.New("invalid JSON")
	}
	return Metadata(data), nil
}

func NewMetadataFromMap(data map[string]any) (Metadata, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return Metadata{}, err
	}
	return Metadata(b), nil
}

func (m Metadata) Equals(other any) bool {
	otherM, ok := other.(Metadata)
	if !ok {
		return false
	}

	canon1, err1 := canonicalize(m)
	canon2, err2 := canonicalize(otherM)
	if err1 != nil || err2 != nil {
		return bytes.Equal(m, otherM)
	}
	return bytes.Equal(canon1, canon2)
}

func (m Metadata) IsValid() bool {
	return json.Valid(m)
}

func (m Metadata) IsZero() bool {
	return len(m) == 0
}

func (m Metadata) String() string {
	return string(m)
}

func canonicalize(m Metadata) ([]byte, error) {
	if len(m) == 0 {
		return nil, errors.New("nil or empty metadata")
	}
	var obj any
	dec := json.NewDecoder(bytes.NewReader(m))
	dec.UseNumber()
	if err := dec.Decode(&obj); err != nil {
		return nil, err
	}
	return json.Marshal(obj)
}
