package binance

import (
	"context"
	"time"
)

// Client provides the methods relating to Binance's REST API.
type Client interface {
	AccountInfo(context.Context) (*AccountInfo, error)
	AggregateTrades(context.Context, Symbol, int) ([]AggregateTrade, error)
	AggregateTradesAfter(context.Context, Symbol, time.Time, int) ([]AggregateTrade, error)
	AggregateTradesBetween(context.Context, Symbol, time.Time, time.Time, int) ([]AggregateTrade, error)
	AggregateTradesFrom(context.Context, Symbol, int64, int) ([]AggregateTrade, error)
	AveragePrice(context.Context, Symbol) (*AveragePrice, error)
	HistoricalTrades(context.Context, Symbol, int) ([]Trade, error)
	HistoricalTradesFrom(context.Context, Symbol, int, int64) ([]Trade, error)
	Klines(context.Context, Symbol, Interval, int) ([]Kline, error)
	KlinesBetween(context.Context, Symbol, Interval, time.Time, time.Time, int) ([]Kline, error)
	LimitMaker(context.Context, Symbol, Side, float64, float64) (*OrderAck, error)
	LimitOrder(context.Context, Symbol, Side, float64, float64, TimeInForce) (*OrderAck, error)
	ListOrderBookTickers(context.Context) ([]OrderBookTicker, error)
	ListPriceTickers(context.Context) ([]PriceTicker, error)
	ListTickerStats(context.Context) ([]TickerStats, error)
	MarketOrder(context.Context, Symbol, Side, float64) (*OrderAck, error)
	MarketOrderSpend(context.Context, Symbol, Side, float64) (*OrderAck, error)
	OrderBook(context.Context, Symbol, int) (*OrderBook, error)
	OrderBookTicker(context.Context, Symbol) (*OrderBookTicker, error)
	Ping(context.Context) error
	PriceTicker(context.Context, Symbol) (*PriceTicker, error)
	ServerTime(context.Context) (time.Time, error)
	StopLossLimitOrder(context.Context, Symbol, Side, float64, float64, float64, TimeInForce) (*OrderAck, error)
	StopLossOrder(context.Context, Symbol, Side, float64, float64) (*OrderAck, error)
	TakeProfitLimitOrder(context.Context, Symbol, Side, float64, float64, float64, TimeInForce) (*OrderAck, error)
	TakeProfitOrder(context.Context, Symbol, Side, float64, float64) (*OrderAck, error)
	TickerStats(context.Context, Symbol) (*TickerStats, error)
	RecentTrades(context.Context, Symbol, int) ([]Trade, error)
}
