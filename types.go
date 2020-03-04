package binance

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/luno/jettison/errors"
)

type averagePriceResponse struct {
	Minutes int    `json:"mins"`
	Price   string `json:"price"`
}

// AveragePrice contains the aggregation of price movements over a period of
// time.
type AveragePrice struct {
	Minutes int
	Price   float64
}

func (price *AveragePrice) UnmarshalJSON(data []byte) error {
	var resp averagePriceResponse
	err := json.Unmarshal(data, &resp)
	if err != nil {
		return err
	}

	p, err := strconv.ParseFloat(resp.Price, 64)
	if err != nil {
		return err
	}

	price.Minutes = resp.Minutes
	price.Price = p

	return nil
}

// Interval represents constant durations of time.
type Interval string

const (
	Minute         Interval = "1m"
	ThreeMinutes   Interval = "3m"
	FiveMinutes    Interval = "5m"
	FifteenMinutes Interval = "15m"
	ThirtyMinutes  Interval = "30m"
	Hour           Interval = "1h"
	TwoHours       Interval = "2h"
	FourHours      Interval = "4h"
	SixHours       Interval = "6h"
	EightHours     Interval = "8h"
	TwelveHours    Interval = "12h"
	Day            Interval = "1d"
	ThreeDays      Interval = "3d"
	Week           Interval = "1w"
	Month          Interval = "1M"
)

type klineResponse []interface{}

// Kline contains candlestick data over a period of time.
type Kline struct {
	OpenTime    time.Time
	Open        float64
	High        float64
	Low         float64
	Close       float64
	Volume      float64
	CloseTime   time.Time
	QuoteVolume float64
	TradeCount  int64
}

func (k *Kline) UnmarshalJSON(data []byte) error {
	var resp klineResponse
	err := json.Unmarshal(data, &resp)
	if err != nil {
		return err
	}

	// Note: JSON numbers are treated at floats. Large numbers (outside of
	// the int32 range) will fail to be type asserted to int64 because of
	// scientific notation. A way around this is to type assert into a
	// float64 and cast with int64(...).
	ot, ok := resp[0].(float64)
	if !ok {
		return errors.New(fmt.Sprintf("got data of type %T (%v), wanted float64",
			resp[0], resp[0]))
	}

	oString, ok := resp[1].(string)
	if !ok {
		return errors.New(fmt.Sprintf("got data of type %T (%v), wanted string",
			resp[1], resp[1]))
	}

	o, err := strconv.ParseFloat(oString, 64)
	if err != nil {
		return err
	}

	hString, ok := resp[2].(string)
	if !ok {
		return errors.New(fmt.Sprintf("got data of type %T (%v), wanted string",
			resp[2], resp[2]))
	}

	h, err := strconv.ParseFloat(hString, 64)
	if err != nil {
		return err
	}

	lString, ok := resp[3].(string)
	if !ok {
		return errors.New(fmt.Sprintf("got data of type %T (%v), wanted string",
			resp[3], resp[3]))
	}

	l, err := strconv.ParseFloat(lString, 64)
	if err != nil {
		return err
	}

	cString, ok := resp[4].(string)
	if !ok {
		return errors.New(fmt.Sprintf("got data of type %T (%v), wanted string",
			resp[4], resp[4]))
	}

	c, err := strconv.ParseFloat(cString, 64)
	if err != nil {
		return err
	}

	vString, ok := resp[5].(string)
	if !ok {
		return errors.New(fmt.Sprintf("got data of type %T (%v), wanted string",
			resp[5], resp[5]))
	}

	v, err := strconv.ParseFloat(vString, 64)
	if err != nil {
		return err
	}

	ct, ok := resp[6].(float64)
	if !ok {
		return errors.New(fmt.Sprintf("got data of type %T (%v) for , wanted float64",
			resp[6], resp[6]))
	}

	qvString, ok := resp[7].(string)
	if !ok {
		return errors.New(fmt.Sprintf("got data of type %T (%v), wanted string",
			resp[7], resp[7]))
	}

	qv, err := strconv.ParseFloat(qvString, 64)
	if err != nil {
		return err
	}

	tc, ok := resp[8].(float64)
	if !ok {
		return errors.New(fmt.Sprintf("got data of type %T (%v) for trade count, wanted float64",
			resp[8], resp[8]))
	}

	k.OpenTime = time.Unix(0, int64(ot)*1e6)
	k.Open = o
	k.High = h
	k.Low = l
	k.Close = c
	k.Volume = v
	k.CloseTime = time.Unix(0, int64(ct)*1e6)
	k.QuoteVolume = qv
	k.TradeCount = int64(tc)
	return nil
}

type orderBookResponse struct {
	LastUpdateID int64      `json:"lastUpdateId"`
	Bids         [][]string `json:"bids"`
	Asks         [][]string `json:"asks"`
}

// OrderBook contains a list of open bids and asks on the exchange.
type OrderBook struct {
	LastUpdateID int64
	Bids         []OrderBookEntry
	Asks         []OrderBookEntry
}

func (book *OrderBook) UnmarshalJSON(data []byte) error {
	var resp orderBookResponse
	err := json.Unmarshal(data, &resp)
	if err != nil {
		return err
	}

	book.LastUpdateID = resp.LastUpdateID

	for _, b := range resp.Bids {
		p, err := strconv.ParseFloat(b[0], 64)
		if err != nil {
			return err
		}

		v, err := strconv.ParseFloat(b[1], 64)
		if err != nil {
			return err
		}

		book.Bids = append(book.Bids, OrderBookEntry{p, v})
	}

	for _, a := range resp.Asks {
		p, err := strconv.ParseFloat(a[0], 64)
		if err != nil {
			return err
		}

		v, err := strconv.ParseFloat(a[1], 64)
		if err != nil {
			return err
		}

		book.Asks = append(book.Asks, OrderBookEntry{p, v})
	}

	return nil
}

// OrderBookEntry is an aggregation of open orders at a given price.
type OrderBookEntry struct {
	Price  float64
	Volume float64
}

type orderBookTickerResponse struct {
	Symbol    string `json:"symbol"`
	BidPrice  string `json:"bidPrice"`
	BidVolume string `json:"bidQty"`
	AskPrice  string `json:"askPrice"`
	AskVolume string `json:"askQty"`
}

// OrderBookTicker contains the best bid and ask prices for a Symbol at any
// given time.
type OrderBookTicker struct {
	Symbol    string
	BidPrice  float64
	BidVolume float64
	AskPrice  float64
	AskVolume float64
}

func (ticker *OrderBookTicker) UnmarshalJSON(data []byte) error {
	var resp orderBookTickerResponse
	err := json.Unmarshal(data, &resp)
	if err != nil {
		return err
	}

	bp, err := strconv.ParseFloat(resp.BidPrice, 64)
	if err != nil {
		return err
	}

	bv, err := strconv.ParseFloat(resp.BidVolume, 64)
	if err != nil {
		return err
	}

	ap, err := strconv.ParseFloat(resp.AskPrice, 64)
	if err != nil {
		return err
	}

	av, err := strconv.ParseFloat(resp.AskVolume, 64)
	if err != nil {
		return err
	}

	ticker.Symbol = resp.Symbol
	ticker.BidPrice = bp
	ticker.BidVolume = bv
	ticker.AskPrice = ap
	ticker.AskVolume = av

	return nil
}

type priceTickerResponse struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

// PriceTicker contains a price for a Symbol.
type PriceTicker struct {
	Symbol string
	Price  float64
}

func (ticker *PriceTicker) UnmarshalJSON(data []byte) error {
	var resp priceTickerResponse
	err := json.Unmarshal(data, &resp)
	if err != nil {
		return err
	}

	p, err := strconv.ParseFloat(resp.Price, 64)
	if err != nil {
		return err
	}

	ticker.Symbol = resp.Symbol
	ticker.Price = p

	return nil
}

type ticketStatsResponse struct {
	Symbol             string `json:"symbol"`
	PriceChange        string `json:"priceChange"`
	PriceChangePercent string `json:"priceChangePercent"`
	WeightedAvgPrice   string `json:"weightedAvgPrice"`
	PrevClosePrice     string `json:"prevClosePrice"`
	LastPrice          string `json:"lastPrice"`
	LastVolume         string `json:"lastQty"`
	BidPrice           string `json:"bidPrice"`
	AskPrice           string `json:"askPrice"`
	OpenPrice          string `json:"openPrice"`
	HighPrice          string `json:"highPrice"`
	LowPrice           string `json:"lowPrice"`
	Volume             string `json:"volume"`
	QuoteVolume        string `json:"quoteVolume"`
	OpenTime           int64  `json:"openTime"`
	CloseTime          int64  `json:"closeTime"`
	FirstID            int64  `json:"firstID"`
	LastID             int64  `json:"lastID"`
	TradeCount         int64  `json:"count"`
}

// TickerStats contains price change statistics over a 24 hour rolling window.
type TickerStats struct {
	Symbol             string
	PriceChange        float64
	PriceChangePercent float64
	WeightedAvgPrice   float64
	PrevClosePrice     float64
	LastPrice          float64
	LastVolume         float64
	BidPrice           float64
	AskPrice           float64
	OpenPrice          float64
	HighPrice          float64
	LowPrice           float64
	Volume             float64
	QuoteVolume        float64
	OpenTime           time.Time
	CloseTime          time.Time
	FirstID            int64
	LastID             int64
	TradeCount         int64
}

func (stats *TickerStats) UnmarshalJSON(data []byte) error {
	var resp ticketStatsResponse
	err := json.Unmarshal(data, &resp)
	if err != nil {
		return err
	}

	pc, err := strconv.ParseFloat(resp.PriceChange, 64)
	if err != nil {
		return err
	}

	pcp, err := strconv.ParseFloat(resp.PriceChangePercent, 64)
	if err != nil {
		return err
	}

	wap, err := strconv.ParseFloat(resp.WeightedAvgPrice, 64)
	if err != nil {
		return err
	}

	prevCp, err := strconv.ParseFloat(resp.PrevClosePrice, 64)
	if err != nil {
		return err
	}

	lp, err := strconv.ParseFloat(resp.LastPrice, 64)
	if err != nil {
		return err
	}

	lv, err := strconv.ParseFloat(resp.LastVolume, 64)
	if err != nil {
		return err
	}

	bp, err := strconv.ParseFloat(resp.BidPrice, 64)
	if err != nil {
		return err
	}

	ap, err := strconv.ParseFloat(resp.AskPrice, 64)
	if err != nil {
		return err
	}

	op, err := strconv.ParseFloat(resp.OpenPrice, 64)
	if err != nil {
		return err
	}

	hp, err := strconv.ParseFloat(resp.HighPrice, 64)
	if err != nil {
		return err
	}

	low, err := strconv.ParseFloat(resp.LowPrice, 64)
	if err != nil {
		return err
	}

	v, err := strconv.ParseFloat(resp.Volume, 64)
	if err != nil {
		return err
	}

	qv, err := strconv.ParseFloat(resp.QuoteVolume, 64)
	if err != nil {
		return err
	}

	stats.Symbol = resp.Symbol
	stats.PriceChange = pc
	stats.PriceChangePercent = pcp
	stats.WeightedAvgPrice = wap
	stats.PrevClosePrice = prevCp
	stats.LastPrice = lp
	stats.LastVolume = lv
	stats.BidPrice = bp
	stats.AskPrice = ap
	stats.OpenPrice = op
	stats.HighPrice = hp
	stats.LowPrice = low
	stats.Volume = v
	stats.QuoteVolume = qv
	stats.OpenTime = time.Unix(0, resp.OpenTime*1e6)
	stats.CloseTime = time.Unix(0, resp.CloseTime*1e6)
	stats.FirstID = resp.FirstID
	stats.LastID = resp.LastID
	stats.TradeCount = resp.TradeCount

	return nil
}

type tradeResponse struct {
	ID           int64  `json:"id"`
	Price        string `json:"price"`
	Volume       string `json:"qty"`
	QuoteVolume  string `json:"quoteQty"`
	Timestamp    int64  `json:"time"`
	IsBuyerMaker bool   `json:"isBuyerMaker"`
	IsBestMatch  bool   `json:"isBestMatch"`
}

// Trade represents an order that has been successfully executed.
type Trade struct {
	ID           int64
	Price        float64
	Volume       float64
	QuoteVolume  float64
	Timestamp    time.Time
	IsBuyerMaker bool
	IsBestMatch  bool
}

func (t *Trade) UnmarshalJSON(data []byte) error {
	var resp tradeResponse
	err := json.Unmarshal(data, &resp)
	if err != nil {
		return err
	}

	p, err := strconv.ParseFloat(resp.Price, 64)
	if err != nil {
		return err
	}

	v, err := strconv.ParseFloat(resp.Volume, 64)
	if err != nil {
		return err
	}

	qv, err := strconv.ParseFloat(resp.QuoteVolume, 64)
	if err != nil {
		return err
	}

	t.ID = resp.ID
	t.Price = p
	t.Volume = v
	t.QuoteVolume = qv
	t.Timestamp = time.Unix(0, resp.Timestamp*1e6)
	t.IsBuyerMaker = resp.IsBuyerMaker
	t.IsBestMatch = resp.IsBestMatch

	return nil
}
