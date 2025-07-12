package client

import (
	"context"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
)

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type HTTPClient struct {
	BaseURL string
	client  *resty.Client
}

func NewClient(BaseURL string) *HTTPClient {
	client := resty.New().
		SetTimeout(10*time.Second).
		SetRetryCount(3).
		SetHeader("Accept", "application/json")

	return &HTTPClient{
		BaseURL: BaseURL,
		client:  client,
	}
}

func (c *HTTPClient) GetPost(ctx context.Context, id int) (*Post, error) {
	url := fmt.Sprintf("%sposts/%d", c.BaseURL, id)

	var post Post
	resp, err := c.client.R().
		SetContext(ctx).
		SetResult(&post).
		Get(url)

	if err != nil {
		return nil, fmt.Errorf("request error: %w", err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("enexpected status: %s", resp.Status())
	}
	return &post, nil
}
