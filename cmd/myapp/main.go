package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/Lovodia/ProxyAPI/internal/client"
	"github.com/Lovodia/ProxyAPI/internal/handlers"
	"github.com/Lovodia/ProxyAPI/internal/parse"
	"github.com/Lovodia/ProxyAPI/pkg/logger"
)

func main() {
	cfg, err := parse.LoadConfig()
	if err != nil {
		slog.Error("failed to load config", "error", err)
		os.Exit(1)
	}

	logger := logger.NewLogger(cfg.LogLevel)
	slog.SetDefault(logger)

	slog.Info("Logger initialized", "level", cfg.LogLevel)

	httpClient := client.NewClient(cfg.APIBaseURL)
	h := handlers.New(httpClient)

	http.HandleFunc("/post", h.GetPostHandler)

	slog.Info("The server is running", "port", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, nil); err != nil {
		slog.Error("server failed", "error", err)
	}
}
