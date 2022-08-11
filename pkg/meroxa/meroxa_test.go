package meroxa

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestClient_MakeRequest(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, r.Body)
		}))
	defer svr.Close()

	u, err := url.Parse(svr.URL)
	if err != nil {
		t.Errorf("expected no error, got %+v", err)
	}
	c := &client{
		baseURL:    u,
		userAgent:  "meroxa-go",
		httpClient: svr.Client(),
	}

	body := bytes.NewBuffer([]byte("test body"))
	response, err := c.MakeRequest(
		context.Background(),
		http.MethodGet,
		"/api/test",
		body,
		url.Values{"test": []string{"param1", "param2"}},
		http.Header{"Test-Header": []string{"value1", "value2"}})
	if err != nil {
		t.Errorf("expected no error, got %+v", err)
	}
	if response.StatusCode != http.StatusOK {
		t.Errorf("expected http status code %d, got %d", http.StatusOK,
			response.StatusCode)
	}
	if response.Request.Method != http.MethodGet {
		t.Errorf("expected http method %s, got %s", http.MethodGet,
			response.Request.Method)
	}
	if response.Request.URL.Path != "/api/test" {
		t.Errorf("expected http path %s, got %s", "/api/test", response.Request.URL.Path)
	}
	b, err := io.ReadAll(response.Body)
	if err != nil {
		t.Errorf("expected no error, got %+v", err)
	}
	if bytes.Equal(body.Bytes(), b) {
		t.Errorf("expected http body %s, got %s", body, b)
	}
	if response.Request.URL.Query()["test"][0] != "param1" {
		t.Errorf("expected http parameter %s, got %s", "param1",
			response.Request.URL.Query()["test"][0])
	}
	if response.Request.URL.Query()["test"][1] != "param2" {
		t.Errorf("expected http parameter %s, got %s", "param2",
			response.Request.URL.Query()["test"][1])
	}
	if response.Request.Header["Test-Header"][0] != "value1,value2" {
		t.Errorf("expected http header %s, got %s", "value1,value2",
			response.Request.Header["Test-Header"][0])
	}
}
