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

package repository

import (
	"context"

	"github.com/ReallyWeirdCat/brainiac/pkg/domain/entity"
	"github.com/ReallyWeirdCat/brainiac/pkg/domain/valueobject"
)

type MessageRepository interface {
	Save(ctx context.Context, message entity.Message) error
	Delete(ctx context.Context, guid valueobject.GUID) error
	GetByGUID(ctx context.Context, guid valueobject.GUID) (*entity.Message, error)
	GetByUsername(ctx context.Context, username string) ([]*entity.Message, error)
	GetByUsernameInChat(ctx context.Context, username string, chatGuid valueobject.GUID) ([]*entity.Message, error)
	GetByUsernameInChatWithOffset(ctx context.Context, username string, chatGuid valueobject.GUID, count int, offset int) ([]*entity.Message, error)
	GetByChatGUID(ctx context.Context, guid string) ([]*entity.Message, error)
	GetByChatGUIDWithOffset(ctx context.Context, guid string, count int, offset int) ([]*entity.Message, error)
	ExistsByGUID(ctx context.Context, guid valueobject.GUID) (bool, error)
}
