package meroxa

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

func TestCreateConnector(t *testing.T) {
	name := "test"
	resourceID := 1
	configuration := map[string]string{
		"custom_config": "true",
	}
	metadata := map[string]string{
		"region": "us-east-1",
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// Test request
		type connectionRequest struct {
			Name          string            `json:"name"`
			Configuration map[string]string `json:"config"`
			ResourceID    int               `json:"resource_id"`
			Metadata      map[string]string `json:"metadata"`
		}

		var cr connectionRequest
		err := json.NewDecoder(req.Body).Decode(&cr)
		if err != nil {
			t.Errorf("expected no error, got %+v", err)
		}

		if cr.Name != name {
			t.Errorf("expected name %s, got %s", name, cr.Name)
		}

		if cr.ResourceID != resourceID {
			t.Errorf("expected resource ID %d, got %d", resourceID, cr.ResourceID)
		}

		if !reflect.DeepEqual(cr.Configuration, configuration) {
			t.Errorf("expected configuration %+v, got %+v", configuration, cr.Configuration)
		}

		if !reflect.DeepEqual(cr.Metadata, metadata) {
			t.Errorf("expected metadata %+v, got %+v", metadata, cr.Metadata)
		}
		defer req.Body.Close()

		// Return response to satisfy client and test response
		c := generateConnector(name, resourceID, configuration, metadata)
		json.NewEncoder(w).Encode(c)
	}))
	// Close the server when test finishes
	defer server.Close()

	c := testClient(server.Client(), server.URL)

	resp, err := c.CreateConnector(context.Background(), name, resourceID, configuration, metadata)

	if err != nil {
		t.Errorf("expected no error, got %+v", err)
	}

	if resp.Name != name {
		t.Errorf("expected name %s, got %s", name, resp.Name)
	}

}

func TestUpdateConnectorStatus(t *testing.T) {
	connectorKey := "test-key"
	state := "pause"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if want, got := fmt.Sprintf("/v1/connectors/%s/status", connectorKey), req.URL.Path; want != got {
			t.Fatalf("mismatched of request path: want=%s got=%s", want, got)
		}

		// Test request
		type connectionRequest struct {
			State string `json:"state"`
		}

		var cr connectionRequest
		if err := json.NewDecoder(req.Body).Decode(&cr); err != nil {
			t.Errorf("expected no error, got %+v", err)
		}
		defer req.Body.Close()

		if cr.State != state {
			t.Errorf("expected state %s, got %s", state, cr.State)
		}

		// Return response to satisfy client and test response
		c := generateConnector(connectorKey, 0, nil, nil)
		c.State = state
		json.NewEncoder(w).Encode(c)
	}))
	// Close the server when test finishes
	defer server.Close()

	c := testClient(server.Client(), server.URL)

	resp, err := c.UpdateConnectorStatus(context.Background(), connectorKey, state)
	if err != nil {
		t.Errorf("expected no error, got %+v", err)
	}

	if resp.State != state {
		t.Errorf("expected state %s, got %s", state, resp.State)
	}
}

func testClient(c *http.Client, u string) *Client {
	parsedURL, _ := url.Parse(u)
	return &Client{
		BaseURL:    parsedURL,
		httpClient: c,
	}
}

func generateConnector(name string, id int, config, metadata map[string]string) Connector {
	if name == "" {
		name = "test"
	}

	if id == 0 {
		id = rand.Intn(10000)
	}

	return Connector{
		ID:            id,
		Kind:          "postgres",
		Name:          name,
		Configuration: config,
		Metadata:      metadata,
	}
}
