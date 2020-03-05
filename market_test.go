package binance

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestAccountInfo_OK(t *testing.T) {
	srv, err := createTestServer(t, http.StatusOK)
	require.NoError(t, err)
	defer srv.Close()

	c := NewClient(WithBaseURL(srv.URL))
	info, err := c.AccountInfo(context.Background())
	require.NoError(t, err)
	require.NotNil(t, info)
	require.NotNil(t, info.Balances)
	require.Equal(t, 2, len(info.Balances))
}

func TestOrderBook_OK(t *testing.T) {
	srv, err := createTestServer(t, http.StatusOK)
	require.NoError(t, err)
	defer srv.Close()

	c := NewClient(WithBaseURL(srv.URL))
	book, err := c.OrderBook(context.Background(), ETHBTC, 10)
	require.NoError(t, err)
	require.NotNil(t, book)

	require.Equal(t, 1, len(book.Bids))
	require.Equal(t, 1, len(book.Asks))
}

func TestRecentTrades_OK(t *testing.T) {
	srv, err := createTestServer(t, http.StatusOK)
	require.NoError(t, err)
	defer srv.Close()

	c := NewClient(WithBaseURL(srv.URL))
	trades, err := c.RecentTrades(context.Background(), ETHBTC, 10)
	require.NoError(t, err)
	require.NotNil(t, trades)
	require.Equal(t, 1, len(trades))
}

func TestHistoricalTradesFrom_OK(t *testing.T) {
	srv, err := createTestServer(t, http.StatusOK)
	require.NoError(t, err)
	defer srv.Close()

	c := NewClient(WithBaseURL(srv.URL))
	trades, err := c.HistoricalTradesFrom(context.Background(), ETHBTC, 10,
		int64(1))
	require.NoError(t, err)
	require.NotNil(t, trades)
	require.Equal(t, 1, len(trades))
}

func TestHistoricalTrades_OK(t *testing.T) {
	srv, err := createTestServer(t, http.StatusOK)
	require.NoError(t, err)
	defer srv.Close()

	c := NewClient(WithBaseURL(srv.URL))
	trades, err := c.HistoricalTrades(context.Background(), ETHBTC, 10)
	require.NoError(t, err)
	require.NotNil(t, trades)
	require.Equal(t, 1, len(trades))
}

func TestKlines_OK(t *testing.T) {
	srv, err := createTestServer(t, http.StatusOK)
	require.NoError(t, err)
	defer srv.Close()

	c := NewClient(WithBaseURL(srv.URL))
	klines, err := c.Klines(context.Background(), ETHBTC, Minute, 10)
	require.NoError(t, err)
	require.NotNil(t, klines)
	require.Equal(t, 1, len(klines))
}

func TestKlinesBetween_OK(t *testing.T) {
	srv, err := createTestServer(t, http.StatusOK)
	require.NoError(t, err)
	defer srv.Close()

	c := NewClient(WithBaseURL(srv.URL))
	klines, err := c.KlinesBetween(context.Background(), ETHBTC, Minute,
		time.Now().Add(-1*time.Hour*24), time.Now(), 10)
	require.NoError(t, err)
	require.NotNil(t, klines)
	require.Equal(t, 1, len(klines))
}

func TestAveragePrice_OK(t *testing.T) {
	srv, err := createTestServer(t, http.StatusOK)
	require.NoError(t, err)
	defer srv.Close()

	c := NewClient(WithBaseURL(srv.URL))
	price, err := c.AveragePrice(context.Background(), ETHBTC)
	require.NoError(t, err)
	require.NotNil(t, price)
}

func TestTickerStats_OK(t *testing.T) {
	srv, err := createTestServer(t, http.StatusOK)
	require.NoError(t, err)
	defer srv.Close()

	c := NewClient(WithBaseURL(srv.URL))
	stats, err := c.TickerStats(context.Background(), ETHBTC)
	require.NoError(t, err)
	require.NotNil(t, stats)
}

func TestListTickerStats_OK(t *testing.T) {
	srv, err := createTestServer(t, http.StatusOK)
	require.NoError(t, err)
	defer srv.Close()

	c := NewClient(WithBaseURL(srv.URL))
	stats, err := c.ListTickerStats(context.Background())
	require.NoError(t, err)
	require.NotNil(t, stats)
	require.Equal(t, 1, len(stats))
}

func TestPriceTicker_OK(t *testing.T) {
	srv, err := createTestServer(t, http.StatusOK)
	require.NoError(t, err)
	defer srv.Close()

	c := NewClient(WithBaseURL(srv.URL))
	ticker, err := c.PriceTicker(context.Background(), ETHBTC)
	require.NoError(t, err)
	require.NotNil(t, ticker)
}

func TestListPriceTickers_OK(t *testing.T) {
	srv, err := createTestServer(t, http.StatusOK)
	require.NoError(t, err)
	defer srv.Close()

	c := NewClient(WithBaseURL(srv.URL))
	tickers, err := c.ListPriceTickers(context.Background())
	require.NoError(t, err)
	require.NotNil(t, tickers)
	require.Equal(t, 2, len(tickers))
}

func TestOrderBookTicker_OK(t *testing.T) {
	srv, err := createTestServer(t, http.StatusOK)
	require.NoError(t, err)
	defer srv.Close()

	c := NewClient(WithBaseURL(srv.URL))
	ticker, err := c.OrderBookTicker(context.Background(), ETHBTC)
	require.NoError(t, err)
	require.NotNil(t, ticker)
}

func TestListOrderBookTickers_OK(t *testing.T) {
	srv, err := createTestServer(t, http.StatusOK)
	require.NoError(t, err)
	defer srv.Close()

	c := NewClient(WithBaseURL(srv.URL))
	tickers, err := c.ListOrderBookTickers(context.Background())
	require.NoError(t, err)
	require.NotNil(t, tickers)
	require.Equal(t, 2, len(tickers))
}

func TestAggregateTrades_OK(t *testing.T) {
	srv, err := createTestServer(t, http.StatusOK)
	require.NoError(t, err)
	defer srv.Close()

	c := NewClient(WithBaseURL(srv.URL))
	trades, err := c.AggregateTrades(context.Background(), ETHBTC, 10)
	require.NoError(t, err)
	require.NotNil(t, trades)
	require.Equal(t, 1, len(trades))
}

func TestAggregateTradesAfter_OK(t *testing.T) {
	srv, err := createTestServer(t, http.StatusOK)
	require.NoError(t, err)
	defer srv.Close()

	c := NewClient(WithBaseURL(srv.URL))
	trades, err := c.AggregateTradesAfter(context.Background(), ETHBTC,
		time.Now(), 10)
	require.NoError(t, err)
	require.NotNil(t, trades)
	require.Equal(t, 1, len(trades))
}

func TestAggregateTradesBetween_OK(t *testing.T) {
	srv, err := createTestServer(t, http.StatusOK)
	require.NoError(t, err)
	defer srv.Close()

	c := NewClient(WithBaseURL(srv.URL))
	trades, err := c.AggregateTradesBetween(context.Background(), ETHBTC,
		time.Now().Add(-1*time.Hour*24), time.Now(), 10)
	require.NoError(t, err)
	require.NotNil(t, trades)
	require.Equal(t, 1, len(trades))
}

func TestAggregateTradesFrom_OK(t *testing.T) {
	srv, err := createTestServer(t, http.StatusOK)
	require.NoError(t, err)
	defer srv.Close()

	c := NewClient(WithBaseURL(srv.URL))
	trades, err := c.AggregateTradesFrom(context.Background(), ETHBTC, 1234,
		10)
	require.NoError(t, err)
	require.NotNil(t, trades)
	require.Equal(t, 1, len(trades))
}
