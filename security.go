package binance

import (
	"net/http"
	"net/url"
)

// SecurityLevel represents the required authentication a Client is required
// to provide when sending requests.
type SecurityLevel int

const (
	// SecurityLevelNone requires no authentication.
	SecurityLevelNone SecurityLevel = 0

	// SecurityLevelUserStream requires a valid API key.
	SecurityLevelUserStream SecurityLevel = 1

	// SecurityLevelMarketData requires a valid API key.
	SecurityLevelMarketData SecurityLevel = 2

	// SecurityLevelTrade requires a valid API key and a signature.
	SecurityLevelTrade SecurityLevel = 3

	// SecurityLevelUserData requires a valid API key and a signature.
	SecurityLevelUserData SecurityLevel = 4

	securityLevelSentinel SecurityLevel = 5
)

// Valid returns whether `level` is a declared SecurityLevel constant.
func (level SecurityLevel) Valid() bool {
	return level >= SecurityLevelNone && level < securityLevelSentinel
}

// RequiresAuth returns whether a SecurityLevel requires a request to have an
// authentication header present.
func (level SecurityLevel) RequiresAuth() bool {
	return level.Valid() && level >= SecurityLevelUserStream
}

// RequiresSigning returns whether a SecurityLevel requires a request to be
// signed.
func (level SecurityLevel) RequiresSigning() bool {
	return level.Valid() && level >= SecurityLevelTrade
}

func getSecurityLevel(u *url.URL, method string) SecurityLevel {
	return securityGroups[u.Path][method]
}

var securityGroups = map[string]map[string]SecurityLevel{
	"/api/v3/account": {
		http.MethodGet: SecurityLevelUserData,
	},

	"/api/v3/allOrderList": {
		http.MethodGet: SecurityLevelUserData,
	},

	"/api/v3/allOrders": {
		http.MethodGet: SecurityLevelUserData,
	},

	"/api/v3/historicalTrades": {
		http.MethodGet: SecurityLevelMarketData,
	},

	"/api/v3/myTrades": {
		http.MethodGet: SecurityLevelUserData,
	},

	"/api/v3/openOrderList": {
		http.MethodGet: SecurityLevelUserData,
	},

	"/api/v3/openOrders": {
		http.MethodGet: SecurityLevelUserData,
	},

	"/api/v3/order": {
		http.MethodDelete: SecurityLevelTrade,
		http.MethodGet:    SecurityLevelUserData,
		http.MethodPost:   SecurityLevelTrade,
	},

	"/api/v3/order/oco": {
		http.MethodPost:   SecurityLevelTrade,
		http.MethodDelete: SecurityLevelTrade,
	},

	"/api/v3/order/test": {
		http.MethodPost: SecurityLevelTrade,
	},

	"/api/v3/orderList": {
		http.MethodGet: SecurityLevelUserData,
	},
}
