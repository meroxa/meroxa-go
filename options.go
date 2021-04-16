package meroxa

import (
	"io"
	"net/url"
	"time"
)

type Option func(*Client) error

func WithBaseURL(rawurl string) Option {
	return func(client *Client) error {
		u, err := url.Parse(rawurl)
		if err != nil {
			return err
		}
		client.baseURL = u
		return nil
	}
}

func WithClientTimeout(timeout time.Duration) Option {
	return func(client *Client) error {
		client.httpClient.Timeout = timeout
		return nil
	}
}

func WithUserAgent(ua string) Option {
	return func(client *Client) error {
		client.userAgent = ua
		return nil
	}
}

func WithDebugOutput(writer io.Writer) Option {
	return func(client *Client) error {
		client.httpClient.Transport = &dumpTransport{
			out:       writer,
			transport: client.httpClient.Transport,
		}
		return nil
	}
}
