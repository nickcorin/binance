package binance

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSecurityLevels(t *testing.T) {
	tests := []struct {
		url           *url.URL
		method        string
		securityLevel SecurityLevel
	}{
		{
			url: &url.URL{
				Scheme:   "https",
				Host:     "api.binance.com",
				Path:     "/api/v3/recentTrades",
				RawQuery: "symbol=ETHBTC",
			},
			method:        http.MethodGet,
			securityLevel: SecurityLevelNone,
		},
		{
			url: &url.URL{
				Scheme:   "https",
				Host:     "api.binance.com",
				Path:     "/api/v3/order",
				RawQuery: "symbol=ETHBTC",
			},
			method:        http.MethodGet,
			securityLevel: SecurityLevelUserData,
		},
		{
			url: &url.URL{
				Scheme: "https",
				Host:   "api.binance.com",
				Path:   "/api/v3/order",
			},
			method:        http.MethodPost,
			securityLevel: SecurityLevelTrade,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			level := getSecurityLevel(test.url, test.method)
			require.Equal(t, test.securityLevel, level)
		})
	}
}
