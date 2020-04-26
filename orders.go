package binance

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/luno/jettison/errors"
)

// OrderResponseType defines the type of response you'd like to receive after
// creating a new order.
type OrderResponseType string

// Enumerated types for OrderResponseType.
const (
	OrderResponseTypeAck    = "ACK"
	OrderResponseTypeResult = "RESULT"
	OrderResponseTypeFull   = "FULL"
)

// OrderSide is an enumerated string type representing whether an order is a
// buy or sell.
type OrderSide string

// Enumerated types for OrderSide.
const (
	Buy  OrderSide = "BUY"
	Sell OrderSide = "SELL"
)

// OrderStatus describes the current state of an order.
type OrderStatus string

const (
	// OrderStatusNew indicates a newly created order.
	OrderStatusNew OrderStatus = "NEW"

	// OrderStatusPartiallyFilled indicates an order which has had some of
	// its bought or sold.
	OrderStatusPartiallyFilled OrderStatus = "PARTIALLY_FILLED"

	// OrderStatusFilled indicates a completed order.
	OrderStatusFilled OrderStatus = "FILLED"

	// OrderStatusCancelled indicates an order that has been cancelled.
	OrderStatusCancelled OrderStatus = "CANCELLED"

	// OrderStatusPendingCancel is currently unused.
	OrderStatusPendingCancel OrderStatus = "PENDING_CANCEL"

	// OrderStatusRejected indicates an order that has been rejected by
	// the exchange. This could be due to insuffienct funds in an account
	// or invalid parameters.
	OrderStatusRejected OrderStatus = "REJECTED"

	// OrderStatusExpired indicates an order that has outlived its
	// TimeInForce configuration.
	OrderStatusExpired OrderStatus = "EXPIRED"
)

// OrderType describes the behavior of an order's execution.
type OrderType string

const (
	// OrderTypeLimit is a limit order which has a maximum or minimum price
	// to buy or sell.
	OrderTypeLimit OrderType = "LIMIT"

	// OrderTypeMarket is a market order which only specifies a quantity to
	// buy or sell at the current market price.
	OrderTypeMarket = "MARKET"

	// OrderTypeStopLoss is a market order that only executes when a given
	// stop price is reached. Usually used to minimize loss when the market
	// drops.
	OrderTypeStopLoss OrderType = "STOP_LOSS"

	// OrderTypeStopLossLimit is a limit order that only executes when a
	// given stop price is reached.
	OrderTypeStopLossLimit OrderType = "STOP_LOSS_LIMIT"

	// OrderTypeTakeProfit is a market order that only executes when a given
	// stop price is reached. Usually used to lock in profits when the
	// market suddenly rises.
	OrderTypeTakeProfit OrderType = "TAKE_PROFIT"

	// OrderTypeTakeProfitLimit is a limit order that only executed when a
	// given stop price is reached. Usually used to lock in profits when
	// the market suddenly rises.
	OrderTypeTakeProfitLimit OrderType = "TAKE_PROFIT_LIMIT"

	// OrderTypeLimitMaker is a limit order that is rejected if it would
	// get executed immediately and trade as a taker.
	OrderTypeLimitMaker OrderType = "LIMIT_MAKER"
)

// NewOrderRequest contains all request parameters for creating a new order.
type NewOrderRequest struct {
	// Symbol represents the market to place the order on.
	//
	// Required for all order types.
	Symbol string `schema:"symbol"`

	// Side represents whether this order is a buy or sell.
	//
	// Required for all order types.
	Side OrderSide `schema:"side"`

	// Type represents what kind of order to place.
	//
	// Required for all order types.
	Type OrderType `schema:"type"`

	// TimeInForce represents the duration of validity of the order.
	//
	// Required for orders of type LIMIT, STOP_LOSS_LIMIT and TAKE_PROFIT_LIMIT.
	TimeInForce TimeInForce `schema:"timeInForce,omitempty"`

	// Qty represents the quantity to buy or sell.
	//
	// Required for orders of type MARKET, STOP_LOSS, TAKE_PROFIT and
	// LIMIT_MAKER.
	Qty float64 `schema:"quantity,omitempty"`

	// Required for order of type MARKET if Qty is not set.
	QuoteOrderQty float64 `schema:"quoteOrderQty,omitempty"`

	// Price represents the price at which to place the order.
	//
	// Required for orders of type LIMIT, STOP_LOSS_LIMIT and TAKE_PROFIT_LIMIT.
	Price float64 `schema:"price,omitempty"`

	// NewClientOrderID represents a unique identifier for the order, supplied
	// by the client.
	//
	// Optional.
	// Default is a randomly generated string.
	NewClientOrderID string `schema:"newClientOrderId,omitempty"`

	// StopPrice represents the price the market needs to reach before placing
	// the order as a market order.
	//
	// Required for orders of type STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT and
	// TAKE_PROFIT_LIMIT.
	StopPrice float64 `schema:"stopPrice,omitempty"`

	// IcebergQty represents the maximum amount per sub-order until the total
	// quantity of the order has been filled. Orders with type LIMIT or
	// LIMIT_MAKER are automatically made an iceberg order if an IcebergQty is
	// sent. Any order with IcebergQty set, MUST have it's TimeInForce set to
	// GTC.
	//
	// Optional.
	IcebergQty float64 `schema:"icebergQty,omitempty"`

	// ResponseType represents the kind of response you want to receive back.
	//
	// Optional.
	// Default: ACK for orders of type LIMIT or MARKET, FULL otherwise.
	ResponseType OrderResponseType `schema:"newOrderRespType,omitempty"`

	// ReceiveWindow represents the duration of validity in ms of the request.
	//
	// Optional.
	// Default: 5000ms. Maximum: 60000ms.
	ReceiveWindow int64 `schema:"recvWindow,omitempty"`
}

// NewOrderResponse contains
type NewOrderResponse struct {
	// Symbol represents the market the order was placed on.
	Symbol string `schema:"symbol"`

	// OrderID represents a unique identifier for the order, generated by
	// Binance.
	OrderID int64 `schema:"orderId"`

	// OrderListID will always be -1 if the order was not an OCO order.
	OrderListID int64 `schema:"orderListId"`

	// ClientOrderID represents the unique identifier for the order, sent by
	// the client on creation. If NewClientOrderID was empty in the request,
	// this will be a randomly generated string.
	ClientOrderID string `schema:"clientOrderId"`

	// TransactTime represents the unix timestamp in milliseconds of the time
	// the order was placed.
	//
	// TODO: Verify this comment.
	TransactTime int64 `schema:"transactTime"`

	// Price represents the price at which the order was placed.
	//
	// Returned with response types RESULT and FULL.
	Price string `schema:"price,omitempty"`

	// OriginalQty represents the original quantity the order was placed for.
	//
	// Returned with response types RESULT and FULL.
	OriginalQty string `schema:"origQty,omitempty"`

	// ExecutedQty represents how much of the original quantity has been
	// executed.
	//
	// Returned with response types RESULT and FULL.
	ExecutedQty         string `schema:"executedQty,omitempty"`
	CummulativeQuoteQty string `schema:"cummulativeQuoteQty,omitempty"`

	// Status represents the current status of the order.
	//
	// Returned with response types RESULT and FULL.
	Status OrderStatus `schema:"status,omitempty"`

	// TimeInForce represents the duration of validity of the order.
	//
	// Returned with response types RESULT and FULL.
	TimeInForce TimeInForce `schema:"timeInForce,omitempty"`

	// Type represents the type of the order.
	//
	// Returned with response types RESULT and FULL.
	Type OrderType `schema:"type,omitempty"`

	// Side represents whether the order was a buy or sell.
	//
	// Returned with response types RESULT and FULL.
	Side OrderSide `schema:"side,omitempty"`

	// Fills contains sub-orders executed in order to fully execute an order.
	//
	// Returned wtih response type FULL.
	Fills []OrderFill `schema:"fills,omitempty"`
}

// OrderFill represents a sub-order executed as part of a larger order.
type OrderFill struct {
	// Price represents the price at which a sub-order was executed.
	Price string `schema:"price"`

	// Qty represents the quantity of the sub-order.
	Qty string `schema:"qty"`

	// Commission represents the amount of commission earned by a sub-order.
	Commission string `schema:"commission"`

	// CommissionAsset represents the asset that commission is paid out in.
	CommissionAsset string `schema:"commissionAsset"`
}

// NewOrder places a new order on the exchange.
func (c *client) NewOrder(ctx context.Context, r *NewOrderRequest) (
	*NewOrderResponse, error) {
	params := make(url.Values)
	err := c.encoder.Encode(r, params)
	if err != nil {
		return nil, errors.Wrap(err, "failed to encode new order request")
	}

	res, err := c.post(ctx, "/order", []byte(params.Encode()))
	if err != nil {
		return nil, err
	}

	var orderResponse NewOrderResponse
	if err = json.Unmarshal(res, &orderResponse); err != nil {
		return nil, errors.Wrap(err, "failed to parse new order response")
	}

	return &orderResponse, nil
}

// NewOrderTest creates and validates a new order on the exchange, but does
// not send it to the matching engine.
func (c *client) NewOrderTest(ctx context.Context, r *NewOrderRequest) error {
	params := make(url.Values)
	err := c.encoder.Encode(r, params)
	if err != nil {
		return errors.Wrap(err, "failed to encode new order request")
	}

	_, err = c.post(ctx, "/order/test", []byte(params.Encode()))
	if err != nil {
		return err
	}

	return nil
}
