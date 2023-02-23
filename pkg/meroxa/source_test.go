package meroxa

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateSourceV1(t *testing.T) {
	getUrl := "https://s3-get.url"
	putUrl := "https://s3-put.url"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()

		// Return response to satisfy client and test response
		s := generateSource(getUrl, putUrl)
		if err := json.NewEncoder(w).Encode(s); err != nil {
			t.Errorf("expected no error, got %+v", err)
		}
	}))
	// Close the server when test finishes
	defer server.Close()

	c := testClient(testRequester(server.Client(), server.URL))
	resp, err := c.CreateSource(context.Background())
	if err != nil {
		t.Errorf("expected no error, got %+v", err)
	}
	if resp.GetUrl != getUrl {
		t.Errorf("expected getUrl %s, got %s", getUrl, resp.GetUrl)
	}
	if resp.PutUrl != putUrl {
		t.Errorf("expected getUrl %s, got %s", putUrl, resp.PutUrl)
	}
}

func TestCreateSourceV2(t *testing.T) {
	sourceInput := &CreateSourceInputV2{
		Environment: &EntityIdentifier{UUID: uuid.NewString()},
	}

	getUrl := "https://s3-get.url"
	putUrl := "https://s3-put.url"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()

		// Return response to satisfy client and test response
		s := generateSource(getUrl, putUrl)
		if err := json.NewEncoder(w).Encode(s); err != nil {
			t.Errorf("expected no error, got %+v", err)
		}
	}))
	// Close the server when test finishes
	defer server.Close()
	c := testClient(testRequester(server.Client(), server.URL))

	resp, err := c.CreateSourceV2(context.Background(), sourceInput)
	if err != nil {
		t.Errorf("expected no error, got %+v", err)
	}
	if resp.GetUrl != getUrl {
		t.Errorf("expected getUrl %s, got %s", getUrl, resp.GetUrl)
	}
	if resp.PutUrl != putUrl {
		t.Errorf("expected getUrl %s, got %s", putUrl, resp.PutUrl)
	}
}

func generateSource(getUrl, putUrl string) Source {
	if getUrl == "" {
		getUrl = "https://meroxa-get-url.com"
	}
	if putUrl == "" {
		putUrl = "https://meroxa-put-url.com"
	}
	return Source{GetUrl: getUrl, PutUrl: putUrl}
}
