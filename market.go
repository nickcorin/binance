package binance

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/luno/jettison/errors"
)

// AggregateTrades returns a list of the most recent trades. Trades are
// aggregated if they were executed as part of the same order, at the same time
// and for the same price for a given Symbol. Max limit is 1000.
func (c *client) AggregateTrades(ctx context.Context, symbol Symbol,
	limit int) ([]AggregateTrade, error) {
	res, err := c.get(ctx, fmt.Sprintf("/aggTrades?symbol=%s&limit=%d",
		string(symbol), limit))
	if err != nil {
		return nil, err
	}

	var trades []AggregateTrade
	if err = json.Unmarshal(res, &trades); err != nil {
		return nil, errors.Wrap(err, "failed to parse aggregate trades")
	}

	return trades, err
}

// AggregateTradesAfter returns a list of trades that executed after a given
// time inclusive. Trades are aggregated if they were executed as part of the
// same order, at the same time and for the same price for a given Symbol. Max
// limit is 1000.
func (c *client) AggregateTradesAfter(ctx context.Context, symbol Symbol,
	from time.Time, limit int) ([]AggregateTrade, error) {
	res, err := c.get(ctx,
		fmt.Sprintf("/aggTrades?symbol=%s&startTime=%d&limit=%d",
			string(symbol), from.UnixNano()/1e6, limit))
	if err != nil {
		return nil, err
	}

	var trades []AggregateTrade
	if err = json.Unmarshal(res, &trades); err != nil {
		return nil, errors.Wrap(err, "failed to parse aggregate trades")
	}

	return trades, nil
}

// AggregateTradesBetween returns a list of trades that executed between a
// given time interval inclusive. Trades are aggregated if they were executed
// as part of the same order, at the same time and for the same price for a
// given Symbol. Max limit is 1000.
func (c *client) AggregateTradesBetween(ctx context.Context, symbol Symbol,
	from time.Time, to time.Time, limit int) ([]AggregateTrade, error) {
	res, err := c.get(ctx,
		fmt.Sprintf("/aggTrades?symbol=%s&startTime=%d&endTime=%d&limit=%d",
			string(symbol), from.UnixNano()/1e6, to.UnixNano()/1e6, limit))
	if err != nil {
		return nil, err
	}

	var trades []AggregateTrade
	if err = json.Unmarshal(res, &trades); err != nil {
		return nil, errors.Wrap(err, "failed to parse aggregate trades")
	}

	return trades, nil
}

// AggregateTradesFrom returns a list of trades that executed after a given
// trade ID inclusive. Trades are aggregated if they were executed as part of
// the same order, at the same time and for the same price for a given Symbol.
// Max limit is 1000.
func (c *client) AggregateTradesFrom(ctx context.Context, symbol Symbol,
	from int64, limit int) ([]AggregateTrade, error) {
	res, err := c.get(ctx,
		fmt.Sprintf("/aggTrades?symbol=%s&fromId=%d&limit=%d",
			string(symbol), from, limit))
	if err != nil {
		return nil, err
	}

	var trades []AggregateTrade
	if err = json.Unmarshal(res, &trades); err != nil {
		return nil, errors.Wrap(err, "failed to parse aggregate trades")
	}

	return trades, nil
}

// AveragePrice returns an aggregation of price movements over a period of time.
func (c *client) AveragePrice(ctx context.Context, symbol Symbol) (*AveragePrice,
	error) {
	res, err := c.get(ctx, fmt.Sprintf("/avgPrice&symbol=%s",
		string(symbol)))
	if err != nil {
		return nil, err
	}

	var price AveragePrice
	if err = json.Unmarshal(res, &price); err != nil {
		return nil, errors.Wrap(err, "failed to parse average price")
	}

	return &price, err
}

// HistoricalTrades returns a historical list of trades executed on the
// exchange. Max limit is 1000.
func (c *client) HistoricalTrades(ctx context.Context, symbol Symbol,
	limit int) ([]Trade, error) {
	res, err := c.get(ctx, fmt.Sprintf("/historicalTrades?symbol=%s&limit=%d",
		symbol, limit))
	if err != nil {
		return nil, err
	}

	var trades []Trade
	if err = json.Unmarshal(res, &trades); err != nil {
		return nil, errors.Wrap(err, "failed to parse trades")
	}

	return trades, err
}

// HistoricalTradesFrom returns a historical list of trades executed on the
// exchange after a given trade ID. Max limit is 1000.
func (c *client) HistoricalTradesFrom(ctx context.Context, symbol Symbol,
	limit int, from int64) ([]Trade, error) {
	res, err := c.get(ctx,
		fmt.Sprintf("/historicalTrades?symbol=%s&limit=%d&fromId=%d",
			symbol, limit, from))
	if err != nil {
		return nil, err
	}

	var trades []Trade
	if err = json.Unmarshal(res, &trades); err != nil {
		return nil, errors.Wrap(err, "failed to parse trades")
	}

	return trades, err
}

// Klines returns candlestick data for a Symbol. Max limit is 1000.
func (c *client) Klines(ctx context.Context, symbol Symbol, interval Interval,
	limit int) ([]Kline, error) {
	res, err := c.get(ctx,
		fmt.Sprintf("/klines?symbol=%s&interval=%s&limit=%d",
			string(symbol), interval, limit))
	if err != nil {
		return nil, err
	}

	var klines []Kline
	if err = json.Unmarshal(res, &klines); err != nil {
		return nil, errors.Wrap(err, "failed to parse klines")
	}

	return klines, err
}

// KlinesBetween returns candlestick data for a Symbol between some interval
// of time. Max limit is 1000.
func (c *client) KlinesBetween(ctx context.Context, symbol Symbol,
	interval Interval, from, to time.Time, limit int) ([]Kline, error) {
	start := from.UnixNano() / 1e6
	end := to.UnixNano() / 1e6
	res, err := c.get(ctx,
		fmt.Sprintf("/klines?symbol=%s&interval=%s&limit=%d&startTime=%d&endTime=%d",
			string(symbol), interval, limit, start, end))
	if err != nil {
		return nil, err
	}

	var klines []Kline
	if err = json.Unmarshal(res, &klines); err != nil {
		return nil, errors.Wrap(err, "failed to parse klines")
	}

	return klines, err
}

// ListOrderBookTickers returns the current best bid and ask prices for all
// Symbols.
func (c *client) ListOrderBookTickers(ctx context.Context) (
	[]OrderBookTicker, error) {
	res, err := c.get(ctx, "/ticker/bookTicker")
	if err != nil {
		return nil, err
	}

	var tickers []OrderBookTicker
	if err = json.Unmarshal(res, &tickers); err != nil {
		return nil, errors.Wrap(err, "failed to parse order book tickers")
	}

	return tickers, nil
}

// PriceTicker returns the latest price tickers for all Symbols.
func (c *client) ListPriceTickers(ctx context.Context) (
	[]PriceTicker, error) {
	res, err := c.get(ctx, "/ticker/price")
	if err != nil {
		return nil, err
	}

	var tickers []PriceTicker
	if err = json.Unmarshal(res, &tickers); err != nil {
		return nil, errors.Wrap(err, "failed to parse price ticker")
	}

	return tickers, nil
}

// ListTickerStats returns 24 hour rolling price change statistics for all
// Symbols.
func (c *client) ListTickerStats(ctx context.Context) (
	[]TickerStats, error) {
	res, err := c.get(ctx, "/ticker/24h")
	if err != nil {
		return nil, err
	}

	var stats []TickerStats
	if err = json.Unmarshal(res, &stats); err != nil {
		return nil, errors.Wrap(err, "failed to parse ticker stats")
	}

	return stats, nil
}

// OrderBook returns the current state of the exchange's order book with a list
// of bids and asks. All orders at the same price are aggreated together.
//
// Valid limits are 5, 10, 20, 50, 100, 500, 1000, 5000.
func (c *client) OrderBook(ctx context.Context, symbol Symbol, limit int) (
	*OrderBook, error) {
	res, err := c.get(ctx, fmt.Sprintf("/depth?symbol=%s&limit=%d",
		string(symbol), limit))
	if err != nil {
		return nil, err
	}

	var book OrderBook
	if err = json.Unmarshal(res, &book); err != nil {
		return nil, errors.Wrap(err, "failed to parse order book")
	}

	return &book, nil
}

// OrderBookTicker returns the current best bid and ask price for a Symbol.
func (c *client) OrderBookTicker(ctx context.Context, symbol Symbol) (
	*OrderBookTicker, error) {
	res, err := c.get(ctx, fmt.Sprintf("/ticker/bookTicker?symbol=%s",
		string(symbol)))
	if err != nil {
		return nil, err
	}

	var ticker OrderBookTicker
	if err = json.Unmarshal(res, &ticker); err != nil {
		return nil, errors.Wrap(err, "failed to parse order book ticker")
	}

	return &ticker, nil
}

// PriceTicker returns the latest price for a given Symbol.
func (c *client) PriceTicker(ctx context.Context, symbol Symbol) (
	*PriceTicker, error) {
	res, err := c.get(ctx, fmt.Sprintf("/ticker/price?symbol=%s",
		string(symbol)))
	if err != nil {
		return nil, err
	}

	var ticker PriceTicker
	if err = json.Unmarshal(res, &ticker); err != nil {
		return nil, errors.Wrap(err, "failed to parse price ticker")
	}

	return &ticker, nil
}

// TickerStats returns 24 hour rolling price change statistics for a given
// Symbol.
func (c *client) TickerStats(ctx context.Context, symbol Symbol) (*TickerStats,
	error) {
	res, err := c.get(ctx, fmt.Sprintf("/ticker/24h?symbol=%s",
		string(symbol)))
	if err != nil {
		return nil, err
	}

	var stats TickerStats
	if err = json.Unmarshal(res, &stats); err != nil {
		return nil, errors.Wrap(err, "failed to parse ticker stats")
	}

	return &stats, nil
}

// Recent trades returns a list of recent trades executed on the exchange. Max
// limit is 1000.
func (c *client) RecentTrades(ctx context.Context, symbol Symbol, limit int) (
	[]Trade, error) {
	res, err := c.get(ctx, fmt.Sprintf("/trades?symbol=%s&limit=%d",
		string(symbol), limit))
	if err != nil {
		return nil, err
	}

	var trades []Trade
	if err = json.Unmarshal(res, &trades); err != nil {
		return nil, errors.Wrap(err, "failed to parse trades")
	}

	return trades, nil
}
