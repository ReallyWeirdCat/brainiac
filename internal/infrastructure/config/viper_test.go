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
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ReallyWeirdCat/brainiac/pkg/domain/config"
	"go.yaml.in/yaml/v3"
)

// helper to write a config YAML file
func writeConfigFile(t *testing.T, dir, name, content string) string {
	t.Helper()
	path := filepath.Join(dir, name)
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatalf("failed to write config file: %v", err)
	}
	return path
}

// helper to recover from panic and return the recovered value
func recoverPanic(t *testing.T, fn func()) (recovered interface{}) {
	t.Helper()
	defer func() {
		recovered = recover()
	}()
	fn()
	return nil
}

func TestNewViperConfig_Get_LoadsFromFile(t *testing.T) {
	tmpDir := t.TempDir()
	cfgContent := `
registration:
  enable: true
  require_email: true
login:
  enforce_email: true
security:
  passwords:
    compromised:
      check_compromised_passwords: true
      compromised_passwords_file_path: "/path/to/file"
      compromised_passwords_repo_url: "https://repo.url"
smtp:
  enable: true
  host: "smtp.example.com"
  port: 587
  username: "user"
  password: "secret"
  use_tls: false
  from: "from@example.com"
database:
  uri: "postgres://test:test@localhost/testdb"
`
	cfgPath := writeConfigFile(t, tmpDir, "config.yaml", cfgContent)

	provider := NewViperConfig(cfgPath)
	cfg := provider.Get()

	if !cfg.Registration.Enable {
		t.Error("expected Registration.Enable true")
	}
	if !cfg.Registration.RequireEmail {
		t.Error("expected Registration.RequireEmail true")
	}
	if !cfg.Login.EnforceEmail {
		t.Error("expected Login.EnforceEmail true")
	}
	if !cfg.Security.Passwords.Compromised.CheckPasswords {
		t.Error("expected CheckPasswords true")
	}
	if cfg.Security.Passwords.Compromised.FilePath != "/path/to/file" {
		t.Errorf("expected FilePath /path/to/file, got %s", cfg.Security.Passwords.Compromised.FilePath)
	}
	if cfg.Security.Passwords.Compromised.RepoURL != "https://repo.url" {
		t.Errorf("expected RepoURL https://repo.url, got %s", cfg.Security.Passwords.Compromised.RepoURL)
	}
	if !cfg.SMTP.Enable {
		t.Error("expected SMTP.Enable true")
	}
	if cfg.SMTP.Host != "smtp.example.com" {
		t.Errorf("expected SMTP.Host smtp.example.com, got %s", cfg.SMTP.Host)
	}
	if cfg.SMTP.Port != 587 {
		t.Errorf("expected SMTP.Port 587, got %d", cfg.SMTP.Port)
	}
	if cfg.SMTP.Username != "user" {
		t.Errorf("expected SMTP.Username user, got %s", cfg.SMTP.Username)
	}
	if cfg.SMTP.Password != "secret" {
		t.Errorf("expected SMTP.Password secret, got %s", cfg.SMTP.Password)
	}
	if cfg.SMTP.UseTLS {
		t.Error("expected UseTLS false")
	}
	if cfg.SMTP.From != "from@example.com" {
		t.Errorf("expected SMTP.From from@example.com, got %s", cfg.SMTP.From)
	}
	if cfg.Database.URI != "postgres://test:test@localhost/testdb" {
		t.Errorf("expected Database.URI postgres://test:test@localhost/testdb, got %s", cfg.Database.URI)
	}
}

func TestNewViperConfig_Get_CachesConfig(t *testing.T) {
	tmpDir := t.TempDir()
	cfgPath := writeConfigFile(t, tmpDir, "config.yaml", `
registration:
  enable: true
`)

	provider := NewViperConfig(cfgPath)
	cfg1 := provider.Get()

	// Remove the file to ensure second call does not re‑read
	if err := os.Remove(cfgPath); err != nil {
		t.Fatalf("failed to remove config file: %v", err)
	}

	cfg2 := provider.Get()

	if !cfg2.Registration.Enable {
		t.Error("cached config should still have Registration.Enable true")
	}
	// Verify it is the same instance (since cfg is stored by value, they should be equal)
	if cfg1 != cfg2 {
		t.Error("expected identical cached config")
	}
}

func TestNewViperConfig_Get_CreatesDefaultConfig(t *testing.T) {
	tmpDir := t.TempDir()
	cfgPath := filepath.Join(tmpDir, "nonexistent.yaml")

	provider := NewViperConfig(cfgPath)
	cfg := provider.Get()

	// Verify the file was created
	if _, err := os.Stat(cfgPath); os.IsNotExist(err) {
		t.Fatal("expected config file to be created")
	}

	// Verify defaults are applied
	defaults := config.AppConfig{}
	defaults.SetDefault()
	if cfg != defaults {
		t.Error("loaded config should match defaults")
	}

	// Verify file content contains defaults (header + YAML)
	fileContent, err := os.ReadFile(cfgPath)
	if err != nil {
		t.Fatalf("reading created config file: %v", err)
	}
	if !strings.Contains(string(fileContent), "# Default configuration for Brainiac.") {
		t.Error("created file missing header")
	}
	// Try to unmarshal the YAML portion (skip header)
	lines := strings.SplitN(string(fileContent), "\n", 2)
	if len(lines) < 2 {
		t.Fatal("created file too short")
	}
	var parsed config.AppConfig
	if err := yaml.Unmarshal([]byte(lines[1]), &parsed); err != nil {
		t.Fatalf("unmarshal created file: %v", err)
	}
	if parsed != defaults {
		t.Error("created file YAML content does not match defaults")
	}
}

func TestNewViperConfig_Get_EnvOverride(t *testing.T) {
	tmpDir := t.TempDir()
	// Provide a file with smtp.host = "filehost" – should be overridden by env
	cfgPath := writeConfigFile(t, tmpDir, "config.yaml", `
smtp:
  enable: true
  host: "filehost"
`)

	// Set env var SMTP_HOST (Viper automatic env binding)
	t.Setenv("SMTP_HOST", "envhost")
	// Also set another env var to override a field not present in file
	t.Setenv("SMTP_USERNAME", "envuser")

	provider := NewViperConfig(cfgPath)
	cfg := provider.Get()

	if cfg.SMTP.Host != "envhost" {
		t.Errorf("expected SMTP.Host 'envhost' from env, got '%s'", cfg.SMTP.Host)
	}
	if cfg.SMTP.Username != "envuser" {
		t.Errorf("expected SMTP.Username 'envuser' from env, got '%s'", cfg.SMTP.Username)
	}
}

func TestNewViperConfig_Get_ValidationError(t *testing.T) {
	tmpDir := t.TempDir()
	// SMTP enabled but missing host -> validation should fail
	cfgPath := writeConfigFile(t, tmpDir, "config.yaml", `
smtp:
  enable: true
  host: ""
`)

	panicVal := recoverPanic(t, func() {
		provider := NewViperConfig(cfgPath)
		provider.Get()
	})

	if panicVal == nil {
		t.Fatal("expected panic due to validation error")
	}

	panicMsg := fmt.Sprint(panicVal)
	if !strings.Contains(panicMsg, "SMTP host not specified") {
		t.Errorf("panic message does not mention missing host: %v", panicVal)
	}
}

func TestNewViperConfig_Get_InvalidYAML(t *testing.T) {
	tmpDir := t.TempDir()
	cfgPath := writeConfigFile(t, tmpDir, "config.yaml", `broken: yaml: :`)

	provider := NewViperConfig(cfgPath)
	panicVal := recoverPanic(t, func() {
		provider.Get()
	})

	if panicVal == nil {
		t.Fatal("expected panic due to invalid YAML")
	}
	panicMsg := fmt.Sprint(panicVal)
	if !strings.Contains(panicMsg, "error reading config file") {
		t.Errorf("expected 'error reading config file' in panic, got: %v", panicVal)
	}
}

func TestNewViperConfig_createDefaultConfigFile_AlreadyExists(t *testing.T) {
	tmpDir := t.TempDir()
	cfgPath := writeConfigFile(t, tmpDir, "config.yaml", "some: content")

	v := &viperConfig{configPath: cfgPath}
	err := v.createDefaultConfigFile(cfgPath)
	if err == nil {
		t.Error("expected error when file already exists")
	}
	if !strings.Contains(err.Error(), "already exists") {
		t.Errorf("error message should mention already exists, got: %v", err)
	}
}

func TestNewViperConfig_createDefaultConfigFile_Success(t *testing.T) {
	tmpDir := t.TempDir()
	cfgPath := filepath.Join(tmpDir, "new.yaml")

	v := &viperConfig{configPath: cfgPath}
	err := v.createDefaultConfigFile(cfgPath)
	if err != nil {
		t.Fatalf("createDefaultConfigFile failed: %v", err)
	}

	// Verify file exists
	if _, err := os.Stat(cfgPath); os.IsNotExist(err) {
		t.Fatal("file was not created")
	}

	// Read and check header + default YAML
	data, err := os.ReadFile(cfgPath)
	if err != nil {
		t.Fatalf("reading created file: %v", err)
	}
	if !strings.HasPrefix(string(data), "# Default configuration for Brainiac.") {
		t.Error("missing header in created file")
	}

	// Extract YAML part and compare to defaults
	parts := strings.SplitN(string(data), "\n", 2)
	if len(parts) < 2 {
		t.Fatal("created file too short")
	}
	var fromFile config.AppConfig
	if err := yaml.Unmarshal([]byte(parts[1]), &fromFile); err != nil {
		t.Fatalf("unmarshal created file: %v", err)
	}
	expected := config.AppConfig{}
	expected.SetDefault()
	if fromFile != expected {
		t.Errorf("default config in file does not match SetDefault() output")
	}
}

func TestNewViperConfig_Get_PanicsOnLoadError(t *testing.T) {
	// Path to a file that can't be read or a directory to provoke error.
	// For example, a directory instead of a file.
	tmpDir := t.TempDir()
	// Provide a directory as config "file"
	provider := NewViperConfig(tmpDir)
	panicVal := recoverPanic(t, func() {
		provider.Get()
	})
	if panicVal == nil {
		t.Fatal("expected panic when loading invalid config file")
	}
}

func TestNewViperConfig_Get_DefaultConfigValidation(t *testing.T) {
	// Ensure default config passes validation (Get should not panic)
	tmpDir := t.TempDir()
	cfgPath := filepath.Join(tmpDir, "config.yaml")

	provider := NewViperConfig(cfgPath)
	// Recover just in case
	panicVal := recoverPanic(t, func() {
		_ = provider.Get()
	})
	if panicVal != nil {
		t.Fatalf("default config failed validation: %v", panicVal)
	}
}

// Ensures viperConfig satisfies the interface (compile‑time check, already present in code)
func TestViperConfig_ImplementsProvider(t *testing.T) {
	var _ config.AppConfigProvider = &viperConfig{}
}
