package binance

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/luno/jettison/errors"
)

// KlineInterval represents the time interval aggregated per candlestick.
type KlineInterval string

// Enumerated types for KlineInterval.
const (
	OneMinute      KlineInterval = "1m"
	ThreeMinutes   KlineInterval = "3m"
	FiveMinutes    KlineInterval = "5m"
	FifteenMinutes KlineInterval = "15m"
	ThirtyMinutes  KlineInterval = "30m"
	OneHour        KlineInterval = "1h"
	TwoHours       KlineInterval = "2h"
	FourHours      KlineInterval = "4h"
	SixHours       KlineInterval = "6h"
	EightHours     KlineInterval = "8h"
	TwelveHours    KlineInterval = "12h"
	OneDay         KlineInterval = "1d"
	ThreeDays      KlineInterval = "3d"
	OneWeek        KlineInterval = "1w"
	OneMonth       KlineInterval = "1M"
)

// KlinesRequest contains the parameters to query kline / candlestick data.
type KlinesRequest struct {
	// EndTime represents the time to query until.
	//
	// If startTime and endTime is not sent, the most recent klines are
	// returned.
	EndTime int64 `schema:"endTime,omitempty"`

	// Interval represents the time interval to aggregate trades.
	//
	// Required.
	Interval KlineInterval `schema:"interval"`

	// Limit represents the maximum amount of klines to query.
	//
	// Default: 500.
	// Max: 1000.
	Limit int64 `schema:"limit,omitempty"`

	// StartTime represents the time to query from.
	//
	// If startTime and endTime is not sent, the most recent klines are
	// returned.
	StartTime int64 `schema:"startTime,omitempty"`

	// Symbol represents the market to query.
	//
	// Required.
	Symbol string `schema:"symbol"`
}

// Kline contains kline / candlestick data.
type Kline struct {
	Close             string
	CloseTime         int64
	High              string
	OpenTime          int64
	Open              string
	Low               string
	QuoteassertVolume string
	TradeCount        int64
	Volume            string
}

// UnmarshalJSON satisfies the json.Unmarshaler interface for the Kline type.
func (k *Kline) UnmarshalJSON(data []byte) error {
	raw := make([]interface{}, 0)
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	var ok bool
	if ot, ok := raw[0].(float64); ok {
		k.OpenTime = int64(ot)
	} else {
		return fmt.Errorf("failed to assert %T to %T", raw[0], k.OpenTime)
	}

	if k.Open, ok = raw[1].(string); !ok {
		return fmt.Errorf("failed to assert %T to %T", raw[0], k.Open)
	}

	if k.High, ok = raw[2].(string); !ok {
		return fmt.Errorf("failed to assert %T to %T", raw[0], k.High)
	}

	if k.Low, ok = raw[3].(string); !ok {
		return fmt.Errorf("failed to assert %T to %T", raw[0], k.Low)
	}

	if k.Close, ok = raw[4].(string); !ok {
		return fmt.Errorf("failed to assert %T to %T", raw[0], k.Close)
	}

	if k.Volume, ok = raw[5].(string); !ok {
		return fmt.Errorf("failed to assert %T to %T", raw[0], k.Volume)
	}

	if ct, ok := raw[6].(float64); ok {
		k.CloseTime = int64(ct)
	} else {
		return fmt.Errorf("failed to assert %T to %T", raw[0], k.CloseTime)
	}

	if k.QuoteassertVolume, ok = raw[7].(string); !ok {
		return fmt.Errorf("failed to assert %T to %T", raw[0],
			k.QuoteassertVolume)
	}

	if tc, ok := raw[8].(float64); ok {
		k.TradeCount = int64(tc)
	} else {
		return fmt.Errorf("failed to assert %T to %T", raw[0], k.TradeCount)
	}

	return nil
}

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

// Klines queries candlestick data.
func (c *client) Klines(ctx context.Context, r *KlinesRequest) ([]Kline,
	error) {
	params := make(url.Values)
	if err := c.encoder.Encode(r, params); err != nil {
		return nil, errors.Wrap(err, "failed to encode klines request")
	}

	res, err := c.get(ctx, fmt.Sprintf("/klines?%s", params.Encode()))
	if err != nil {
		return nil, err
	}

	var klines []Kline
	if err = json.Unmarshal(res, &klines); err != nil {
		return nil, errors.Wrap(err, "failed to parse klines")
	}

	return klines, nil
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
