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

package config

import (
	"errors"
)

type AppConfigProvider interface {
	Get() AppConfig
}

type AppConfig struct {
	Registration struct {
		Enable       bool `yaml:"enable" envconfig:"REGISTRATION_ENABLE"`
		RequireEmail bool `yaml:"require_email" envconfig:"REGISTRATION_REQUIRE_EMAIL"`
	} `yaml:"registration"`
	Login struct {
		EnforceEmail bool `yaml:"enforce_email" envconfig:"LOGIN_ENFORCE_EMAIL"`
	} `yaml:"login"`
	Security struct {
		Passwords struct {
			Compromised struct {
				CheckPasswords bool   `yaml:"check_compromised_passwords" envconfig:"SECURITY_PASSWORDS_COMPROMISED_CHECK_PASSWORDS"`
				FilePath       string `yaml:"compromised_passwords_file_path" envconfig:"SECURITY_PASSWORDS_COMPROMISED_FILE_PATH"`
				RepoURL        string `yaml:"compromised_passwords_repo_url" envconfig:"SECURITY_PASSWORDS_COMPROMISED_REPO_URL"`
			} `yaml:"compromised"`
		} `yaml:"passwords"`
	} `yaml:"security"`
	SMTP struct {
		Enable   bool   `yaml:"enable" envconfig:"SMTP_ENABLE"`
		Host     string `yaml:"host" envconfig:"SMTP_HOST"`
		Port     uint16 `yaml:"port" envconfig:"SMTP_PORT"`
		Username string `yaml:"username" envconfig:"SMTP_USERNAME"`
		Password string `yaml:"password" envconfig:"SMTP_PASSWORD"`
		UseTLS   bool   `yaml:"use_tls" envconfig:"SMTP_USE_TLS"`
		From     string `yaml:"from" envconfig:"SMTP_FROM"`
	} `yaml:"smtp"`
	Database struct {
		URI string `yaml:"uri" envconfig:"DATABASE_URI"`
	} `yaml:"database"`
}

func (a *AppConfig) Validate() error {
	var errs []error

	if a.Security.Passwords.Compromised.CheckPasswords {
		if a.Security.Passwords.Compromised.FilePath == "" {
			errs = append(errs, errors.New("Security.Passwords.Compromised.FilePath not specified. Set valid file path or disable Security.Passwords.Compromised.CheckPasswords"))
		}
	}

	if a.SMTP.Enable {
		if a.SMTP.Host == "" {
			errs = append(errs, errors.New("SMTP host not specified"))
		}
		if a.SMTP.Port == 0 {
			errs = append(errs, errors.New("SMTP port not specified"))
		}
		if a.SMTP.Username == "" {
			errs = append(errs, errors.New("SMTP username not specified"))
		}
		if a.SMTP.From == "" {
			errs = append(errs, errors.New("SMTP from not specified"))
		}
	}

	if a.Database.URI == "" {
		errs = append(errs, errors.New("Database URI not specified"))
	}

	if len(errs) == 0 {
		return nil
	}

	return errors.Join(errs...)
}

func (a *AppConfig) SetDefault() {
	a.Registration.Enable = false
	a.Registration.RequireEmail = false

	a.Login.EnforceEmail = false

	a.Security.Passwords.Compromised.CheckPasswords = false
	a.Security.Passwords.Compromised.FilePath = ""
	a.Security.Passwords.Compromised.RepoURL = ""

	a.SMTP.Enable = false
	a.SMTP.Host = "example.com"
	a.SMTP.Port = 465
	a.SMTP.Username = "user"
	a.SMTP.Password = "password"
	a.SMTP.UseTLS = true
	a.SMTP.From = "user@example.com"

	a.Database.URI = "postgres://user:password@localhost:5432/dbname"
}
