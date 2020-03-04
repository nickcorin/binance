package binance

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

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
