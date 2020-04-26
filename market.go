package binance

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/luno/jettison/errors"
)

// OrderBookTicker contains the best price and quantity on an order book.
type OrderBookTicker struct {
	// AskPrice represents the lowest ask price in the order book.
	AskPrice string `json:"askPrice"`

	// AskQty represents the volume at the current ask price.
	AskQty string `json:"askQty"`

	// BidPrice represents the highest bid in the order book.
	BidPrice string `json:"bidPrice"`

	// BidQty represents the volume at the current bid price.
	BidQty string `json:"bidQty"`

	// Symbol represents the market queried.
	Symbol string `json:"symbol"`
}

func (c *client) OrderBookTicker(ctx context.Context, symbol string) (
	*OrderBookTicker, error) {
	params := make(url.Values)
	params.Set("symbol", symbol)

	res, err := c.get(ctx, fmt.Sprintf("/ticker/bookTicker?%s", params.Encode()))
	if err != nil {
		return nil, err
	}

	var bookTicker OrderBookTicker
	if err = json.Unmarshal(res, &bookTicker); err != nil {
		return nil, errors.Wrap(err, "failed to parse order book ticker")
	}

	return &bookTicker, nil

}
