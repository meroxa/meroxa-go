package meroxa

import (
	"context"
	"io"
	"net/http"
	"net/http/httputil"
	"strings"
)

type dumpTransport struct {
	out       io.Writer
	transport http.RoundTripper
}

func (d *dumpTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	err := d.dumpRequest(req)
	if err != nil {
		return nil, err
	}

	resp, err := d.transport.RoundTrip(req)
	if err != nil {
		return nil, err
	}

	err = d.dumpResponse(resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (d *dumpTransport) dumpRequest(req *http.Request) error {
	cloned := req.Clone(context.Background())

	// Makes sure we don't log out the bearer token by accident when it's not nil
	if !strings.Contains(cloned.Header.Get("Authorization"), "nil") {
		cloned.Header.Set("Authorization", "REDACTED")
	}

	dump, err := httputil.DumpRequestOut(cloned, true)
	if err != nil {
		return err
	}
	_, err = d.out.Write(dump)
	if err != nil {
		return err
	}
	return nil
}

func (d *dumpTransport) dumpResponse(resp *http.Response) error {
	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		return err
	}
	_, err = d.out.Write(dump)
	if err != nil {
		return err
	}
	return nil
}
