package meroxa

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetEnvironments(t *testing.T) {
	env := &Environment{
		Type:     "dedicated",
		Name:     "environment-1234",
		Provider: "aws",
		Region:   "aws:us-east",
		State:    "provisioned",
		ID:       "1234",
	}

	environments := []*Environment{env}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if want, got := environmentsBasePath, req.URL.Path; want != got {
			t.Fatalf("Path mismatched: want=%v got=%v", want, got)
		}

		if err := json.NewEncoder(w).Encode(environments); err != nil {
			t.Fatal(err)
		}

	}))
	defer server.Close()

	c := testClient(server.Client(), server.URL)

	gotEnv, err := c.ListEnvironments(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if want, got := environments, gotEnv; !reflect.DeepEqual(want, got) {
		t.Fatalf("Environments mismatched: want=%v got=%v", want, got)
	}
}
