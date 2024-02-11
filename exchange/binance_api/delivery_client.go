package binance_api

import (
	"github.com/adshao/go-binance/v2"
	"github.com/adshao/go-binance/v2/delivery"
)

type DeliveryClient struct {
	client *delivery.Client
}

func NewDeliveryClient(apiKey string, secretKey string, useTestnet bool) *DeliveryClient {
	if useTestnet {
		delivery.UseTestnet = true
	}
	return &DeliveryClient{
		client: binance.NewDeliveryClient(apiKey, secretKey),
	}
}
