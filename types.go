package binance

import (
	"encoding/json"
	"strconv"
	"time"
)

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
