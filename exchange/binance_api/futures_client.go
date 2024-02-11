package binance_api

import (
	"github.com/adshao/go-binance/v2"
	"github.com/adshao/go-binance/v2/futures"
)

type FuturesClient struct {
	client *futures.Client
}

func NewFuturesClient(apiKey string, secretKey string, useTestnet bool) *FuturesClient {
	if useTestnet {
		futures.UseTestnet = true
	}
	return &FuturesClient{
		client: binance.NewFuturesClient(apiKey, secretKey),
	}
}
