package binance

import "strings"

//go:generate stringer -type=Symbol

// Symbol represents a trading market. Created by concatenating the quote
// asset and the base asset.
type Symbol int

const (
	SymbolUnknown Symbol = 0

	// BTC markets.
	ETHBTC      Symbol = 1
	LTCBTC      Symbol = 2
	BNBBTC      Symbol = 3
	NEOBTC      Symbol = 4
	BCCBTC      Symbol = 5
	GASBTC      Symbol = 6
	HSDBTC      Symbol = 7
	MCOBTC      Symbol = 8
	WTCBTC      Symbol = 9
	LRCBTC      Symbol = 10
	QTUMBTC     Symbol = 11
	YOYOBTC     Symbol = 12
	OMGBTC      Symbol = 13
	ZRXBTC      Symbol = 14
	STRATBTC    Symbol = 15
	btcSentinel Symbol = 16

	// ETH markets.
	QTUMETH     Symbol = 101
	EOSETH      Symbol = 102
	SNTETH      Symbol = 103
	BNTETH      Symbol = 104
	BNBETH      Symbol = 105
	OAXETH      Symbol = 106
	DNTETH      Symbol = 107
	MCOETH      Symbol = 108
	ICNETH      Symbol = 109
	WTCETH      Symbol = 110
	LRCETH      Symbol = 111
	OMGETH      Symbol = 112
	ZRXETH      Symbol = 113
	STRATETH    Symbol = 114
	ethSentinel Symbol = 115

	// Stablecoin markets.

	BTCUSDT        Symbol = 201
	ETHUSDT        Symbol = 202
	stableSentinel        = 203

	symbolSentinel Symbol = 204
)

// Valid returns whether `s` is a declared Symbol constant.
func (s Symbol) Valid() bool {
	return s.BTCBase() || s.ETHBase() || s.StableBase()
}

// BTCBase checks if a Symbol has a base asset of BTC.
func (s Symbol) BTCBase() bool {
	return s > SymbolUnknown && s < btcSentinel
}

// ETHBase checks if a Symbol has a base asset of ETH.
func (s Symbol) ETHBase() bool {
	return s > btcSentinel && s < ethSentinel
}

// StableBase checks if a Symbol has a base asset of some stable coin.
func (s Symbol) StableBase() bool {
	return s > ethSentinel && s < symbolSentinel
}

// Is compares a Symbol string to a given string and tests equality.
func (s Symbol) Is(s2 string) bool {
	return strings.EqualFold(s.String(), s2)
}

// IsAny compares a Symbol string to a list of strings and returns if it is
// contained.
func (s Symbol) IsAny(sl ...string) bool {
	for _, s2 := range sl {
		if s.Is(s2) {
			return true
		}
	}
	return false
}
