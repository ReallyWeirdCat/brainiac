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

package auth

import (
	"context"
	"errors"
	"math/rand"
	"time"

	"github.com/ReallyWeirdCat/brainiac/pkg/domain/app/ports"
	"github.com/ReallyWeirdCat/brainiac/pkg/domain/app/usecase/mail"
	"github.com/ReallyWeirdCat/brainiac/pkg/domain/config"
	"github.com/ReallyWeirdCat/brainiac/pkg/domain/entity"
	"github.com/ReallyWeirdCat/brainiac/pkg/domain/repository"
	"github.com/ReallyWeirdCat/brainiac/pkg/domain/valueobject"
)

type EmailSender interface {
	Execute(ctx context.Context, req mail.SendEmailRequest) (*mail.SendEmailResponse, error)
}

type RegistrationBeginUsecase struct {
	uow              ports.UnitOfWorkProvider
	regCodeCache     ports.Cache[entity.RegistrationCode]
	hasher           ports.PasswordHasher
	pwdChecker       ports.CompromisedPasswordChecker
	logger           ports.Logger
	sendEmailUsecase EmailSender
	config           config.AppConfig
}

func NewRegistrationBeginUsecase(
	uow ports.UnitOfWorkProvider,
	regCodeCache ports.Cache[entity.RegistrationCode],
	hasher ports.PasswordHasher,
	pwdChecker ports.CompromisedPasswordChecker,
	logger ports.Logger,
	sendEmailUsecase *mail.SendEmailUsecase,
	config config.AppConfigProvider,
) *RegistrationBeginUsecase {
	return &RegistrationBeginUsecase{
		uow:              uow,
		regCodeCache:     regCodeCache,
		hasher:           hasher,
		pwdChecker:       pwdChecker,
		logger:           logger,
		sendEmailUsecase: sendEmailUsecase,
		config:           config.Get(),
	}
}

func (r *RegistrationBeginUsecase) Execute(ctx context.Context, req RegistrationBeginRequest) (*RegistrationBeginResponse, error) {
	if !r.config.Registration.Enable {
		return nil, ErrRegistrationDisabled
	}

	randDelay(ctx)

	email, err := valueobject.NewEmail(req.Email)
	if err != nil {
		return nil, err
	}
	username, err := valueobject.NewUsername(req.Username)
	if err != nil {
		return nil, err
	}
	password, err := valueobject.NewPassword(req.Password)
	if err != nil {
		return nil, err
	}

	passwordCompromised, err := r.pwdChecker.IsCompromised(string(password))
	if err != nil {
		return nil, err
	}
	if passwordCompromised {
		return nil, valueobject.ErrCompromisedPassword
	}

	uow, err := r.uow.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer uow.Rollback(ctx)

	usernameTaken, err := uow.AppUsers().ExistsByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	if usernameTaken {
		return nil, ErrUsernameTaken
	}

	existingCredential, err := uow.AppUserCredentials().GetByEmail(ctx, email)
	if err != nil && !errors.Is(err, repository.ErrEntityNotFound) {
		return nil, err
	}

	exists, err := r.regCodeCache.Exists(ctx, username.String())
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrCodeAlreadySent
	}

	regcode, err := entity.NewRegistrationCode(username)
	if err != nil {
		return nil, err
	}

	_, err = r.regCodeCache.SetNX(ctx, username.String(), regcode, regcode.ExpireAt.Sub(time.Now()))
	if err != nil {
		return nil, err
	}

	if existingCredential == nil {
		_, err = r.sendEmailUsecase.Execute(ctx, mail.SendEmailRequest{
			To:      []string{email.String()},
			Subject: "Brainiac Registration",
			Body:    "Your registration code is " + regcode.Code.String(),
			IsHTML:  false,
		})
		if err != nil {
			errDelCache := r.regCodeCache.Delete(ctx, username.String())
			if errDelCache != nil {
				return nil, errors.Join(err, errDelCache)
			}
			return nil, err
		}
		r.logger.Info("requested registration", "username", username.String())
		return &RegistrationBeginResponse{
			Username: username.String(),
		}, nil
	}

	// Proceed as if email is not yet registered to prevent email enumeration attacks
	r.logger.Warn("attempted registration via registered credentials", "user", existingCredential.AppUserGUID, "username", username.String())
	return &RegistrationBeginResponse{
		Username: username.String(),
	}, nil
}

func randDelay(ctx context.Context) {
	select {
	case <-time.After(time.Duration(rand.Intn(3)+1) * time.Second):
	case <-ctx.Done():
	}
}
