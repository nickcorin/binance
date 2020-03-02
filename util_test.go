package binance

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStripQueryParams(t *testing.T) {

	tests := []struct {
		Path          string
		SanitizedPath string
	}{
		{
			"https://api.binance.com/api/v3/depth?symbol=BTC",
			"https://api.binance.com/api/v3/depth",
		},
		{
			"https://api.binance.com/api/v3/depth?symbol=BTC&limit=10",
			"https://api.binance.com/api/v3/depth",
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			output := stripQueryParams(test.Path)
			require.Equal(t, test.SanitizedPath, output)
		})
	}
}
