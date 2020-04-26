
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

	// Create an order request.
	req := binance.NewOrderRequest{
		Price:			0.81,
		Side:			binance.OrderSideBuy,
		Symbol:			"ETHBTC",
		TimeInForce:	binance.GTC,
		Type:			binance.OrderTypeLimit,
		Quantity:		0.5,
	}

	// Test the order.
	err := client.NewOrderTest(context.Background(), &req)
	if err != nil {
		log.Fatal(err)
	}

	// Place the order.
	res, err := client.NewOrder(context.Background(), &req) 
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

- [ ] Order Book
- [ ] Recent Trades
- [ ] Old Trade Data
- [ ] Aggregated Trades
- [ ] Kline / Candlestick Data
- [ ] Current Average Price
- [ ] 24 Hour Ticker
- [ ] Price Ticker
- [ ] Order Book Ticker

### Account

- [x] New Order
- [x] Test New Order
- [x] Query Order
- [x] Cancel Order
- [ ] Current Open Orders
- [ ] All Orders
- [ ] New OCO
- [ ] Cancel OCO
- [ ] Query OCO
- [ ] Query All OCO
- [ ] Query Open OCO
- [ ] Account Info
- [ ] Account Trade List

## Donations

If this package helped you out, feel free to donate.

- BTC: 34vfza4AWUsYtSG2zheNntrUtHircvkgUC
- BCH: pq8lwev3qykw3yqcfrrcelzafd0uact3wyjp3j3pnc
- ETH: 0xa9a84846b43FAb5d3694222De29F0973FDfc07a2
