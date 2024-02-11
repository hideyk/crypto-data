package binance

import (
	"github.com/adshao/go-binance/v2"
	"github.com/adshao/go-binance/v2/delivery"
	"github.com/adshao/go-binance/v2/futures"
)

type Client struct {
	client         *binance.Client
	futuresClient  *futures.Client
	deliveryClient *delivery.Client
}

func NewClient(apiKey string, secretKey string, useTestnet ...bool) *Client {
	if useTestnet {
		binance.UseTestnet = true
		futures.UseTestnet = true
		delivery.UseTestnet = true
	}
	return &Client{
		client:         binance.NewClient(apiKey, secretKey),
		futuresClient:  binance.NewFuturesClient(apiKey, secretKey),
		deliveryClient: binance.NewDeliveryClient(apiKey, secretKey),
	}
}
