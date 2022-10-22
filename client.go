package postcodesio

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

const (
	baseURL        = "https://api.postcodes.io"
	defaultTimeout = 30 * time.Second
)

// Client is the base struct for the postcode.io API client.
type Client struct {
	baseURL    string
	httpClient *http.Client
}

// ClientOption describes the type for functional options used when creating a Client.
type ClientOption func(*Client)

// WithTransport is the option to set a custom transport layer to the Client.
func WithTransport(rt http.RoundTripper) ClientOption {
	return func(c *Client) {
		c.httpClient.Transport = rt
	}
}

// WithTimeout is the option to set a timeout to Client's requests.
func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *Client) {
		c.httpClient.Timeout = timeout
	}
}

// New creates a new Client.
func New(opts ...ClientOption) *Client {
	c := &Client{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: defaultTimeout,
		},
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

// get executes a http get request.
func (c *Client) get(ctx context.Context, url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, http.NoBody)
	if err != nil {
		return nil, err
	}

	return c.doRequest(req)
}

// post executes a http post request.
func (c *Client) post(ctx context.Context, url string, body interface{}) ([]byte, error) {
	payload, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	return c.doRequest(req)
}

// doRequest encapsulates an http request-response.
func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return b, nil
}
