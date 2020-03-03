package binance

import (
	"context"
	"net/http"
	"testing"

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
