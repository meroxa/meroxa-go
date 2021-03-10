package meroxa

import (
	"log"
	"net/http"
	"net/http/httputil"
)

type DumpTransport struct {
	r http.RoundTripper
}

func (d *DumpTransport) RoundTrip(h *http.Request) (*http.Response, error) {
	dump, _ := httputil.DumpRequestOut(h, true)
	log.Printf(string(dump))

	resp, err := d.r.RoundTrip(h)
	dump, _ = httputil.DumpResponse(resp, true)
	log.Printf(string(dump))

	return resp, err
}
