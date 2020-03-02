package binance

import "context"

// Client represents the interface to build a Client to communicate with
// Binance.
type Client interface {
	Ping(ctx context.Context) error
}
