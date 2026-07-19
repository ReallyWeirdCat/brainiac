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
	"crypto/tls"
	"fmt"
	"net"
	"net/smtp"

	"github.com/ReallyWeirdCat/brainiac/pkg/domain/app/ports"
	cfg "github.com/ReallyWeirdCat/brainiac/pkg/domain/config"
)

const smtpPortSMTPS = 465

type smtpMailer struct {
	config cfg.AppConfig
}

func NewSmtpMailer(provider cfg.AppConfigProvider) ports.Mailer {
	return &smtpMailer{config: provider.Get()}
}

// Send sends a raw email message to the given recipients using the configured SMTP server.
// The from address is used as the envelope sender (MAIL FROM).
// The msg parameter must be a complete MIME message including headers and body.
func (s *smtpMailer) Send(ctx context.Context, from string, to []string, msg []byte) error {
	cfg := s.config.SMTP
	if !cfg.Enable {
		return fmt.Errorf("SMTP is disabled")
	}
	if cfg.Host == "" || cfg.Port == 0 {
		return fmt.Errorf("SMTP host or port not configured")
	}

	// Port 465 needs TLS regardless
	useTLS := cfg.UseTLS || cfg.Port == smtpPortSMTPS

	// Authentication requires TLS.
	if (cfg.Username != "" || cfg.Password != "") && !useTLS {
		return fmt.Errorf("SMTP authentication requires TLS")
	}
	if (cfg.Username != "") != (cfg.Password != "") {
		return fmt.Errorf("both username and password must be provided for authentication")
	}

	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)

	client, cancel, err := dialSMTP(ctx, addr, cfg.Host, useTLS, int(cfg.Port))
	if err != nil {
		return err
	}
	defer cancel()
	defer client.Close()

	if err := ctx.Err(); err != nil {
		return err
	}

	if cfg.Username != "" && cfg.Password != "" {
		auth := smtp.PlainAuth("", cfg.Username, cfg.Password, cfg.Host)
		if err := client.Auth(auth); err != nil {
			return fmt.Errorf("failed to authenticate: %w", err)
		}
		if err := ctx.Err(); err != nil {
			return err
		}
	}

	if err := client.Mail(from); err != nil {
		return fmt.Errorf("failed to set sender: %w", err)
	}
	if err := ctx.Err(); err != nil {
		return err
	}

	for _, recipient := range to {
		if err := client.Rcpt(recipient); err != nil {
			return fmt.Errorf("failed to set recipient %q: %w", recipient, err)
		}
		if err := ctx.Err(); err != nil {
			return err
		}
	}

	w, err := client.Data()
	if err != nil {
		return fmt.Errorf("failed to get data writer: %w", err)
	}
	defer func() {
		if closeErr := w.Close(); closeErr != nil && err == nil {
			err = fmt.Errorf("failed to close data writer: %w", closeErr)
		}
	}()

	if err := ctx.Err(); err != nil {
		return err
	}

	if _, err := w.Write(msg); err != nil {
		return fmt.Errorf("failed to write message: %w", err)
	}

	return nil
}

// dialSMTP establishes a connection to the SMTP server, optionally using TLS,
// and returns a ready client along with a cancel function to stop the internal
// context monitor. The caller must cancel the monitor when done with the client.
func dialSMTP(ctx context.Context, addr, host string, useTLS bool, port int) (*smtp.Client, context.CancelFunc, error) {
	var conn net.Conn
	var err error

	if useTLS && port == smtpPortSMTPS {
		tlsDialer := &tls.Dialer{
			NetDialer: &net.Dialer{},
			Config:    &tls.Config{ServerName: host},
		}
		conn, err = tlsDialer.DialContext(ctx, "tcp", addr)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to dial with TLS: %w", err)
		}
	} else {
		dialer := &net.Dialer{}
		conn, err = dialer.DialContext(ctx, "tcp", addr)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to dial: %w", err)
		}
	}

	if deadline, ok := ctx.Deadline(); ok {
		if err := conn.SetDeadline(deadline); err != nil {
			conn.Close()
			return nil, nil, fmt.Errorf("failed to set deadline: %w", err)
		}
	}

	monitorCtx, monitorCancel := context.WithCancel(context.Background())
	go func() {
		select {
		case <-ctx.Done():
			conn.Close()
		case <-monitorCtx.Done():
		}
	}()

	client, err := smtp.NewClient(conn, host)
	if err != nil {
		conn.Close()
		monitorCancel()
		return nil, nil, fmt.Errorf("failed to create SMTP client: %w", err)
	}

	if useTLS && port != smtpPortSMTPS {
		if err := client.StartTLS(&tls.Config{ServerName: host}); err != nil {
			client.Close()
			monitorCancel()
			return nil, nil, fmt.Errorf("failed to start TLS: %w", err)
		}
	}

	return client, monitorCancel, nil
}

var _ ports.Mailer = &smtpMailer{}
