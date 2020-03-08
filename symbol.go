package binance

import "strings"

// Symbol represents a trading market. Created by concatenating the quote
// asset and the base asset.
type Symbol string

const (
	ETHBTC   Symbol = "ETHBTC"
	LTCBTC   Symbol = "LTCBTC"
	BNBBTC   Symbol = "BNBBTC"
	NEOBTC   Symbol = "NEOBTC"
	BCCBTC   Symbol = "BCCBTC"
	GASBTC   Symbol = "GASBTC"
	HSDBTC   Symbol = "HSDBTC"
	MCOBTC   Symbol = "MCOBTC"
	WTCBTC   Symbol = "WTCBTC"
	LRCBTC   Symbol = "LRCBTC"
	QTUMBTC  Symbol = "QTUMBTC"
	YOYOBTC  Symbol = "YOYOBTC"
	OMGBTC   Symbol = "OMGBTC"
	ZRXBTC   Symbol = "ZRXBTC"
	STRATBTC Symbol = "STRATBTC"

	QTUMETH  Symbol = "QTUMETH"
	EOSETH   Symbol = "EOSETH"
	SNTETH   Symbol = "SNTETH"
	BNTETH   Symbol = "BNTETH"
	BNBETH   Symbol = "BNBETH"
	OAXETH   Symbol = "OAXETH"
	DNTETH   Symbol = "DNTETH"
	MCOETH   Symbol = "MCOETH"
	ICNETH   Symbol = "ICNETH"
	WTCETH   Symbol = "WTCETH"
	LRCETH   Symbol = "LRCETH"
	OMGETH   Symbol = "OMGETH"
	ZRXETH   Symbol = "ZRXETH"
	STRATETH Symbol = "STRATETH"

	BTCUSDT Symbol = "BTCUSDT"
	ETHUSDT Symbol = "ETHUSDT"
)

// Valid returns whether `s` is a declared Symbol constant.
func (s Symbol) Valid() bool {
	return s.BTCBase() || s.ETHBase() || s.USDTBase()
}

// BTCBase checks if a Symbol has a base asset of BTC.
func (s Symbol) BTCBase() bool {
	return strings.EqualFold(string(s[len(s)-3:len(s)]), "BTC")
}

// ETHBase checks if a Symbol has a base asset of ETH.
func (s Symbol) ETHBase() bool {
	return strings.EqualFold(string(s[len(s)-3:len(s)]), "ETH")
}

// StableBase checks if a Symbol has a base asset of USDT.
func (s Symbol) USDTBase() bool {
	return strings.EqualFold(string(s[len(s)-3:len(s)]), "USDT")
}

// Is compares a Symbol string to a given string and tests equality.
func (s Symbol) Is(s2 string) bool {
	return strings.EqualFold(string(s), s2)
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
