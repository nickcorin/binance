package binance

import (
	"context"
	"time"
)

// Client provides the methods relating to Binance's REST API.
type Client interface {
	AccountInfo(context.Context) (*AccountInfo, error)
	NewOrder(context.Context, *NewOrderRequest) (*NewOrderResponse, error)
	NewOrderTest(context.Context, *NewOrderRequest) error
	ServerTime(context.Context) (time.Time, error)
}
