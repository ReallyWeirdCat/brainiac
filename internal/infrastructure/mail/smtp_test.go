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
	"testing"
	"time"

	smtpmock "github.com/mocktools/go-smtp-mock/v2"
	"github.com/stretchr/testify/require"

	cfg "github.com/ReallyWeirdCat/brainiac/pkg/domain/config"
)

type testConfigProvider struct {
	config cfg.AppConfig
}

func (p *testConfigProvider) Get() cfg.AppConfig {
	return p.config
}

func startMockServer(t *testing.T, opts ...func(*smtpmock.ConfigurationAttr)) *smtpmock.Server {
	t.Helper()

	config := smtpmock.ConfigurationAttr{
		LogToStdout:       false,
		LogServerActivity: false,
	}
	for _, opt := range opts {
		opt(&config)
	}

	server := smtpmock.New(config)
	err := server.Start()
	require.NoError(t, err)

	t.Cleanup(func() {
		_ = server.Stop()
	})

	return server
}

// withBlacklistedRcptTo configures the mock to reject specific recipient emails.
func withBlacklistedRcptTo(emails ...string) func(*smtpmock.ConfigurationAttr) {
	return func(c *smtpmock.ConfigurationAttr) {
		c.BlacklistedRcpttoEmails = emails
	}
}

// withNotRegisteredEmails configures the mock to treat certain emails as non-existent.
func withNotRegisteredEmails(emails ...string) func(*smtpmock.ConfigurationAttr) {
	return func(c *smtpmock.ConfigurationAttr) {
		c.NotRegisteredEmails = emails
	}
}

// withResponseDelay adds a delay to the DATA command to test timeouts.
func withResponseDelay(delaySeconds int) func(*smtpmock.ConfigurationAttr) {
	return func(c *smtpmock.ConfigurationAttr) {
		c.ResponseDelayData = delaySeconds
	}
}

// extractAddressFromMailFrom extracts the email address from the full MAIL FROM command.
// Example: "MAIL FROM:<sender@example.com>" -> "sender@example.com"
func extractAddressFromMailFrom(cmd string) string {
	cmd = strings.TrimPrefix(cmd, "MAIL FROM:")
	cmd = strings.TrimSpace(cmd)
	cmd = strings.Trim(cmd, "<>")
	return cmd
}

// extractAddressFromRcptTo extracts the email address from the full RCPT TO command.
// Example: "RCPT TO:<recipient@example.com>" -> "recipient@example.com"
func extractAddressFromRcptTo(cmd string) string {
	cmd = strings.TrimPrefix(cmd, "RCPT TO:")
	cmd = strings.TrimSpace(cmd)
	cmd = strings.Trim(cmd, "<>")
	return cmd
}

func TestSmtpMailer_Send(t *testing.T) {
	const (
		from    = "sender@example.com"
		to      = "recipient@example.com"
		subject = "Test Subject"
		body    = "Hello, this is a test email."
	)

	msg := []byte(fmt.Sprintf(
		"From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s",
		from, to, subject, body,
	))

	t.Run("successful send (no TLS)", func(t *testing.T) {
		server := startMockServer(t)

		cfg := cfg.AppConfig{}
		cfg.SMTP.Enable = true
		cfg.SMTP.Host = "127.0.0.1"
		cfg.SMTP.Port = uint16(server.PortNumber())
		cfg.SMTP.UseTLS = false
		cfg.SMTP.From = from

		provider := &testConfigProvider{config: cfg}
		mailer := NewSmtpMailer(provider)

		ctx := context.Background()
		err := mailer.Send(ctx, from, []string{to}, msg)
		require.NoError(t, err)

		messages, err := server.WaitForMessagesAndPurge(1, 2*time.Second)
		require.NoError(t, err)
		require.Len(t, messages, 1)

		received := messages[0]

		// Envelope sender: extract from the full command.
		mailFromCmd := received.MailfromRequest()
		extractedFrom := extractAddressFromMailFrom(mailFromCmd)
		require.Equal(t, from, extractedFrom)

		// Envelope recipients: each element is [request, response].
		rcptPairs := received.RcpttoRequestResponse()
		require.Len(t, rcptPairs, 1)
		rcptRequest := rcptPairs[0][0]
		extractedTo := extractAddressFromRcptTo(rcptRequest)
		require.Equal(t, to, extractedTo)

		// Message body (headers + content).
		rawMsg := received.MsgRequest()
		require.Contains(t, rawMsg, body)
		require.Contains(t, rawMsg, subject)
	})

	t.Run("send with authentication (StartTLS) - skipped because mock does not support STARTTLS", func(t *testing.T) {
		t.Skip("go-smtp-mock does not support STARTTLS; skipping authentication test")
	})

	t.Run("implicit TLS on port 465 - skipped because mock does not support TLS", func(t *testing.T) {
		t.Skip("implicit TLS requires a TLS-enabled mock server; skipping")
	})

	t.Run("disabled SMTP", func(t *testing.T) {
		cfg := cfg.AppConfig{}
		cfg.SMTP.Enable = false
		provider := &testConfigProvider{config: cfg}
		mailer := NewSmtpMailer(provider)

		err := mailer.Send(context.Background(), from, []string{to}, msg)
		require.Error(t, err)
		require.Contains(t, err.Error(), "SMTP is disabled")
	})

	t.Run("missing host or port", func(t *testing.T) {
		cfg := cfg.AppConfig{}
		cfg.SMTP.Enable = true
		cfg.SMTP.Host = ""
		cfg.SMTP.Port = 2525
		provider := &testConfigProvider{config: cfg}
		mailer := NewSmtpMailer(provider)

		err := mailer.Send(context.Background(), from, []string{to}, msg)
		require.Error(t, err)
		require.Contains(t, err.Error(), "SMTP host or port not configured")

		cfg.SMTP.Host = "localhost"
		cfg.SMTP.Port = 0
		provider = &testConfigProvider{config: cfg}
		mailer = NewSmtpMailer(provider)
		err = mailer.Send(context.Background(), from, []string{to}, msg)
		require.Error(t, err)
		require.Contains(t, err.Error(), "SMTP host or port not configured")
	})

	t.Run("authentication requires TLS", func(t *testing.T) {
		cfg := cfg.AppConfig{}
		cfg.SMTP.Enable = true
		cfg.SMTP.Host = "localhost"
		cfg.SMTP.Port = 2525
		cfg.SMTP.UseTLS = false
		cfg.SMTP.Username = "user"
		cfg.SMTP.Password = "pass"
		provider := &testConfigProvider{config: cfg}
		mailer := NewSmtpMailer(provider)

		err := mailer.Send(context.Background(), from, []string{to}, msg)
		require.Error(t, err)
		require.Contains(t, err.Error(), "SMTP authentication requires TLS")
	})

	t.Run("username and password mismatch", func(t *testing.T) {
		cfg := cfg.AppConfig{}
		cfg.SMTP.Enable = true
		cfg.SMTP.Host = "localhost"
		cfg.SMTP.Port = 2525
		cfg.SMTP.UseTLS = true
		cfg.SMTP.Username = "user"
		cfg.SMTP.Password = ""
		provider := &testConfigProvider{config: cfg}
		mailer := NewSmtpMailer(provider)

		err := mailer.Send(context.Background(), from, []string{to}, msg)
		require.Error(t, err)
		require.Contains(t, err.Error(), "both username and password must be provided")
	})

	t.Run("context cancellation before send", func(t *testing.T) {
		server := startMockServer(t)

		cfg := cfg.AppConfig{}
		cfg.SMTP.Enable = true
		cfg.SMTP.Host = "127.0.0.1"
		cfg.SMTP.Port = uint16(server.PortNumber())
		cfg.SMTP.UseTLS = false
		cfg.SMTP.From = from

		provider := &testConfigProvider{config: cfg}
		mailer := NewSmtpMailer(provider)

		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		err := mailer.Send(ctx, from, []string{to}, msg)
		require.Error(t, err)
		require.True(t, strings.Contains(err.Error(), "canceled") ||
			strings.Contains(err.Error(), "operation was canceled"),
			"error should indicate cancellation, got: %v", err)
	})

	t.Run("context cancellation during send (with delay)", func(t *testing.T) {
		server := startMockServer(t, withResponseDelay(5))

		cfg := cfg.AppConfig{}
		cfg.SMTP.Enable = true
		cfg.SMTP.Host = "127.0.0.1"
		cfg.SMTP.Port = uint16(server.PortNumber())
		cfg.SMTP.UseTLS = false
		cfg.SMTP.From = from

		provider := &testConfigProvider{config: cfg}
		mailer := NewSmtpMailer(provider)

		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		err := mailer.Send(ctx, from, []string{to}, msg)
		require.Error(t, err)
		require.True(t,
			strings.Contains(err.Error(), "context deadline exceeded") ||
				strings.Contains(err.Error(), "i/o timeout") ||
				strings.Contains(err.Error(), "connection reset") ||
				strings.Contains(err.Error(), "broken pipe") ||
				strings.Contains(err.Error(), "use of closed network connection"),
			"error should indicate timeout/cancellation, got: %v", err)
	})

	t.Run("recipient rejected by server", func(t *testing.T) {
		badEmail := "bad@example.com"
		server := startMockServer(t, withBlacklistedRcptTo(badEmail))

		cfg := cfg.AppConfig{}
		cfg.SMTP.Enable = true
		cfg.SMTP.Host = "127.0.0.1"
		cfg.SMTP.Port = uint16(server.PortNumber())
		cfg.SMTP.UseTLS = false
		cfg.SMTP.From = from

		provider := &testConfigProvider{config: cfg}
		mailer := NewSmtpMailer(provider)

		ctx := context.Background()
		err := mailer.Send(ctx, from, []string{badEmail}, msg)
		require.Error(t, err)
		require.Contains(t, err.Error(), "failed to set recipient")
	})

	t.Run("recipient not registered (non-existent)", func(t *testing.T) {
		badEmail := "nobody@example.com"
		server := startMockServer(t, withNotRegisteredEmails(badEmail))

		cfg := cfg.AppConfig{}
		cfg.SMTP.Enable = true
		cfg.SMTP.Host = "127.0.0.1"
		cfg.SMTP.Port = uint16(server.PortNumber())
		cfg.SMTP.UseTLS = false
		cfg.SMTP.From = from

		provider := &testConfigProvider{config: cfg}
		mailer := NewSmtpMailer(provider)

		ctx := context.Background()
		err := mailer.Send(ctx, from, []string{badEmail}, msg)
		require.Error(t, err)
		require.Contains(t, err.Error(), "failed to set recipient")
	})
}

func TestDialSMTP(t *testing.T) {
	t.Skip("dialSMTP is internal and covered by Send tests")
}
