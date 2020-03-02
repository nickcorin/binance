package binance

import "net/http"

var defaultOptions = ClientOptions{
	baseURL:   "https://api.binance.com/api/v3",
	logLevel:  LogLevelNone,
	transport: http.DefaultClient,
}

// ClientOptions provides configurable fields for Client.
type ClientOptions struct {
	apiKey    string
	baseURL   string
	logLevel  LogLevel
	transport *http.Client
}

// ClientOption is a function adapter for configuring ClientOptions.
type ClientOption func(*ClientOptions)

// LogLevel configures the extent to which a Client should write output logs.
type LogLevel int

const (
	// LogLevelNone disables logging.
	LogLevelNone = 0

	// LogLevelError only logs errors.
	LogLevelError LogLevel = 1

	// LogLevelInfo logs most activity.
	LogLevelInfo LogLevel = 2

	logLevelSentinel = 3
)

// Valid returns whether `level` is a declared LogLevel constant.
func (level LogLevel) Valid() bool {
	return level >= LogLevelNone && level < logLevelSentinel
}

// WithAPIKey returns a ClientOption to set the API Key a Client uses to
// authenticate requests. Not using this option will cause all requests to be
// unauthenticated.
func WithAPIKey(key string) ClientOption {
	return func(opts *ClientOptions) {
		opts.apiKey = key
	}
}

// WithBaseURL returns a ClientOption to set the prefix a Client uses to prefix
// request. This is useful for testing.
func WithBaseURL(url string) ClientOption {
	return func(opts *ClientOptions) {
		opts.baseURL = url
	}
}

// WithLogLevel returns a ClientOption to set the verbosity of a Client's logs.
// Defaults to `LogLevelNone`.
func WithLogLevel(level LogLevel) ClientOption {
	if !level.Valid() {
		return func(opts *ClientOptions) {}
	}

	return func(opts *ClientOptions) {
		opts.logLevel = level
	}
}

// WithTransport returns a client option to set the underlying HTTP Client used
// for requests. Defaults to the DefaultClient.
func WithTransport(transport *http.Client) ClientOption {
	return func(opts *ClientOptions) {
		opts.transport = transport
	}
}
