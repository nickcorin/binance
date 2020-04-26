package binance

import (
	"context"
	"encoding/json"

	"github.com/luno/jettison/errors"
)

// AccountInfo contains all information pertaining to a user's account.
type AccountInfo struct {
	AccountType      string    `json:"accountType"`
	Balances         []Balance `json:"balances"`
	BuyerCommission  int       `json:"buyerCommission"`
	CanDeposit       bool      `json:"canDesposit"`
	CanTrade         bool      `json:"canTrade"`
	CanWithdraw      bool      `json:"canWithdraw"`
	MakerCommission  int       `json:"makerCommission"`
	SellerCommission int       `json:"sellerCommission"`
	TakerCommission  int       `json:"takerCommission"`
	UpdateTime       int64     `json:"updateTime"`
}

// Balance contains a breakdown of a wallet's funds.
type Balance struct {
	Asset  string
	Free   float64
	Locked float64
}

// AccountInfo returns all information and balances for a user account.
func (c *client) AccountInfo(ctx context.Context) (*AccountInfo, error) {
	res, err := c.get(ctx, "/account")
	if err != nil {
		return nil, err
	}

	var info AccountInfo
	if err = json.Unmarshal(res, &info); err != nil {
		return nil, errors.Wrap(err, "failed to parse account info")
	}

	return &info, err
}
