package main

import (
	"flag"
	"os"
	"strconv"
)

func StringArgParser(env_key, flag_key, description, preset string) string {
	envArg, _ := os.LookupEnv(env_key)
	arg := flag.String(flag_key, envArg, description)
	flag.Parse()
	return *arg
}

func BoolArgParser(env_key, flag_key, description string, preset bool) bool {
	envArg, _ := os.LookupEnv(env_key)
	arg := flag.String(flag_key, envArg, description)
	flag.Parse()
	val, ok := strconv.ParseBool(*arg)
	if ok != nil {
		return preset
	}
	return val
}

type BinanceArgs struct {
	ApiKey     string
	SecretKey  string
	UseTestnet bool
}

func NewBinanceArgs() *BinanceArgs {
	apiKey := StringArgParser("BINANCE_API_KEY", "binance_api_key", "Binance API Key", "EMPTY")
	secretKey := StringArgParser("BINANCE_SECRET_KEY", "binance_secret_key", "Binance Secret Key", "EMPTY")
	useTestnet := BoolArgParser("BINANCE_USE_TESTNET", "binance_use_testnet", "Use Binance Testnet", true)
	return &BinanceArgs{
		ApiKey:     apiKey,
		SecretKey:  secretKey,
		UseTestnet: useTestnet,
	}
}
