package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`

	API struct {
		BaseURL string `yaml:"base_url"`
	} `yaml:"api"`

	Log struct {
		Level string `yaml:"level"`
	} `yaml:"log"`
}

func LoadConfig(filename string) (Config, error) {
	var cfg Config

	data, err := os.ReadFile(filename)
	if err != nil {
		return Config{}, fmt.Errorf("failed to read config file %q: %w", filename, err)
	}

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return Config{}, fmt.Errorf("failed to unmarshal YAML %q: %w", filename, err)
	}

	return cfg, nil
}
