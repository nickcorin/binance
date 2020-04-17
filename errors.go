package binance

// Error defines the structured error that is returned in some reponses.
//
// More information regarding errors can be found in the error codes
// documentation:
//
// https://github.com/binance-exchange/binance-official-api-docs/blob/master/errors.md
type Error struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"msg"`
}

// Error returns the error as a string. Satisfied the stdlib error interface.
func (e Error) Error() string {
	return e.Message
}

// Is allows easy comparisons to be done between an Error and ErrorCode.
func (e Error) Is(code ErrorCode) bool {
	return e.Code == code
}

// IsAny returns whether `e` is contained within a list of ErrorCode types.
func (e Error) IsAny(codes ...ErrorCode) bool {
	for _, code := range codes {
		if e.Is(code) {
			return true
		}
	}
	return false
}

// ErrorCode defines a more granular error type that can be returned from the
// API.
type ErrorCode int

var (
	ErrUnknown                        ErrorCode = -1000
	ErrDisconnected                   ErrorCode = -1001
	ErrUnauthorized                   ErrorCode = -1002
	ErrTooManyRequests                ErrorCode = -1003
	ErrUnexpectedResponse             ErrorCode = -1006
	ErrTimeout                        ErrorCode = -1007
	ErrUnknownOrderComposition        ErrorCode = -1014
	ErrTooManyOrders                  ErrorCode = -1015
	ErrServiceShuttingDown            ErrorCode = -1016
	ErrUnsupportedOperation           ErrorCode = -1020
	ErrInvalidTimestamp               ErrorCode = -1021
	ErrInvalidSignature               ErrorCode = -1022
	ErrIllegalChars                   ErrorCode = -1100
	ErrTooManyParams                  ErrorCode = -1101
	ErrMandatoryParamEmptyOrMalformed ErrorCode = -1102
	ErrUnknownParam                   ErrorCode = -1103
	ErrUnreadParams                   ErrorCode = -1104
	ErrParamEmpty                     ErrorCode = -1105
	ErrParamNotRequired               ErrorCode = -1106
	ErrBadPrecision                   ErrorCode = -1111
	ErrNoDepth                        ErrorCode = -1112
	ErrTIFNotRequired                 ErrorCode = -1114
	ErrInvalidTIF                     ErrorCode = -1115
	ErrInvalidOrderType               ErrorCode = -1116
	ErrInvalidSide                    ErrorCode = -1117
	ErrEmptyNewClientOrderID          ErrorCode = -1118
	ErrEmptyOriginalClientOrderID     ErrorCode = -1119
	ErrBadInterval                    ErrorCode = -1120
	ErrBadSymbol                      ErrorCode = -1121
	ErrInvalidListenKey               ErrorCode = -1125
	ErrMoreThanXHours                 ErrorCode = -1127
	ErrOptionalParamsBadCombo         ErrorCode = -1128
	ErrInvalidParam                   ErrorCode = -1130
	ErrNewOrderRejected               ErrorCode = -2010
	ErrCancelRejected                 ErrorCode = -2011
	ErrNoSuchOrder                    ErrorCode = -2013
	ErrAPIKeyFormat                   ErrorCode = -2014
	ErrRejectedMBXKey                 ErrorCode = -2015
	ErrNoTradingWindow                ErrorCode = -2016
)
