package binance

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/luno/jettison/errors"
	"github.com/luno/jettison/j"
	"github.com/luno/jettison/log"
)

type client struct {
	options *ClientOptions
}

// New returns a Client implementation.
func New(opts ...ClientOption) Client {
	c := client{
		options: &defaultOptions,
	}

	// Apply each of the options to the client.
	for _, o := range opts {
		o(c.options)
	}

	return &c
}

func (c *client) setAuthHeader(r *http.Request) *http.Request {
	r.Header.Set("X-MBX-APIKEY", c.options.apiKey)
	return r
}

func (c *client) signRequest(r *http.Request, body []byte) *http.Request {
	sig := hmac.New(sha256.New, []byte(c.options.secretKey))

	timestamp := time.Now().UnixNano() / 1e6
	values := r.URL.Query()

	values.Set("timestamp", fmt.Sprintf("%d", timestamp))
	fmt.Printf("DEBUG (query): %s\n", values.Encode())
	fmt.Printf("DEBUG (body): %s\n", body)
	sig.Write([]byte(values.Encode()))
	sig.Write(body)

	values.Set("signature", hex.EncodeToString(sig.Sum(nil)))
	r.URL.RawQuery = values.Encode()
	return r
}

func (c *client) call(ctx context.Context, method, path string,
	body []byte) (int, []byte, error) {

	// Add useful data into the context to be included in logs.
	ctx = log.ContextWith(ctx, j.MKV{"method": method, "path": path})

	u, err := url.ParseRequestURI(fmt.Sprintf("%s%s", c.options.baseURL,
		path))
	if err != nil {
		return 0, nil, errors.Wrap(err, "failed to parse uri")
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(),
		bytes.NewBuffer(body))
	if err != nil {
		return 0, nil, errors.Wrap(err, "failed to create request")
	}

	// Set required headers and sign request.
	req.Header.Set("Content-type", "application/x-www-form-urlencoded")

	securityGroup := securityGroups[u.Path][method]
	if securityGroup.RequiresAuth() {
		req = c.setAuthHeader(req)
	}

	if securityGroup.RequiresSigning() {
		req = c.signRequest(req, body)
	}

	reqStart := time.Now()
	res, err := c.options.transport.Do(req)
	if err != nil {
		return 0, nil, errors.Wrap(err, "request failed")
	}
	latency := time.Since(reqStart)

	// Record system metrics.
	httpRequestLatency.WithLabelValues(u.Path).Observe(latency.Seconds())
	httpResponseCodes.WithLabelValues(u.Path, fmt.Sprintf("%d",
		res.StatusCode)).Inc()

	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0, nil, errors.Wrap(err, "failed to read response bytes")
	}

	return res.StatusCode, b, nil
}

func (c *client) get(ctx context.Context, path string) (int, []byte, error) {
	return c.call(ctx, http.MethodGet, path, nil)
}

func (c *client) put(ctx context.Context, path string, body []byte) (int,
	[]byte, error) {
	return c.call(ctx, http.MethodPut, path, body)
}

func (c *client) post(ctx context.Context, path string, body []byte) (int,
	[]byte, error) {
	return c.call(ctx, http.MethodPost, path, body)
}

func (c *client) delete(ctx context.Context, path string, body []byte) (int,
	[]byte, error) {
	return c.call(ctx, http.MethodDelete, path, body)
}

// Ping tests the connectivity to the API.
func (c *client) Ping(ctx context.Context) error {
	code, _, err := c.get(ctx, "/ping")
	if err != nil {
		return err
	}

	if code != http.StatusOK {
		return errors.New("ping: non-2xx status code received")
	}

	return nil
}
