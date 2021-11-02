package meroxa

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestCreateEndpoint(t *testing.T) {
	var (
		name     = "test"
		protocol = "http"
		stream   = "stream-1"
	)
	er := &CreateEndpointInput{
		Name: name,
		Protocol: EndpointProtocol(protocol),
		Stream: stream,
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if want, got := endpointBasePath, req.URL.Path; want != got {
			t.Fatalf("Path mismatched: want=%v got=%v", want, got)
		}

		if err := json.NewDecoder(req.Body).Decode(er); err != nil {
			t.Fatal(err)
		}
		defer req.Body.Close()

		if want, got := name, er.Name; want != got {
			t.Fatalf("Name mismatched: want=%s got=%s", want, got)
		}
		if want, got := EndpointProtocol(protocol), er.Protocol; want != got {
			t.Fatalf("Name protocol: want=%s got=%s", want, got)
		}
		if want, got := stream, er.Stream; want != got {
			t.Fatalf("Name stream: want=%s got=%s", want, got)
		}
	}))
	defer server.Close()

	c := testClient(server.Client(), server.URL)

	if err := c.CreateEndpoint(context.Background(), er); err != nil {
		t.Fatal(err)
	}
}

func TestGetEndpoint(t *testing.T) {
	end := &Endpoint{
		Name:              "endpoint",
		Protocol:          EndpointProtocolHttp,
		Host:              "https://endpoint.test",
		Stream:            "stream",
		Ready:             true,
		BasicAuthUsername: "root",
		BasicAuthPassword: "secret",
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if want, got := endpointBasePath+"/"+end.Name, req.URL.Path; want != got {
			t.Fatalf("Path mismatched: want=%v got=%v", want, got)
		}

		if err := json.NewEncoder(w).Encode(end); err != nil {
			t.Fatal(err)
		}

	}))
	defer server.Close()

	c := testClient(server.Client(), server.URL)

	gotEnd, err := c.GetEndpoint(context.Background(), end.Name)
	if err != nil {
		t.Fatal(err)
	}

	if want, got := end, gotEnd; !reflect.DeepEqual(want, got) {
		t.Fatalf("Endpoint mismatched: want=%v got=%v", want, got)
	}
}
