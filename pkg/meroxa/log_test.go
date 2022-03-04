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

		w.Write([]byte("[timestamp] log message"))
	}))
	defer server.Close()

	c := testClient(server.Client(), server.URL)

	resp, err := c.GetConnectorLogs(context.Background(), "test")
	if err != nil {
		t.Fatal(err)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if want, got := "[timestamp] log message", string(b); want != got {
		t.Fatalf("mismatched of log message: want=%s got=%s", want, got)
	}
}

func TestGetFunctionLogs(t *testing.T) {
	appName := "app-name"
	funcName := "my-func"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		want := functionsPath(appName, funcName) + "/logs"
		if got := req.URL.Path; want != got {
			t.Fatalf("mismatched of request path: want=%s got=%s", want, got)
		}

		w.Write([]byte("[timestamp] log message"))
	}))
	defer server.Close()

	c := testClient(server.Client(), server.URL)

	resp, err := c.GetFunctionLogs(context.Background(), appName, funcName)
	if err != nil {
		t.Fatal(err)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if want, got := "[timestamp] log message", string(b); want != got {
		t.Fatalf("mismatched of log message: want=%s got=%s", want, got)
	}
}
