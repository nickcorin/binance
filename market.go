package binance

import (
	"context"
	"encoding/json"
	"fmt"

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
