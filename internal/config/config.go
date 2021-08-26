package config

import (
	"github.com/mcuadros/go-defaults"
)

type Config struct {
	MainProxyTarget          string   `yaml:"main" default:"http://localhost:8888"`
	ListenAddress            string   `yaml:"listen" default:":8080"`
	TargetsEndpoint          string   `yaml:"targets" default:"targets"`
	TargetsListenAddress     string   `yaml:"targets-address"`
	PasswordFile             string   `yaml:"password"`
	PersistentFailureTimeout int      `yaml:"fail-after" default:"30"`
	RetryAfter               int      `yaml:"retry-after" default:"1"`
	Mirrors                  []string `yaml:"mirror"`
}

func (s *Config) UnmarshalYAML(unmarshal func(interface{}) error) error {
	defaults.SetDefaults(s)

	type cfg Config

	if err := unmarshal((*cfg)(s)); err != nil {
		return err
	}

	return nil
}

func Default() *Config {
	c := &Config{}
	defaults.SetDefaults(c)

	return c
}
