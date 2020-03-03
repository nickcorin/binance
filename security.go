package binance

import "net/http"

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

var securityGroups = map[string]map[string]SecurityLevel{
	"/account": {
		http.MethodGet: SecurityLevelUserData,
	},

	"/allOrderList": {
		http.MethodGet: SecurityLevelUserData,
	},

	"/allOrders": {
		http.MethodGet: SecurityLevelUserData,
	},

	"/historicalTrades": {
		http.MethodGet: SecurityLevelMarketData,
	},

	"/myTrades": {
		http.MethodGet: SecurityLevelUserData,
	},

	"/openOrderList": {
		http.MethodGet: SecurityLevelUserData,
	},

	"/openOrders": {
		http.MethodGet: SecurityLevelUserData,
	},

	"/order": {
		http.MethodDelete: SecurityLevelTrade,
		http.MethodGet:    SecurityLevelUserData,
		http.MethodPost:   SecurityLevelTrade,
	},

	"/order/oco": {
		http.MethodPost:   SecurityLevelTrade,
		http.MethodDelete: SecurityLevelTrade,
	},

	"order/test": {
		http.MethodPost: SecurityLevelTrade,
	},

	"/orderList": {
		http.MethodGet: SecurityLevelUserData,
	},
}
