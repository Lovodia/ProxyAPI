package client

import (
	"context"
	"fmt"
	"time"

	"github.com/Lovodia/ProxyAPI/internal/models"
	"github.com/go-resty/resty/v2"
)

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

func (c *HTTPClient) GetPost(ctx context.Context, id int) (post *models.Post, err error) {
	url := fmt.Sprintf("%sposts/%d", c.BaseURL, id)

	post = &models.Post{}
	resp, err := c.client.R().
		SetContext(ctx).
		SetResult(post).
		Get(url)

	if err != nil {
		err = fmt.Errorf("request error: %w", err)
		return
	}

	if resp.IsError() {
		err = fmt.Errorf("unexpected status: %s", resp.Status())
		return
	}

	return
}
