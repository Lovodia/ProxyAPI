package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Lovodia/ProxyAPI/internal/client"
	"github.com/Lovodia/ProxyAPI/internal/handlers"
)

func main() {

	BaseURL := os.Getenv("API_BASE_URL")
	if BaseURL == "" {
		BaseURL = "https://jsonplaceholder.typicode.com/"
	}

	httpClient := client.NewClient(BaseURL)

	h := handlers.NewHandler(httpClient)

	http.HandleFunc("/post", h.GetPostHandler)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
