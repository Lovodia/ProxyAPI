// @title Proxy API
// @version 1.0
// @description Пример прокси сервера на Go с документацией Swagger
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.example.com/support
// @contact.email support@example.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /
package main

import (
	"log/slog"
	"net/http"
	"os"

	_ "github.com/Lovodia/ProxyAPI/docs"
	"github.com/Lovodia/ProxyAPI/internal/client"
	"github.com/Lovodia/ProxyAPI/internal/handlers"
	"github.com/Lovodia/ProxyAPI/internal/parse"
	"github.com/Lovodia/ProxyAPI/pkg/logger"
	httpSwagger "github.com/swaggo/http-swagger"
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

	httpClient := client.NewClient(cfg)
	h := handlers.New(httpClient)

	http.HandleFunc("/post", h.GetPostHandler)
	http.Handle("/swagger/", httpSwagger.WrapHandler)

	slog.Info("The server is running", "port", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, nil); err != nil {
		slog.Error("server failed", "error", err)
	}
}
