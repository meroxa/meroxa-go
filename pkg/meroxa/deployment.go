package meroxa

import (
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
	UUID            string           `json:"uuid"`
	GitSha          string           `json:"git_sha"`
	ApplicationUUID string           `json:"application_uuid"`
	OutputLog       null.String      `json:"output_log,omitempty"`
	CreatedAt       time.Time        `json:"created_at"`
	DeletedAt       time.Time        `json:"deleted_at,omitempty"`
	State           DeploymentStatus `json:"state"`
	Spec            null.String      `json:"spec,omitempty"`
	SpecVersion     null.String      `json:"spec_version,omitempty"`
}
