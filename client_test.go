package binance

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestServerTime_OK(t *testing.T) {
	srv, err := createTestServer(t, http.StatusOK)
	require.NoError(t, err)
	defer srv.Close()

	c := NewClient(WithBaseURL(srv.URL))
	serverTime, err := c.ServerTime(context.Background())
	require.NoError(t, err)
	require.False(t, serverTime.IsZero())
}
