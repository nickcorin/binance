package binance

import (
	"context"
	"time"
)

// Client provides the methods relating to Binance's REST API.
type Client interface {
	Ping(ctx context.Context) error
	ServerTime(ctx context.Context) (time.Time, error)
}
