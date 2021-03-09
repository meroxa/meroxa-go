package meroxa

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"time"
)

type DumpTransport struct {
	r http.RoundTripper
}

func (d *DumpTransport) RoundTrip(h *http.Request) (*http.Response, error) {
	dump, _ := httputil.DumpRequestOut(h, true)
	fmt.Println(string(dump))
	resp, err := d.r.RoundTrip(h)
	dump, _ = httputil.DumpResponse(resp, true)
	fmt.Println(string(dump))
	return resp, err
}

func httpDebugClient() *http.Client {
	return &http.Client{
		Transport: &DumpTransport{http.DefaultTransport},
		Timeout:   5 * time.Second,
	}
}
