
# Binance
[![Go Report Card](https://goreportcard.com/badge/github.com/nickcorin/binance)](https://goreportcard.com/report/github.com/nickcorin/binance)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/nickcorin/binance)

## About the package
This package provides an unofficial client library to interface with Binance's REST API. I am in no way affiliated with Binance.

**This package is a work in progress, and not fully tested. Use at your own risk.**

I am following the API documentation for Binance found [here](https://github.com/binance-exchange/binance-official-api-docs/blob/master/rest-api.md).

Contributions are welcome.

## Example Usage

```go
import (
	"github.com/nickcorin/binance"
)

func main() {
	// Create a new client with an API Key and Secret Key.
	client := binance.NewClient(
		binance.WithAPIKey("MyKey"),
		binance.WithSecretKey("MySecret"),
	)

	// Fetch your account information.
	info, err := client.AccountInfo(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// Fetch the most recent candlesticks.
	klines, err := client.Klines(context.Background(), binance.ETHBTC,
		binance.FiveMinutes, 10)
	if err != nil {
		log.Fatal(err)
	}
}
```

## Supported Endpoints
### Public API

- [x] Ping
- [x] Server Time
- [ ] Exchange Info

### Market Data

- [x] Order Book
- [x] Recent Trades
- [x] Old Trade Data
- [x] Aggregated Trades
- [x] Kline / Candlestick Data
- [x] Current Average Price
- [x] 24 Hour Ticker
- [x] Price Ticker
- [x] Order Book Ticker

### Account

- [ ] New Order
- [ ] Test New Order
- [ ] Query Order
- [ ] Cancel Order
- [ ] Current Open Orders
- [ ] All Orders
- [ ] New OCO
- [ ] Cancel OCO
- [ ] Query OCO
- [ ] Query All OCO
- [ ] Query Open OCO
- [x] Account Info
- [ ] Account Trade List

## Donations

If this package helped you out, feel free to donate.

- BTC: 34vfza4AWUsYtSG2zheNntrUtHircvkgUC
- BCH: pq8lwev3qykw3yqcfrrcelzafd0uact3wyjp3j3pnc
- ETH: 0xa9a84846b43FAb5d3694222De29F0973FDfc07a2
