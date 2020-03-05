
# Binance
[![Go Report Card](https://goreportcard.com/badge/github.com/nickcorin/binance)](https://goreportcard.com/report/github.com/nickcorin/binance)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/nickcorin/binance)

## About the package
This package provides a client library to interface with Binance's REST API.

**This package is a work in progress.**

The package is incomplete and not fully tested. I am following the API documentation for Binance found [here](https://github.com/binance-exchange/binance-official-api-docs/blob/master/rest-api.md).

Contributions are welcome.

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
