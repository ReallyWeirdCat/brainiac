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

package entity

import (
	"time"

	"github.com/ReallyWeirdCat/brainiac/pkg/domain/enum"
	"github.com/ReallyWeirdCat/brainiac/pkg/domain/valueobject"
)

type StudentGroupMember struct {
	GUID                 valueobject.GUID      `json:"guid"`
	AppUserGUID          valueobject.GUID      `json:"app_user_guid"`
	StudentRole          enum.StudentRoleEnum  `json:"student_role"`
	Notes                *string               `json:"notes,omitempty"`
	Team                 *string               `json:"team,omitempty"`
	TeamResponsibilities *valueobject.Metadata `json:"team_responsibilities,omitempty"`
	CreatedAt            time.Time             `json:"created_at"`
	DeletedAt            *time.Time            `json:"deleted_at,omitempty"`
}

var _ Entity = &StudentGroupMember{}

func (s *StudentGroupMember) IsValid() bool {

	if s.Team != nil && len(*s.Team) > 18 {
		return false
	}
	return s.GUID.IsValid() && s.AppUserGUID.IsValid() && s.StudentRole.IsValid()
}
