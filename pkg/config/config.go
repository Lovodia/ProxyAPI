package config

import (
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
}

func LoadConfig(filename string) (Config, error) {
	var cfg Config

	data, err := os.ReadFile(filename)
	if err != nil {
		return cfg, err
	}
	err = yaml.Unmarshal(data, &cfg)
	return cfg, err
}
