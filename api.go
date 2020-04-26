package binance

import (
	"context"
	"time"
)

// Client provides the methods relating to Binance's REST API.
type Client interface {
	AccountInfo(context.Context) (*AccountInfo, error)
	CancelOrder(context.Context, *CancelOrderRequest) (*CancelOrderResponse, error)
	NewOrder(context.Context, *NewOrderRequest) (*NewOrderResponse, error)
	NewOrderTest(context.Context, *NewOrderRequest) error
	Ping(context.Context) error
	ServerTime(context.Context) (time.Time, error)
	QueryOrder(context.Context, *QueryOrderRequest) (*QueryOrderResponse, error)
}
