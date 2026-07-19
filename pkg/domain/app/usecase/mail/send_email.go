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

package mail

import (
	"context"
	"fmt"
	"strings"

	"github.com/ReallyWeirdCat/brainiac/pkg/domain/app/ports"
	"github.com/ReallyWeirdCat/brainiac/pkg/domain/config"
)

type SendEmailUsecase struct {
	config config.AppConfig
	mailer ports.Mailer
}

func NewSendEmailUsecase(config config.AppConfigProvider, mailer ports.Mailer) *SendEmailUsecase {
	return &SendEmailUsecase{config: config.Get(), mailer: mailer}
}

func (s *SendEmailUsecase) Execute(ctx context.Context, req SendEmailRequest) (*SendEmailResponse, error) {
	// Validate request
	if len(req.To) == 0 {
		return nil, ErrNoRecipients
	}
	if req.Subject == "" {
		return nil, ErrEmptySubject
	}
	if req.Body == "" {
		return nil, ErrEmptyBody
	}

	from := req.From
	if from == "" {
		from = s.config.SMTP.From
	}
	if from == "" {
		return nil, ErrSenderNotSet
	}

	var headers strings.Builder
	headers.WriteString(fmt.Sprintf("From: %s\r\n", from))
	headers.WriteString(fmt.Sprintf("To: %s\r\n", strings.Join(req.To, ", ")))
	headers.WriteString(fmt.Sprintf("Subject: %s\r\n", req.Subject))
	headers.WriteString("MIME-Version: 1.0\r\n")
	if req.IsHTML {
		headers.WriteString("Content-Type: text/html; charset=utf-8\r\n")
	} else {
		headers.WriteString("Content-Type: text/plain; charset=utf-8\r\n")
	}
	headers.WriteString("\r\n")

	msg := []byte(headers.String() + req.Body)

	// Collect all recipients (To + CC + BCC)
	allRecipients := make([]string, 0, len(req.To))
	allRecipients = append(allRecipients, req.To...)

	err := s.mailer.Send(ctx, from, allRecipients, msg)
	if err != nil {
		return nil, fmt.Errorf("failed to send email: %w", err)
	}

	return &SendEmailResponse{}, nil
}
