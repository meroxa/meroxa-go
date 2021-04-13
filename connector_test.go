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
	input := CreateConnectorInput{
		Name:         "test",
		ResourceID:   1,
		PipelineID:   2,
		PipelineName: "my-pipeline",

		Configuration: map[string]interface{}{
			"custom_config": true,
		},
		Metadata: map[string]interface{}{
			"custom_metadata": "value",
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// Test request
		type connectionRequest struct {
			Name          string                 `json:"name"`
			Configuration map[string]interface{} `json:"config"`
			ResourceID    int                    `json:"resource_id"`
			PipelineID    int                    `json:"pipeline_id"`
			PipelineName  string                 `json:"pipeline_name"`
			Metadata      map[string]interface{} `json:"metadata"`
		}

		var cr connectionRequest
		err := json.NewDecoder(req.Body).Decode(&cr)
		if err != nil {
			t.Errorf("expected no error, got %+v", err)
		}

		if cr.Name != input.Name {
			t.Errorf("expected name %s, got %s", input.Name, cr.Name)
		}

		if cr.ResourceID != input.ResourceID {
			t.Errorf("expected resource ID %d, got %d", input.ResourceID, cr.ResourceID)
		}

		if cr.PipelineID != input.PipelineID {
			t.Errorf("expected pipeline ID %d, got %d", input.PipelineID, cr.PipelineID)
		}

		if cr.PipelineName != input.PipelineName {
			t.Errorf("expected pipeline name %s, got %s", input.PipelineName, cr.PipelineName)
		}

		if !reflect.DeepEqual(cr.Configuration, input.Configuration) {
			t.Errorf("expected configuration %+v, got %+v", input.Configuration, cr.Configuration)
		}

		if !reflect.DeepEqual(cr.Metadata, input.Metadata) {
			t.Errorf("expected metadata %+v, got %+v", input.Metadata, cr.Metadata)
		}
		defer req.Body.Close()

		// Return response to satisfy client and test response
		c := generateConnector(input.Name, input.ResourceID, input.Configuration, input.Metadata)
		json.NewEncoder(w).Encode(c)
	}))
	// Close the server when test finishes
	defer server.Close()

	c := testClient(server.Client(), server.URL)

	resp, err := c.CreateConnector(context.Background(), input)

	if err != nil {
		t.Errorf("expected no error, got %+v", err)
	}

	if resp.Name != input.Name {
		t.Errorf("expected name %s, got %s", input.Name, resp.Name)
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

func generateConnector(name string, id int, config, metadata map[string]interface{}) Connector {
	if name == "" {
		name = "test"
	}

	if id == 0 {
		id = rand.Intn(10000)
	}

	return Connector{
		ID:            id,
		Type:          "postgres",
		Name:          name,
		Configuration: config,
		Metadata:      metadata,
	}
}
