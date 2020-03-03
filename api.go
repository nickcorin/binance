package binance

import (
	"context"
	"time"
)

// Client provides the methods relating to Binance's REST API.
type Client interface {
	Ping(context.Context) error
	OrderBook(context.Context, Symbol, int) (*OrderBook, error)
	ServerTime(context.Context) (time.Time, error)
	RecentTrades(context.Context, Symbol, int) ([]Trade, error)
}
