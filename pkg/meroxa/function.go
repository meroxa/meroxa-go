package meroxa

import (
	"context"
	"encoding/json"
	"net/http"
)

const functionsBasePath = "/v1/functions"

type Function struct {
	UUID         string              `json:"uuid"`
	Name         string              `json:"name"`
	InputStream  string              `json:"input_stream"`
	OutputStream string              `json:"output_stream"`
	Image        string              `json:"image"`
	Command      []string            `json:"command"`
	Args         []string            `json:"args"`
	EnvVars      map[string]string   `json:"env_vars"`
	Status       FunctionStatus      `json:"status"`
	Pipeline     PipelineIdentifier `json:"pipeline"`
}

type FunctionStatus struct {
	State   string `json:"state"`
	Details string `json:"details"`
}

type CreateFunctionInput struct {
	Name         string            `json:"name"`
	InputStream  string            `json:"input_stream"`
	OutputStream string            `json:"output_stream"`
	PipelineName string            `json:"pipeline_name"`
	Image        string            `json:"image"`
	Command      []string          `json:"command"`
	Args         []string          `json:"args"`
	EnvVars      map[string]string `json:"env_vars"`
}

func (c *client) CreateFunction(ctx context.Context, input *CreateFunctionInput) (*Function, error) {
	resp, err := c.MakeRequest(ctx, http.MethodPost, functionsBasePath, input, nil)
	if err != nil {
		return nil, err
	}

	if err := handleAPIErrors(resp); err != nil {
		return nil, err
	}

	var fun Function
	if err := json.NewDecoder(resp.Body).Decode(&fun); err != nil {
		return nil, err
	}

	return &fun, nil
}
