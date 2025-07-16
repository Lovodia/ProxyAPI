package client

import (
	"context"
	"fmt"

	"github.com/Lovodia/ProxyAPI/internal/models"
	"github.com/Lovodia/ProxyAPI/internal/parse"
	"github.com/go-resty/resty/v2"
)

type PostGetter interface {
	GetPost(ctx context.Context, id int) (*models.Post, error) // для подмены клиента в тестах
}

type HTTPClient struct {
	BaseURL string
	client  *resty.Client
}

func NewClient(cfg *parse.ParsedConfig) *HTTPClient {
	client := resty.New().
		SetTimeout(cfg.Timeout).
		SetRetryCount(cfg.RetryCount).
		SetHeader("Accept", "application/json")

	return &HTTPClient{
		BaseURL: cfg.APIBaseURL,
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
