package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Lovodia/ProxyAPI/internal/client"
	"github.com/Lovodia/ProxyAPI/internal/handlers"
	"github.com/Lovodia/ProxyAPI/pkg/config"
	"github.com/Lovodia/ProxyAPI/pkg/logger"
	"github.com/joho/godotenv"
)

func main() {
	logger.Init()

	_ = godotenv.Load(".env")

	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		logger.Error.Println("failed to load config.yaml", err)
	}

	apiURL := cfg.API.BaseURL
	if envURL := os.Getenv("API_BASE_URL"); envURL != "" {
		apiURL = envURL
	}
	if apiURL == "" {
		apiURL = "https://jsonplaceholder.typicode.com/"
	}

	port := cfg.Server.Port
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}
	if port == "" {
		port = "8080"
	}

	httpClient := client.NewClient(apiURL)
	h := handlers.New(httpClient)

	http.HandleFunc("/post", h.GetPostHandler)

	logger.Info.Printf("The server is running on port : %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
