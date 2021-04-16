package meroxa

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const (
	apiURL          = "https://api.meroxa.io/v1"
	jsonContentType = "application/json"
	textContentType = "text/plain"
	clientTimeOut   = 5 * time.Second
)

// encodeFunc encodes v into w
type encodeFunc func(w io.Writer, v interface{}) error

func jsonEncodeFunc(w io.Writer, v interface{}) error {
	return json.NewEncoder(w).Encode(v)
}

func stringEncodeFunc(w io.Writer, v interface{}) error {
	if s, ok := v.(string); ok {
		_, err := w.Write([]byte(s))
		return err
	}

	return fmt.Errorf("body is not a string")
}

func noopEncodeFunc(w io.Writer, v interface{}) error {
	return nil
}

// Client represents the Meroxa API Client
type Client struct {
	baseURL   *url.URL
	userAgent string
	token     string

	httpClient *http.Client
}

// New returns a configured Meroxa API Client
func New(token string, options ...Option) (*Client, error) {
	u, err := url.Parse(apiURL)
	if err != nil {
		return nil, err
	}

	c := &Client{
		baseURL:   u,
		userAgent: "meroxa-go",
		token:     token,
		httpClient: &http.Client{
			Timeout: clientTimeOut,
		},
	}

	for _, opt := range options {
		err := opt(c)
		if err != nil {
			return nil, err
		}
	}

	return c, nil
}

func (c *Client) makeRequest(ctx context.Context, method, path string, body interface{}, params url.Values) (*http.Response, error) {
	return c.makeRequestRaw(ctx, method, path, body, params, jsonEncodeFunc)
}

func (c *Client) makeRequestRaw(ctx context.Context, method, path string, body interface{}, params url.Values, encode encodeFunc) (*http.Response, error) {
	req, err := c.newRequest(ctx, method, path, body, encode)
	if err != nil {
		return nil, err
	}

	// Merge params
	if params != nil {
		q := req.URL.Query()
		for k, v := range params { // v is a []string
			for _, vv := range v {
				q.Add(k, vv)
			}
			req.URL.RawQuery = q.Encode()
		}
	}
	resp, err := c.httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) newRequest(ctx context.Context, method, path string, body interface{}, encode encodeFunc) (*http.Request, error) {
	u, err := c.baseURL.Parse(path)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	if body != nil {
		if err := encode(buf, body); err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	// Set Auth
	bearer := fmt.Sprintf("Bearer %s", c.token)

	// add authorization header to the req
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Content-Type", jsonContentType)
	req.Header.Add("Accept", jsonContentType)
	req.Header.Add("User-Agent", c.userAgent)
	return req, nil
}
