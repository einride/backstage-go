package catalog

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type clientConfig struct {
	token   string
	baseURL string
}

// ClientOption configures a [Client].
type ClientOption func(*clientConfig)

// WithToken sets the bearer token to use for authentication.
func WithToken(token string) ClientOption {
	return func(config *clientConfig) {
		config.token = token
	}
}

// WithBaseURL sets the backend base URL.
func WithBaseURL(baseURL string) ClientOption {
	return func(config *clientConfig) {
		config.baseURL = baseURL
	}
}

// Client to the Backstage Catalog API.
type Client struct {
	config     clientConfig
	httpClient *http.Client
}

// NewClient creates a new catalog API [Client].
func NewClient(options ...ClientOption) *Client {
	client := &Client{
		httpClient: http.DefaultClient,
	}
	for _, option := range options {
		option(&client.config)
	}
	if client.config.token != "" {
		client.httpClient.Transport = &tokenRoundTripper{
			token: client.config.token,
			next:  http.DefaultTransport,
		}
	}
	return client
}

type tokenRoundTripper struct {
	token string
	next  http.RoundTripper
}

func (t *tokenRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	request.Header.Set("Authorization", "Bearer "+t.token)
	return t.next.RoundTrip(request)
}

func (c *Client) get(
	ctx context.Context,
	path string,
	query url.Values,
	fn func(*http.Response) error,
) error {
	return c.execute(ctx, http.MethodGet, path, query, nil, fn)
}

func (c *Client) execute(
	ctx context.Context,
	method string,
	path string,
	query url.Values,
	body io.Reader,
	fn func(*http.Response) error,
) (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("%s %s: %w", method, path, err)
		}
	}()
	requestURL, err := url.Parse(c.config.baseURL + path)
	if err != nil {
		return err
	}
	if len(query) > 0 {
		requestURL.RawQuery = query.Encode()
	}
	httpRequest, err := http.NewRequestWithContext(ctx, method, requestURL.String(), body)
	if err != nil {
		return err
	}
	httpResponse, err := c.httpClient.Do(httpRequest)
	if err != nil {
		return err
	}
	defer func() {
		_ = httpResponse.Body.Close()
	}()
	if httpResponse.StatusCode != http.StatusOK {
		return newStatusError(httpResponse)
	}
	if fn != nil {
		return fn(httpResponse)
	}
	return nil
}

func (c *Client) delete(
	ctx context.Context,
	path string,
) (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("%s %s: %w", http.MethodDelete, path, err)
		}
	}()
	httpRequest, err := http.NewRequestWithContext(ctx, http.MethodDelete, c.config.baseURL+path, nil)
	if err != nil {
		return err
	}
	httpResponse, err := c.httpClient.Do(httpRequest)
	if err != nil {
		return err
	}
	defer func() {
		_ = httpResponse.Body.Close()
	}()
	if httpResponse.StatusCode != http.StatusNoContent {
		return newStatusError(httpResponse)
	}
	return nil
}
