package binance_api

import (
	"context"
	"time"

	"github.com/adshao/go-binance/v2"
)

type Client struct {
	client *binance.Client
	ctx    context.Context
	cancel context.CancelFunc
}

func NewClient(apiKey string, secretKey string, useTestnet bool) *Client {
	if useTestnet {
		binance.UseTestnet = true
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	return &Client{
		client: binance.NewClient(apiKey, secretKey),
		ctx:    ctx,
		cancel: cancel,
	}
}

func (c *Client) GetAPIKeyPermission() (*binance.APIKeyPermission, error) {
	APIKeyPermission, err := c.client.NewGetAPIKeyPermission().Do(c.ctx)
	return APIKeyPermission, err
}
