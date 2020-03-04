package meroxa

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// Pipeline represents the Meroxa Pipeline type within the Meroxa API
type Pipeline struct {
	Name     string            `json:"name"`
	Metadata map[string]string `json:"metadata"`
}

// ComponentKind enum for Component "kinds" within Pipeline stages
type ComponentKind int32

const (
	// ConnectorComponent is a Pipeline stage component of type Connector
	ConnectorComponent ComponentKind = 0

	// FunctionComponent is a Pipeline stage component of type Function
	FunctionComponent = 1
)

// PipelineStage represents the Meroxa PipelineStage type within the Meroxa API
type PipelineStage struct {
	PipelineID    int32 `json:"pipeline_id"`
	ComponentID   int32 `json:"component_id"`
	ComponentKind int32 `json:"component_kind"`
}

// CreatePipeline provisions a new Pipeline
func (c *Client) CreatePipeline(ctx context.Context, pipeline *Pipeline) (*Pipeline, error) {
	path := fmt.Sprintf("/v1/pipelines")

	resp, err := c.makeRequest(ctx, http.MethodPost, path, pipeline, nil)
	if err != nil {
		return nil, err
	}

	var p Pipeline
	err = json.NewDecoder(resp.Body).Decode(&p)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

// ListPipelines returns an array of Pipelines (scoped to the calling user)
func (c *Client) ListPipelines(ctx context.Context) ([]*Pipeline, error) {
	path := fmt.Sprintf("/v1/pipelines")

	resp, err := c.makeRequest(ctx, http.MethodGet, path, nil, nil)
	if err != nil {
		return nil, err
	}

	var pp []*Pipeline
	err = json.NewDecoder(resp.Body).Decode(&pp)
	if err != nil {
		return nil, err
	}

	return pp, nil
}

// GetPipeline returns a Pipeline with the given id
func (c *Client) GetPipeline(ctx context.Context, id int) (*Pipeline, error) {
	path := fmt.Sprintf("/v1/pipelines/%d", id)

	resp, err := c.makeRequest(ctx, http.MethodGet, path, nil, nil)
	if err != nil {
		return nil, err
	}

	var p Pipeline
	err = json.NewDecoder(resp.Body).Decode(&p)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

// DeletePipeline deletes the Pipeline with the given id
func (c *Client) DeletePipeline(ctx context.Context, id int) error {
	path := fmt.Sprintf("/v1/pipelines/%d", id)

	_, err := c.makeRequest(ctx, http.MethodDelete, path, nil, nil)
	if err != nil {
		return err
	}

	return nil
}

// GetPipelineStages returns an array of Pipeline Stages for the given Pipeline id
func (c *Client) GetPipelineStages(ctx context.Context, pipelineID int) ([]*PipelineStage, error) {
	path := fmt.Sprintf("/v1/pipelines/%d/stages", pipelineID)

	resp, err := c.makeRequest(ctx, http.MethodGet, path, nil, nil)
	if err != nil {
		return nil, err
	}

	var pp []*PipelineStage
	err = json.NewDecoder(resp.Body).Decode(&pp)
	if err != nil {
		return nil, err
	}

	return pp, nil
}

// AddPipelineStage adds the PipelineStage to the given Pipeline
func (c *Client) AddPipelineStage(ctx context.Context, pipelineID int, connectorID int) (*PipelineStage, error) {
	path := fmt.Sprintf("/v1/pipelines/%d/stages", pipelineID)

	params := map[string][]string{
		"connector": []string{strconv.Itoa(connectorID)},
	}

	resp, err := c.makeRequest(ctx, http.MethodPost, path, nil, params)
	if err != nil {
		return nil, err
	}

	var s PipelineStage
	err = json.NewDecoder(resp.Body).Decode(&s)
	if err != nil {
		return nil, err
	}

	return &s, nil
}

// RemovePipelineStage removes the PipelineStage from the given Pipeline
func (c *Client) RemovePipelineStage(ctx context.Context, pipelineID int, stageID int) error {
	path := fmt.Sprintf("/v1/pipelines/%d/stages/%d", pipelineID, stageID)

	resp, err := c.makeRequest(ctx, http.MethodDelete, path, nil, nil)
	if err != nil {
		return err
	}

	if resp.StatusCode > 204 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("Status %d, %v", resp.StatusCode, string(body))
	}

	return nil
}
