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

	"github.com/ReallyWeirdCat/brainiac/pkg/domain/valueobject"
)

type RegistrationInvite struct {
	AppUserGUID          valueobject.GUID      `json:"app_user_guid"`
	InvitedByAppUserGUID *valueobject.GUID     `json:"invited_by_app_user_guid,omitempty"`
	InviteCode           valueobject.GUID      `json:"invite_code"`
	Email                *valueobject.Email    `json:"email,omitempty"`
	Message              *string               `json:"message,omitempty"`
	Name                 *valueobject.Name     `json:"name,omitempty"`
	Surname              *valueobject.Name     `json:"surname,omitempty"`
	Patronymic           *valueobject.Name     `json:"patronymic,omitempty"`
	Nickname             *valueobject.Nickname `json:"nickname,omitempty"`
	Username             *valueobject.Username `json:"username,omitempty"`
	StudentGroups        *valueobject.Metadata `json:"student_groups,omitempty"`
	TeacherGroups        *valueobject.Metadata `json:"teacher_groups,omitempty"`
	Chats                *valueobject.Metadata `json:"chats,omitempty"`
	ExpireAt             *time.Time            `json:"expire_at,omitempty"`
	UsedAt               *time.Time            `json:"used_at,omitempty"`
	CreatedAt            time.Time             `json:"created_at"`
	DeletedAt            *time.Time            `json:"deleted_at,omitempty"`
}

var _ Entity = Language{}

func (r RegistrationInvite) IsValid() bool {
	if r.AppUserGUID == nil || !r.AppUserGUID.IsValid() {
		return false
	}
	if r.InviteCode == nil || !r.InviteCode.IsValid() {
		return false
	}
	if r.Email != nil && !r.Email.IsValid() {
		return false
	}
	if r.InvitedByAppUserGUID != nil && !(*r.InvitedByAppUserGUID).IsValid() {
		return false
	}
	if r.Name != nil && !r.Name.IsValid() {
		return false
	}
	if r.Surname != nil && !r.Surname.IsValid() {
		return false
	}
	if r.Nickname != nil && !r.Nickname.IsValid() {
		return false
	}
	if r.Username != nil && !r.Username.IsValid() {
		return false
	}
	if r.StudentGroups != nil && !(*r.StudentGroups).IsValid() {
		return false
	}
	if r.TeacherGroups != nil && !(*r.TeacherGroups).IsValid() {
		return false
	}
	if r.Chats != nil && !(*r.Chats).IsValid() {
		return false
	}

	return true
}
