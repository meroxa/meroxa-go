package meroxa

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const environmentsBasePath = "/v1/environments"

type EnvironmentState string

const (
	EnvironmentStateProvisioning   EnvironmentState = "provisioning"
	EnvironmentStateProvisioned                     = "provisioned"
	EnvironmentStateUpdating                        = "updating"
	EnvironmentStateError                           = "error"
	EnvironmentStateRepairing                       = "repairing"
	EnvironmentStateDeprovisioning                  = "deprovisioning"
	EnvironmentStateDeprovisioned                   = "deprovisioned"
)

type EnvironmentViewStatus struct {
	State   EnvironmentState `json:"state"`
	Details string           `json:"details,omitempty"`
}

type EnvironmentRegion string

const (
	EnvironmentRegionAfSouth      EnvironmentRegion = "af-south-1"
	EnvironmentRegionApEast                         = "ap-east-1"
	EnvironmentRegionApNortheast1                   = "ap-northeast-1"
	EnvironmentRegionApNortheast2                   = "ap-northeast-2"
	EnvironmentRegionApNortheast3                   = "ap-northeast-3"
	EnvironmentRegionApSouth                        = "ap-south-1"
	EnvironmentRegionApSoutheast1                   = "ap-southeast-1"
	EnvironmentRegionApSoutheast2                   = "ap-southeast-2"
	EnvironmentRegionCaCentral                      = "ca-central-1"
	EnvironmentRegionEuCentral                      = "eu-central-1"
	EnvironmentRegionEuNorth                        = "eu-north-1"
	EnvironmentRegionEuSouth                        = "eu-south-1"
	EnvironmentRegionEuWest1                        = "eu-west-1"
	EnvironmentRegionEuWest2                        = "eu-west-2"
	EnvironmentRegionEuWest3                        = "eu-west-3"
	EnvironmentRegionMeSouth                        = "me-south-1"
	EnvironmentRegionSaEast1                        = "sa-east-1"
	EnvironmentRegionUsEast1                        = "us-east-1"
	EnvironmentRegionUsEast2                        = "us-east-2"
	EnvironmentRegionUsWest2                        = "us-west-2"
)

type EnvironmentType string

const (
	EnvironmentTypeHosted    EnvironmentType = "hosted"
	EnvironmentTypeDedicated                 = "dedicated"
	EnvironmentTypeCommon                    = "common"
)

type EnvironmentProvider string

const (
	EnvironmentProviderAws EnvironmentProvider = "aws"
)

// Environment represents the Meroxa Environment type within the Meroxa API
type Environment struct {
	UUID          string                 `json:"uuid"`
	Name          string                 `json:"name"`
	Provider      EnvironmentProvider    `json:"provider"`
	Region        EnvironmentRegion      `json:"region"`
	Type          EnvironmentType        `json:"type"`
	Configuration map[string]interface{} `json:"config,omitempty"`
	Status        EnvironmentViewStatus  `json:"status"`
	CreatedAt     time.Time              `json:"created_at"`
	UpdatedAt     time.Time              `json:"updated_at"`
}

// CreateEnvironmentInput represents the input for a Meroxa Environment we're creating within the Meroxa API
type CreateEnvironmentInput struct {
	Type          EnvironmentType        `json:"type,omitempty"`
	Provider      EnvironmentProvider    `json:"provider,omitempty"`
	Name          string                 `json:"name,omitempty"`
	Configuration map[string]interface{} `json:"config"`
	Region        EnvironmentRegion      `json:"region,omitempty"`
}

// ListEnvironments returns an array of Environments (scoped to the calling user)
func (c *client) ListEnvironments(ctx context.Context) ([]*Environment, error) {
	resp, err := c.MakeRequest(ctx, http.MethodGet, environmentsBasePath, nil, nil)
	if err != nil {
		return nil, err
	}

	err = handleAPIErrors(resp)
	if err != nil {
		return nil, err
	}

	var ee []*Environment
	err = json.NewDecoder(resp.Body).Decode(&ee)
	if err != nil {
		return nil, err
	}

	return ee, nil
}

// CreateEnvironment creates a new Environment based on a CreateEnvironmentInput
func (c *client) CreateEnvironment(ctx context.Context, input *CreateEnvironmentInput) (*Environment, error) {
	resp, err := c.MakeRequest(ctx, http.MethodPost, environmentsBasePath, input, nil)
	if err != nil {
		return nil, err
	}

	err = handleAPIErrors(resp)
	if err != nil {
		return nil, err
	}

	var e Environment
	err = json.NewDecoder(resp.Body).Decode(&e)
	if err != nil {
		return nil, err
	}

	return &e, nil
}

func (c *client) GetEnvironment(ctx context.Context, nameOrUUID string) (*Environment, error) {
	path := fmt.Sprintf("%s/%s", environmentsBasePath, nameOrUUID)
	resp, err := c.MakeRequest(ctx, http.MethodGet, path, nil, nil)
	if err != nil {
		return nil, err
	}

	err = handleAPIErrors(resp)
	if err != nil {
		return nil, err
	}

	var e *Environment
	err = json.NewDecoder(resp.Body).Decode(&e)
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (c *client) DeleteEnvironment(ctx context.Context, nameOrUUID string) (*Environment, error) {
	path := fmt.Sprintf("%s/%s", environmentsBasePath, nameOrUUID)
	resp, err := c.MakeRequest(ctx, http.MethodDelete, path, nil, nil)
	if err != nil {
		return nil, err
	}

	err = handleAPIErrors(resp)
	if err != nil {
		return nil, err
	}

	var e *Environment
	err = json.NewDecoder(resp.Body).Decode(&e)

	return e, nil
}
