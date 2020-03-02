package binance

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/luno/jettison/errors"
	"github.com/luno/jettison/j"
	"github.com/luno/jettison/log"
)

type client struct {
	options *clientOptions
}

// New returns a concrete Client implementation.
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

func (c *client) call(ctx context.Context, method, path string,
	body io.Reader) (int, []byte, error) {

	ctx = log.ContextWith(ctx, j.MKV{
		method: method,
		path:   path,
	})

	req, err := http.NewRequest(method, fmt.Sprintf("%s%s",
		c.options.baseURL, path), body)
	if err != nil {
		return 0, nil, errors.Wrap(err, "failed to create request")
	}

	reqStart := time.Now()
	res, err := c.options.transport.Do(req)
	if err != nil {
		return 0, nil, errors.Wrap(err, "request failed")
	}
	latency := time.Since(reqStart)

	sanitizedPath := stripQueryParams(path)
	httpRequestLatency.WithLabelValues(sanitizedPath).Observe(latency.Seconds())
	httpResponseCodes.WithLabelValues(sanitizedPath,
		fmt.Sprintf("%d", res.StatusCode)).Inc()

	if c.options.logLevel >= LogLevelInfo {
		log.Info(ctx, "",
			j.MKV{
				"status_code": res.StatusCode,
				"latency":     latency.String(),
			})
	}

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

func (c *client) put(ctx context.Context, path string, body io.Reader) (int,
	[]byte, error) {
	return c.call(ctx, http.MethodPut, path, body)
}

func (c *client) post(ctx context.Context, path string, body io.Reader) (int,
	[]byte, error) {
	return c.call(ctx, http.MethodPost, path, body)
}

// Ping tests the connectivity to the API.
func (c *client) Ping(ctx context.Context) error {
	code, _, err := c.call(ctx, http.MethodGet, "/ping", nil)
	if err != nil {
		return err
	}

	if code != http.StatusOK {
		return errors.New("ping: non-2xx status code received")
	}

	return nil
}
