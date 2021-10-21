package meroxa

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const connectorsBasePath = "/v1/connectors"

type ConnectorState string

const (
	ConnectorStatePending ConnectorState = "pending"
	ConnectorStateRunning                = "running"
	ConnectorStatePaused                 = "paused"
	ConnectorStateCrashed                = "crashed"
	ConnectorStateFailed                 = "failed"
	ConnectorStateDOA                    = "doa"
)

type Action string

var (
	ActionPause   Action = "pause"
	ActionResume  Action = "resume"
	ActionRestart Action = "restart"
)

type Connector struct {
	ID            int                    `json:"id"`
	Type          string                 `json:"type"`
	Name          string                 `json:"name"`
	Configuration map[string]interface{} `json:"config"`
	Streams       map[string]interface{} `json:"streams"`
	State         ConnectorState         `json:"state"`
	Trace         string                 `json:"trace,omitempty"`
	PipelineID    int                    `json:"pipeline_id"`
	PipelineName  string                 `json:"pipeline_name"`
}

type CreateConnectorInput struct {
	Name          string                 `json:"name,omitempty"`
	ResourceID    int                    `json:"resource_id"`
	PipelineID    int                    `json:"pipeline_id,omitempty"`
	PipelineName  string                 `json:"pipeline_name,omitempty"`
	Configuration map[string]interface{} `json:"config,omitempty"`
}

type UpdateConnectorInput struct {
	Name          string                 `json:"name,omitempty"`
	Configuration map[string]interface{} `json:"config,omitempty"`
}

// CreateConnector provisions a connector between the Resource and the Meroxa
// platform
func (c *Client) CreateConnector(ctx context.Context, input CreateConnectorInput) (*Connector, error) {
	resp, err := c.MakeRequest(ctx, http.MethodPost, connectorsBasePath, input, nil)
	if err != nil {
		return nil, err
	}

	err = handleAPIErrors(resp)
	if err != nil {
		return nil, err
	}

	var con Connector
	err = json.NewDecoder(resp.Body).Decode(&con)
	if err != nil {
		return nil, err
	}

	return &con, nil
}

// @TODO implement connector status
// UpdateConnectorStatus updates the status of a connector
func (c *Client) UpdateConnectorStatus(ctx context.Context, nameOrId, state Action) (*Connector, error) {
	path := fmt.Sprintf("%s/%s/status", connectorsBasePath, nameOrId)

	cr := struct {
		State Action `json:"state,omitempty"`
	}{
		State: state,
	}

	resp, err := c.MakeRequest(ctx, http.MethodPost, path, cr, nil)
	if err != nil {
		return nil, err
	}

	err = handleAPIErrors(resp)
	if err != nil {
		return nil, err
	}

	var con Connector
	err = json.NewDecoder(resp.Body).Decode(&con)
	if err != nil {
		return nil, err
	}

	return &con, nil
}

// UpdateConnector updates the name, or a configuration of a connector
func (c *Client) UpdateConnector(ctx context.Context, nameOrId string, input UpdateConnectorInput) (*Connector, error) {
	path := fmt.Sprintf("%s/%s", connectorsBasePath, nameOrId)

	resp, err := c.MakeRequest(ctx, http.MethodPatch, path, input, nil)
	if err != nil {
		return nil, err
	}

	err = handleAPIErrors(resp)
	if err != nil {
		return nil, err
	}

	var con Connector
	err = json.NewDecoder(resp.Body).Decode(&con)
	if err != nil {
		return nil, err
	}

	return &con, nil
}

// ListPipelineConnectors returns an array of Connectors (scoped to the calling user)
func (c *Client) ListPipelineConnectors(ctx context.Context, pipelineID int) ([]*Connector, error) {
	path := fmt.Sprintf("/v1/pipelines/%d/connectors", pipelineID)

	resp, err := c.MakeRequest(ctx, http.MethodGet, path, nil, nil)
	if err != nil {
		return nil, err
	}

	err = handleAPIErrors(resp)
	if err != nil {
		return nil, err
	}

	var rr []*Connector
	err = json.NewDecoder(resp.Body).Decode(&rr)
	if err != nil {
		return nil, err
	}

	return rr, nil
}

// ListConnectors returns an array of Connectors (scoped to the calling user)
func (c *Client) ListConnectors(ctx context.Context) ([]*Connector, error) {
	resp, err := c.MakeRequest(ctx, http.MethodGet, connectorsBasePath, nil, nil)
	if err != nil {
		return nil, err
	}

	err = handleAPIErrors(resp)
	if err != nil {
		return nil, err
	}

	var rr []*Connector
	err = json.NewDecoder(resp.Body).Decode(&rr)
	if err != nil {
		return nil, err
	}

	return rr, nil
}

// GetConnector returns a Connector for the given connector ID
func (c *Client) GetConnector(ctx context.Context, id int) (*Connector, error) {
	path := fmt.Sprintf("%s/%d", connectorsBasePath, id)

	resp, err := c.MakeRequest(ctx, http.MethodGet, path, nil, nil)
	if err != nil {
		return nil, err
	}

	err = handleAPIErrors(resp)
	if err != nil {
		return nil, err
	}

	var con Connector
	err = json.NewDecoder(resp.Body).Decode(&con)
	if err != nil {
		return nil, err
	}

	return &con, nil
}

// GetConnectorByName returns a Connector with the given name
func (c *Client) GetConnectorByName(ctx context.Context, name string) (*Connector, error) {
	params := map[string][]string{
		"name": {name},
	}

	resp, err := c.MakeRequest(ctx, http.MethodGet, connectorsBasePath, nil, params)
	if err != nil {
		return nil, err
	}

	err = handleAPIErrors(resp)
	if err != nil {
		return nil, err
	}

	var con Connector
	err = json.NewDecoder(resp.Body).Decode(&con)
	if err != nil {
		return nil, err
	}

	return &con, nil
}

// DeleteConnector deletes the Connector with the given id
func (c *Client) DeleteConnector(ctx context.Context, id int) error {
	path := fmt.Sprintf("%s/%d", connectorsBasePath, id)

	resp, err := c.MakeRequest(ctx, http.MethodDelete, path, nil, nil)
	if err != nil {
		return err
	}

	err = handleAPIErrors(resp)
	if err != nil {
		return err
	}

	return nil
}
