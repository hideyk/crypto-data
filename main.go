package main

import (
	"fmt"

	"github.com/hideyk/crypto-data/exchange/binance_api"
)

func main() {
	binance_args := NewBinanceArgs()
	fmt.Printf("Binance args: %+v", binance_args)
	binance_client := binance_api.NewClient(
		binance_args.ApiKey,
		binance_args.SecretKey,
		binance_args.UseTestnet)

	fmt.Printf("%+v", binance_client)
}
