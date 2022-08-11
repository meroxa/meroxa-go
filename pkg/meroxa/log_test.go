package meroxa

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetConnectorLogs(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if want, got := "/v1/connectors/test/logs", req.URL.Path; want != got {
			t.Fatalf("mismatched of request path: want=%s got=%s", want, got)
		}
		if _, err := w.Write([]byte("[timestamp] log message")); err != nil {
			t.Fatalf("expected no error, got %+v", err)
		}
	}))
	defer server.Close()

	c := testClient(server.Client(), server.URL)

	resp, err := c.GetConnectorLogs(context.Background(), "test")
	if err != nil {
		t.Fatal(err)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("expected no error, got %+v", err)
	}
	if want, got := "[timestamp] log message", string(b); want != got {
		t.Fatalf("mismatched of log message: want=%s got=%s", want, got)
	}
}

func TestGetFunctionLogs(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if want, got := "/v1/functions/my-func/logs", req.URL.Path; want != got {
			t.Fatalf("mismatched of request path: want=%s got=%s", want, got)
		}

		if _, err := w.Write([]byte("[timestamp] log message")); err != nil {
			t.Fatalf("expected no error, got %+v", err)
		}
	}))
	defer server.Close()

	c := testClient(server.Client(), server.URL)

	resp, err := c.GetFunctionLogs(context.Background(), "my-func")
	if err != nil {
		t.Fatal(err)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("expected no error, got %+v", err)
	}
	if want, got := "[timestamp] log message", string(b); want != got {
		t.Fatalf("mismatched of log message: want=%s got=%s", want, got)
	}
}

func TestGetBuildLogs(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if want, got := "/v1/builds/my-build/logs", req.URL.Path; want != got {
			t.Fatalf("mismatched of request path: want=%s got=%s", want, got)
		}

		if _, err := w.Write([]byte("[timestamp] log message")); err != nil {
			t.Fatalf("expected no error, got %+v", err)
		}
	}))
	defer server.Close()

	c := testClient(server.Client(), server.URL)

	resp, err := c.GetBuildLogs(context.Background(), "my-build")
	if err != nil {
		t.Fatal(err)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("expected no error, got %+v", err)
	}
	if want, got := "[timestamp] log message", string(b); want != got {
		t.Fatalf("mismatched of log message: want=%s got=%s", want, got)
	}
}
