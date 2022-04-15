package meroxa

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"

	"github.com/volatiletech/null/v8"
)

func TestCreateConnector(t *testing.T) {
	input := CreateConnectorInput{
		Name:         "test",
		ResourceName: "my-resource",
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
			ResourceName  string                 `json:"resource_name"`
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

		if cr.ResourceName != input.ResourceName {
			t.Errorf("expected resource Name %s, got %s", input.ResourceName, cr.ResourceName)
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
		c := generateConnector(input.Name, input.Configuration, input.Metadata)
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
		if want, got := fmt.Sprintf("%s/%s/status", connectorsBasePath, connectorKey), req.URL.Path; want != got {
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
		c := generateConnector(connectorKey, nil, nil)
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
	var connector = generateConnector("", nil, nil)
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

func generateConnector(name string, config, metadata map[string]interface{}) Connector {
	if name == "" {
		name = "test"
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
		Type:          "postgres",
		Name:          name,
		Configuration: config,
		Metadata:      metadata,
		Environment:   &EntityIdentifier{Name: null.StringFrom("my-env")},
	}
}

func TestGetConnectorByName(t *testing.T) {
	connector := generateConnector("my-connector", nil, nil)

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

func TestListPipelineConnectors(t *testing.T) {
	p := generatePipeline("", fmt.Sprint(PipelineStateHealthy), nil)
	connector := generateConnector("", nil, nil)
	connector.PipelineName = p.Name
	list := []*Connector{&connector}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if want, got := fmt.Sprintf("%s/%s/connectors", pipelinesBasePath, p.Name), req.URL.Path; want != got {
			t.Fatalf("mismatched of request path: want=%s got=%s", want, got)
		}
		defer req.Body.Close()

		// Return response to satisfy client and test response
		json.NewEncoder(w).Encode(list)
	}))
	// Close the server when test finishes
	defer server.Close()

	c := testClient(server.Client(), server.URL)

	resp, err := c.ListPipelineConnectors(context.Background(), p.Name)
	if err != nil {
		t.Errorf("expected no error, got %+v", err)
	}

	if !reflect.DeepEqual(resp, list) {
		t.Errorf("expected response same as list")
	}
}

func TestListConnectors(t *testing.T) {
	c1 := generateConnector("", nil, nil)
	c2 := generateConnector("", nil, nil)
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
