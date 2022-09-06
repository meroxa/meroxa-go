package meroxa

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/volatiletech/null/v8"
)

type DeploymentState string

const (
	DeploymentStateDeploying        DeploymentState = "deploying"
	DeploymentStateDeployingErrored DeploymentState = "deploying_error"
	DeploymentStateRollingBack      DeploymentState = "rolling_back"
	DeploymentStateDeployed         DeploymentState = "deployed"
)

type DeploymentStatus struct {
	State   DeploymentState `json:"state"`
	Details string          `json:"details,omitempty"`
}

type Deployment struct {
	UUID        string           `json:"uuid"`
	GitSha      string           `json:"git_sha"`
	Application EntityIdentifier `json:"application"`
	OutputLog   null.String      `json:"output_log,omitempty"`
	CreatedAt   time.Time        `json:"created_at"`
	DeletedAt   time.Time        `json:"deleted_at,omitempty"`
	State       DeploymentStatus `json:"state"`
	Spec        null.String      `json:"spec,omitempty"`
	SpecVersion null.String      `json:"spec_version,omitempty"`
	// createdBy?
}

type CreateDeploymentInput struct {
	GitSha      string           `json:"git_sha"`
	Application EntityIdentifier `json:"application"`
	Spec        null.String      `json:"spec,omitempty"`
	SpecVersion null.String      `json:"spec_version,omitempty"`
}

func (c *client) GetLatestDeployment(ctx context.Context, appName string) (*Deployment, error) {
	resp, err := c.MakeRequest(ctx, http.MethodGet, fmt.Sprintf("%s/%s/deployments/latest", applicationsBasePath, appName), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	err = handleAPIErrors(resp)
	if err != nil {
		return nil, err
	}

	var d *Deployment
	err = json.NewDecoder(resp.Body).Decode(&d)
	if err != nil {
		return nil, err
	}

	return d, nil
}

func (c *client) CreateDeployment(ctx context.Context, input *CreateDeploymentInput) (*Deployment, error) {
	appIdentifier, err := input.Application.GetNameOrUUID()

	if err != nil {
		return nil, err
	}

	resp, err := c.MakeRequest(ctx, http.MethodPost, fmt.Sprintf("%s/%s/deployments", applicationsBasePath, appIdentifier), input, nil, nil)
	if err != nil {
		return nil, err
	}

	err = handleAPIErrors(resp)
	if err != nil {
		return nil, err
	}

	var d *Deployment
	err = json.NewDecoder(resp.Body).Decode(&d)
	if err != nil {
		return nil, err
	}

	return d, nil
}
