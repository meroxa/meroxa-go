package meroxa

import (
	"context"
	"io"
	"net/http"
	"net/http/httputil"
	"strings"
	"unicode/utf8"
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
	if auth := cloned.Header.Get("Authorization"); auth != "" {
		tokens := strings.SplitN(auth, " ", 2)
		if len(tokens) == 2 {
			tokens[1] = d.obfuscate(tokens[1])
			auth = strings.Join(tokens, " ")
		}
		cloned.Header.Set("Authorization", auth)
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

func (d *dumpTransport) obfuscate(text string) string {
	const (
		visibleSuffixLen = 4
		minStarsLen      = 4
		star             = "*"
	)

	if len(text) < minStarsLen {
		// hide whole text
		return strings.Repeat(star, len(text))
	} else if len(text) < minStarsLen+visibleSuffixLen {
		// hide only minStarsLen
		return strings.Repeat(star, minStarsLen) + text[minStarsLen:]
	} else {
		// hide everything except visibleSuffixLen
		starsLen := utf8.RuneCountInString(text) - visibleSuffixLen
		return strings.Repeat(star, starsLen) + text[len(text)-visibleSuffixLen:]
	}
}
