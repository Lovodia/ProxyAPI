package parse

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/Lovodia/ProxyAPI/pkg/config"
	"github.com/joho/godotenv"
)

type ParsedConfig struct {
	APIBaseURL string
	Port       string
	LogLevel   string
	Timeout    time.Duration
	RetryCount int
}

func LoadConfig() (*ParsedConfig, error) {
	cfg := &ParsedConfig{}

	timeoutStr := os.Getenv("TIMEOUT_SECONDS")
	if timeoutStr == "" {
		timeoutStr = ""
	}
	timeoutSec, err := strconv.Atoi(timeoutStr)
	if err != nil {
		timeoutSec = 10
	}
	cfg.Timeout = time.Duration(timeoutSec) * time.Second

	retryStr := os.Getenv("RETRY_COUNT")
	if retryStr == "" {
		retryStr = "3"
	}
	retryCount, err := strconv.Atoi(retryStr)
	if err != nil {
		retryCount = 3
	}
	cfg.RetryCount = retryCount

	envPath := os.Getenv("ENV_PATH")
	if envPath == "" {
		envPath = ".env"
	}
	_ = godotenv.Load(envPath)

	cfg.APIBaseURL = os.Getenv("API_BASE_URL")
	cfg.Port = os.Getenv("PORT")
	cfg.LogLevel = os.Getenv("LOG_LEVEL")

	if cfg.APIBaseURL != "" && cfg.Port != "" {
		return cfg, nil
	}

	yamlFiles, err := filepath.Glob("*.y*ml")
	if err != nil {
		return nil, fmt.Errorf("failed to scan YAML files: %w", err)
	}

	for _, file := range yamlFiles {
		c, err := config.LoadConfig(file)
		if err != nil {
			continue
		}

		if cfg.APIBaseURL == "" {
			cfg.APIBaseURL = c.API.BaseURL
		}
		if cfg.Port == "" {
			cfg.Port = c.Server.Port
		}
		if cfg.LogLevel == "" {
			cfg.LogLevel = c.Log.Level
		}

		if cfg.APIBaseURL != "" && cfg.Port != "" {
			return cfg, nil
		}
	}

	return nil, fmt.Errorf("missing required config: API_BASE_URL and/or PORT")
}
