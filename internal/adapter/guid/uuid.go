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

package guid

import (
	"github.com/ReallyWeirdCat/brainiac/pkg/domain/app/ports"
	"github.com/ReallyWeirdCat/brainiac/pkg/domain/valueobject"
	"github.com/google/uuid"
)

type UuidGuidProvider struct{}

func (u UuidGuidProvider) New() (valueobject.GUID, error) {
	uuid := uuid.New()
	guid, err := valueobject.NewGUID(uuid.String())
	if err != nil {
		return valueobject.GUID(""), err
	}
	return guid, nil
}

var _ ports.GuidProvider = UuidGuidProvider{}
