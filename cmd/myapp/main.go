package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Post struct {
	UserID int    `json:"userid"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type HTTPClient struct {
	BaseURL string
	client  *http.Client
}

func (c *HTTPClient) GetPost(ctx context.Context, id int) (*Post, error) {
	url := fmt.Sprintf("%sposts/%d", c.BaseURL, id)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("enexpected status: %s", resp.Status)
	}

	var post Post
	if err := json.NewDecoder(resp.Body).Decode(&post); err != nil {
		return nil, err
	}
	return &post, nil
}

func main() {
	client := &HTTPClient{
		BaseURL: "https://jsonplaceholder.typicode.com/",
		client: &http.Client{
			Timeout: 20 * time.Second,
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	post, err := client.GetPost(ctx, 1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Post #%d:\nTitle: %s\nBody: %s\n", post.ID, post.Title, post.Body)
}
