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
			"region": "us-east-1",
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

		if !reflect.DeepEqual(input.Metadata, cr.Metadata) {
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

	resp, err := c.CreateConnector(context.Background(), &input)

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
		c.State = ConnectorState(state)
		json.NewEncoder(w).Encode(c)
	}))
	// Close the server when test finishes
	defer server.Close()

	c := testClient(server.Client(), server.URL)

	resp, err := c.UpdateConnectorStatus(context.Background(), connectorKey, Action(state))
	if err != nil {
		t.Errorf("expected no error, got %+v", err)
	}

	if resp.State != ConnectorState(state) {
		t.Errorf("expected state %s, got %s", state, resp.State)
	}
}

func TestUpdateConnector(t *testing.T) {
	var connector = generateConnector("", 0, nil, nil)
	var connectorUpdate UpdateConnectorInput
	connectorUpdate.Name = connector.Name
	connectorUpdate.Configuration = connector.Configuration

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if want, got := fmt.Sprintf("%s/%s", connectorsBasePath, connector.Name), req.URL.Path; want != got {
			t.Fatalf("mismatched of request path: want=%s got=%s", want, got)
		}

		var cu UpdateConnectorInput

		if err := json.NewDecoder(req.Body).Decode(&cu); err != nil {
			t.Errorf("expected no error, got %+v", err)
		}
		defer req.Body.Close()

		if cu.Name != connector.Name {
			t.Errorf("expected name %s, got %s", cu.Name, cu.Name)
		}

		if !reflect.DeepEqual(cu.Configuration, connector.Configuration) {
			t.Errorf("expected same configuration")
		}

		// Return response to satisfy client and test response
		json.NewEncoder(w).Encode(connector)
	}))
	// Close the server when test finishes
	defer server.Close()

	c := testClient(server.Client(), server.URL)

	resp, err := c.UpdateConnector(context.Background(), connector.Name, &connectorUpdate)
	if err != nil {
		t.Errorf("expected no error, got %+v", err)
	}

	if !reflect.DeepEqual(resp, &connector) {
		t.Errorf("expected response same as connector")
	}
}

func testClient(c *http.Client, u string) Client {
	parsedURL, _ := url.Parse(u)
	return &client{
		baseURL:    parsedURL,
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

	if config == nil {
		config = map[string]interface{}{
			"key": "value",
		}
	}

	if metadata == nil {
		metadata = map[string]interface{}{
			"mx:key": "value",
		}
	}

	return Connector{
		ID:            id,
		Type:          "postgres",
		Name:          name,
		Configuration: config,
		Metadata:      metadata,
		Environment: &EnvironmentIdentifier{Name: "my-env"},
	}
}

func TestGetConnectorByName(t *testing.T) {
	connector := generateConnector("my-connector", 0, nil, nil)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if want, got := fmt.Sprintf("%s/%s", connectorsBasePath, connector.Name), req.URL.Path; want != got {
			t.Fatalf("mismatched of request path: want=%s got=%s", want, got)
		}

		defer req.Body.Close()

		// Return response to satisfy client and test response
		json.NewEncoder(w).Encode(connector)
	}))
	// Close the server when test finishes
	defer server.Close()

	c := testClient(server.Client(), server.URL)

	resp, err := c.GetConnectorByNameOrID(context.Background(), connector.Name)
	if err != nil {
		t.Errorf("expected no error, got %+v", err)
	}

	if !reflect.DeepEqual(resp, &connector) {
		t.Errorf("expected response same as connector")
	}
}

func TestGetConnectorByID(t *testing.T) {
	connector := generateConnector("", 10, nil, nil)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if want, got := fmt.Sprintf("%s/%d", connectorsBasePath, connector.ID), req.URL.Path; want != got {
			t.Fatalf("mismatched of request path: want=%s got=%s", want, got)
		}

		defer req.Body.Close()

		// Return response to satisfy client and test response
		json.NewEncoder(w).Encode(connector)
	}))
	// Close the server when test finishes
	defer server.Close()

	c := testClient(server.Client(), server.URL)

	resp, err := c.GetConnectorByNameOrID(context.Background(), fmt.Sprint(connector.ID))
	if err != nil {
		t.Errorf("expected no error, got %+v", err)
	}

	if !reflect.DeepEqual(resp, &connector) {
		t.Errorf("expected response same as connector")
	}
}

func TestListPipelineConnectors(t *testing.T) {
	p := generatePipeline("", 0, fmt.Sprint(PipelineStateHealthy), nil)
	connector := generateConnector("", 0, nil, nil)
	connector.PipelineID = p.ID
	list := []*Connector{&connector}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if want, got := fmt.Sprintf("/v1/pipelines/%d/connectors", p.ID), req.URL.Path; want != got {
			t.Fatalf("mismatched of request path: want=%s got=%s", want, got)
		}

		defer req.Body.Close()

		// Return response to satisfy client and test response
		json.NewEncoder(w).Encode(list)
	}))
	// Close the server when test finishes
	defer server.Close()

	c := testClient(server.Client(), server.URL)

	resp, err := c.ListPipelineConnectors(context.Background(), p.ID)
	if err != nil {
		t.Errorf("expected no error, got %+v", err)
	}

	if !reflect.DeepEqual(resp, list) {
		t.Errorf("expected response same as list")
	}
}

func TestListConnectors(t *testing.T) {
	c1 := generateConnector("", 10, nil, nil)
	c2 := generateConnector("", 10, nil, nil)
	list := []*Connector{&c1, &c2}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if want, got := fmt.Sprintf("%s", connectorsBasePath), req.URL.Path; want != got {
			t.Fatalf("mismatched of request path: want=%s got=%s", want, got)
		}

		defer req.Body.Close()

		// Return response to satisfy client and test response
		json.NewEncoder(w).Encode(list)
	}))
	// Close the server when test finishes
	defer server.Close()

	c := testClient(server.Client(), server.URL)

	resp, err := c.ListConnectors(context.Background())
	if err != nil {
		t.Errorf("expected no error, got %+v", err)
	}

	if !reflect.DeepEqual(resp, list) {
		t.Errorf("expected response same as list")
	}
}
