package config

import (
	"fmt"
	"os"
	"sync"

	"github.com/ReallyWeirdCat/brainiac/pkg/domain/config"
	"github.com/go-viper/mapstructure/v2"
	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/viper"
	"go.yaml.in/yaml/v3"
)

const defaultConfigLocation = "./config.yaml"

type viperConfig struct {
	mu         sync.Mutex
	loaded     bool
	configPath string
	cfg        config.AppConfig
}

func NewViperConfig() config.AppConfigProvider {
	return &viperConfig{configPath: defaultConfigLocation}
}

func NewViperConfigWithPath(configPath string) config.AppConfigProvider {
	return &viperConfig{configPath: configPath}
}

func (v *viperConfig) Get() config.AppConfig {
	v.mu.Lock()
	defer v.mu.Unlock()
	if v.loaded {
		return v.cfg
	}
	if err := v.load(); err != nil {
		panic(fmt.Errorf("configuration loading failed: %w", err))
	}
	v.loaded = true
	return v.cfg
}

func (v *viperConfig) load() error {
	if v.configPath == "" {
		v.configPath = defaultConfigLocation
	}

	vpr := viper.New()
	vpr.SetConfigFile(v.configPath)

	if err := vpr.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok && !os.IsNotExist(err) {
			return fmt.Errorf("error reading config file: %w", err)
		}
		if err := v.createDefaultConfigFile(v.configPath); err != nil {
			return err
		}
		return v.load()
	}

	var appCfg config.AppConfig
	appCfg.SetDefault()

	// Unmarshal YAML into the struct (mapstructure tags match YAML keys).
	decoderOpt := viper.DecoderConfigOption(func(c *mapstructure.DecoderConfig) {
		c.TagName = "yaml"
		c.WeaklyTypedInput = true // allows string->int conversions, etc.
	})
	if err := vpr.Unmarshal(&appCfg, decoderOpt); err != nil {
		return fmt.Errorf("unable to decode into config struct: %w", err)
	}

	// Apply environment variable overrides (envconfig tags).
	if err := envconfig.Process("", &appCfg); err != nil {
		return fmt.Errorf("processing environment variables: %w", err)
	}

	if err := appCfg.Validate(); err != nil {
		return fmt.Errorf("configuration validation failed: %w", err)
	}

	v.cfg = appCfg
	return nil
}

func (v *viperConfig) createDefaultConfigFile(path string) error {
	if _, err := os.Stat(path); err == nil {
		return fmt.Errorf("configuration file %q already exists", path)
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("checking file %q: %w", path, err)
	}

	var defaultCfg config.AppConfig
	defaultCfg.SetDefault()

	yamlBytes, err := yaml.Marshal(defaultCfg)
	if err != nil {
		return fmt.Errorf("marshal to YAML: %w", err)
	}

	header := []byte("# Default configuration for Brainiac.\n")
	out := append(header, yamlBytes...)

	if err := os.WriteFile(path, out, 0644); err != nil {
		return fmt.Errorf("write file %q: %w", path, err)
	}
	return nil
}

var _ config.AppConfigProvider = &viperConfig{}
