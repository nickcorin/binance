package binance

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/schema"
	"github.com/luno/jettison"
	"github.com/luno/jettison/errors"
	"github.com/luno/jettison/j"
	"github.com/luno/jettison/log"
)

type client struct {
	encoder *schema.Encoder
	decoder *schema.Decoder
	options *ClientOptions
}

// NewClient returns a Client implementation.
func NewClient(opts ...ClientOption) Client {
	c := client{
		encoder: schema.NewEncoder(),
		decoder: schema.NewDecoder(),
		options: &defaultOptions,
	}

	// Apply each of the options to the client.
	for _, o := range opts {
		o(c.options)
	}

	return &c
}

// debug is a wrapper for jettison.Info logging, but checks that the client has
// a sufficient log level before writing the log.
func (c *client) debug(ctx context.Context, msg string,
	opts ...jettison.Option) {
	if c.options.logLevel >= LogLevelDebug {
		log.Info(ctx, msg, opts...)
	}
}

// error is a wrapper for jettison.Error, but checks that the client has a
// sufficient log level before writing the log. It returns an error for
// convenience of returning c.error(...), although should only be used in this
// context at the highest level in the error stack to prevent duplicate
// logging.
func (c *client) error(ctx context.Context, err error,
	opts ...jettison.Option) error {
	clientErrorsCounter.WithLabelValues(err.Error()).Inc()
	if c.options.logLevel >= LogLevelError {
		log.Error(ctx, err, opts...)
	}
	return err
}

// HeaderAPIKey defines the request header to set with the client's API key.
const HeaderAPIKey = "X-MBX-APIKEY"

func (c *client) setAuthHeader(r *http.Request) *http.Request {
	r.Header.Set(HeaderAPIKey, c.options.apiKey)
	return r
}

func (c *client) signRequest(r *http.Request, body []byte) *http.Request {
	sig := hmac.New(sha256.New, []byte(c.options.secretKey))

	timestamp := time.Now().UnixNano() / 1e6
	values := r.URL.Query()

	values.Set("timestamp", fmt.Sprintf("%d", timestamp))
	sig.Write([]byte(values.Encode()))
	sig.Write(body)

	values.Set("signature", hex.EncodeToString(sig.Sum(nil)))
	r.URL.RawQuery = values.Encode()
	return r
}

func (c *client) call(ctx context.Context, method, path string,
	body []byte) ([]byte, error) {

	// Add useful data into the context to be included in logs.
	ctx = log.ContextWith(ctx, j.MKV{"method": method, "path": path})
	u, err := url.ParseRequestURI(fmt.Sprintf("%s%s", c.options.baseURL,
		path))
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse uri")
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(),
		bytes.NewBuffer(body))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create request")
	}

	// Set required headers and sign request.
	req.Header.Set("Content-type", "application/x-www-form-urlencoded")
	securityLevel := getSecurityLevel(u, method)

	if securityLevel.RequiresAuth() {
		req = c.setAuthHeader(req)
	}

	if securityLevel.RequiresSigning() {
		req = c.signRequest(req, body)
	}

	reqStart := time.Now()
	res, err := c.options.transport.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute request")
	}
	latency := time.Since(reqStart)

	// Record system metrics.
	httpRequestLatencyHist.WithLabelValues(u.Path).Observe(latency.Seconds())
	httpResponseCodesCounter.WithLabelValues(u.Path, fmt.Sprintf("%d",
		res.StatusCode)).Inc()
	c.debug(ctx, "HTTPS client request", j.KV("status_code", res.StatusCode))

	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read response body")
	}

	// We can assume that 2XX response codes mean that the request was
	// successful.
	if res.StatusCode < http.StatusMultipleChoices {
		return b, nil
	}

	// Return a generic error since we didn't receive any extra information
	// about the error.
	if len(b) == 0 {
		return nil, errors.New("unsuccessful response code received",
			j.KV("response_code", res.StatusCode))
	}

	var apiError Error
	err = json.Unmarshal(b, &apiError)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse response error")
	}

	return nil, apiError
}

func (c *client) get(ctx context.Context, path string) ([]byte, error) {
	return c.call(ctx, http.MethodGet, path, nil)
}

func (c *client) put(ctx context.Context, path string, body []byte) ([]byte,
	error) {
	return c.call(ctx, http.MethodPut, path, body)
}

func (c *client) post(ctx context.Context, path string, body []byte) ([]byte,
	error) {
	return c.call(ctx, http.MethodPost, path, body)
}

func (c *client) delete(ctx context.Context, path string, body []byte) ([]byte,
	error) {
	return c.call(ctx, http.MethodDelete, path, body)
}

// Ping tests the connectivity to the API.
func (c *client) Ping(ctx context.Context) error {
	_, err := c.get(ctx, "/ping")
	if err != nil {
		return c.error(ctx, err)
	}

	return nil
}

// ServerTime returns the current time on the REST API server.
func (c *client) ServerTime(ctx context.Context) (time.Time, error) {
	res, err := c.get(ctx, "/time")
	if err != nil {
		return time.Time{}, c.error(ctx, err)
	}

	timeResponse := struct {
		Milliseconds int64 `json:"serverTime"`
	}{}

	if err = json.Unmarshal(res, &timeResponse); err != nil {
		return time.Time{}, c.error(ctx, errors.Wrap(err,
			"failed to parse server time"))
	}

	return time.Unix(0, timeResponse.Milliseconds*1e6), nil
}
