package binance

import "context"

// Client provides the methods relating to Binance's REST API.
type Client interface {
	Ping(ctx context.Context) error
}
