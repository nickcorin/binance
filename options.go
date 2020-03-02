package binance

import "net/http"

var defaultOptions = clientOptions{
	baseURL:   "https://api.binance.com/api/v3",
	logLevel:  LogLevelNone,
	transport: http.DefaultClient,
}

type clientOptions struct {
	apiKey    string
	baseURL   string
	logLevel  LogLevel
	transport *http.Client
}

// ClientOption creates a way for the client to be configured.
type ClientOption func(*clientOptions)

// LogLevel defines the extend to which the client should write output logs
// to stdout.
type LogLevel int

const (
	// LogLevelNone disables logging.
	LogLevelNone = 0

	// LogLevelError only logs errors.
	LogLevelError LogLevel = 1

	// LogLevelInfo logs most client activity.
	LogLevelInfo LogLevel = 2

	logLevelSentinel = 3
)

// Valid returns whether "level" is a declared LogLevel constant.
func (level LogLevel) Valid() bool {
	return level >= LogLevelNone && level < logLevelSentinel
}

// WithAPIKey sets the API key in the client to be used for authenticated
// calls. If this is not set, the request will not be authenticated.
func WithAPIKey(key string) ClientOption {
	return func(opts *clientOptions) {
		opts.apiKey = key
	}
}

// WithBaseURL sets the URL that will prefix all endpoints paths being
// requeted. This is useful for running client tests.
func WithBaseURL(url string) ClientOption {
	return func(opts *clientOptions) {
		opts.baseURL = url
	}
}

// WithLogLevel sets the amount of logging the client provides. It defaults to
// LogLevelNone.
func WithLogLevel(level LogLevel) ClientOption {
	if !level.Valid() {
		return func(opts *clientOptions) {}
	}

	return func(opts *clientOptions) {
		opts.logLevel = level
	}
}

// WithTransport sets the underlying http client used during HTTP request. If
// this is not set, http.DefaultClient is used.
func WithTransport(transport *http.Client) ClientOption {
	return func(opts *clientOptions) {
		opts.transport = transport
	}
}
