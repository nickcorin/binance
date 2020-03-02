package binance

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateSignature(t *testing.T) {
	tests := []struct {
		key       string
		secret    string
		params    string
		body      []byte
		signature string
	}{
		{
			key:       "vmPUZE6mv9SD5VNHk4HlWFsOr6aKE2zvsw0MuIgwCIPy6utIco14y7Ju91duEh8A",
			secret:    "NhqPtmdSJYdKjVHjA7PZj4Mge3R5YNiP1e3UZjInClVN65XAbvqqM6A7H5fATj0j",
			body:      []byte("symbol=LTCBTC&side=BUY&type=LIMIT&timeInForce=GTC&quantity=1&price=0.1&recvWindow=5000&timestamp=1499827319559"),
			signature: "c8db56825ae71d6d79447849e617115f4a920fa2acdcab2b053c4b2838bd6b71",
		},
		{
			key:       "vmPUZE6mv9SD5VNHk4HlWFsOr6aKE2zvsw0MuIgwCIPy6utIco14y7Ju91duEh8A",
			secret:    "NhqPtmdSJYdKjVHjA7PZj4Mge3R5YNiP1e3UZjInClVN65XAbvqqM6A7H5fATj0j",
			params:    "symbol=LTCBTC&side=BUY&type=LIMIT&timeInForce=GTC&quantity=1&price=0.1&recvWindow=5000&timestamp=1499827319559",
			signature: "c8db56825ae71d6d79447849e617115f4a920fa2acdcab2b053c4b2838bd6b71",
		},
		{
			key:       "vmPUZE6mv9SD5VNHk4HlWFsOr6aKE2zvsw0MuIgwCIPy6utIco14y7Ju91duEh8A",
			secret:    "NhqPtmdSJYdKjVHjA7PZj4Mge3R5YNiP1e3UZjInClVN65XAbvqqM6A7H5fATj0j",
			params:    "symbol=LTCBTC&side=BUY&type=LIMIT&timeInForce=GTC",
			body:      []byte("quantity=1&price=0.1&recvWindow=5000&timestamp=1499827319559"),
			signature: "0fd168b8ddb4876a0358a8d14d0c9f3da0e9b20c5d52b2a00fcf7d1c602f9a77",
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *tesing.T) {
			sig := generateSignature(key, secret, path, body)
			require.Equal(t, test.signature, sig)
		})
	}
}
