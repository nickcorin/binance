package binance

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/luno/jettison/errors"
)

// OrderBook returns the current state of the exchange's order book with a list
// of bids and asks. All orders at the same price are aggreated together.
//
// Valid limits are 5, 10, 20, 50, 100, 500, 1000, 5000.
func (c *client) OrderBook(ctx context.Context, symbol Symbol, limit int) (
	*OrderBook, error) {
	res, err := c.get(ctx, fmt.Sprintf("/depth?symbol=%s&limit=%d",
		symbol.String(), limit))
	if err != nil {
		return nil, err
	}

	var book OrderBook
	if err = json.Unmarshal(res, &book); err != nil {
		return nil, errors.Wrap(err, "failed to parse order book")
	}

	return &book, nil
}

// Recent trades returns a list of recent trades executed on the exchange. Max
// limit is 1000.
func (c *client) RecentTrades(ctx context.Context, symbol Symbol, limit int) (
	[]Trade, error) {
	res, err := c.get(ctx, fmt.Sprintf("/trades?symbol=%s&limit=%d",
		symbol.String(), limit))
	if err != nil {
		return nil, err
	}

	var trades []Trade
	if err = json.Unmarshal(res, &trades); err != nil {
		return nil, errors.Wrap(err, "failed to parse trades")
	}

	return trades, nil
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
			symbol.String(), interval, limit))
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
			symbol.String(), interval, limit, start, end))
	if err != nil {
		return nil, err
	}

	var klines []Kline
	if err = json.Unmarshal(res, &klines); err != nil {
		return nil, errors.Wrap(err, "failed to parse klines")
	}

	return klines, err
}

// AveragePrice returns an aggregation of price movements over a period of time.
func (c *client) AveragePrice(ctx context.Context, symbol Symbol) (*AveragePrice,
	error) {
	res, err := c.get(ctx, fmt.Sprintf("/avgPrice&symbol=%s",
		symbol.String()))
	if err != nil {
		return nil, err
	}

	var price AveragePrice
	if err = json.Unmarshal(res, &price); err != nil {
		return nil, errors.Wrap(err, "failed to parse average price")
	}

	return &price, err
}

// TickerStats returns 24 hour rolling price change statistics for a given
// Symbols.
func (c *client) TickerStats(ctx context.Context, symbol Symbol) (*TickerStats,
	error) {
	res, err := c.get(ctx, fmt.Sprintf("/ticker/24h?symbol=%s",
		symbol.String()))
	if err != nil {
		return nil, err
	}

	var stats TickerStats
	if err = json.Unmarshal(res, &stats); err != nil {
		return nil, errors.Wrap(err, "failed to parse ticker stats")
	}

	return &stats, nil
}
