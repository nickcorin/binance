package binance

import (
	"context"
	"time"
)

// Client provides the methods relating to Binance's REST API.
type Client interface {
	AveragePrice(context.Context, Symbol) (*AveragePrice, error)
	HistoricalTrades(context.Context, Symbol, int) ([]Trade, error)
	HistoricalTradesFrom(context.Context, Symbol, int, int64) ([]Trade, error)
	Klines(context.Context, Symbol, Interval, int) ([]Kline, error)
	KlinesBetween(context.Context, Symbol, Interval, time.Time, time.Time, int) ([]Kline, error)
	OrderBook(context.Context, Symbol, int) (*OrderBook, error)
	Ping(context.Context) error
	ServerTime(context.Context) (time.Time, error)
	RecentTrades(context.Context, Symbol, int) ([]Trade, error)
}
