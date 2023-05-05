package meroxa

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const flinkJobsBasePath = "/v1/flink-jobs"

type FlinkJobState string

const (
	FlinkJobStateCancelled     FlinkJobState = "cancelled"
	FlinkJobStateCancelling    FlinkJobState = "cancelling"
	FlinkJobStateCreated       FlinkJobState = "created"
	FlinkJobStateDOA           FlinkJobState = "doa"
	FlinkJobStateFailed        FlinkJobState = "failed"
	FlinkJobStateFailing       FlinkJobState = "failing"
	FlinkJobStateFinished      FlinkJobState = "finished"
	FlinkJobStateInitializing  FlinkJobState = "initializing"
	FlinkJobStateReconciling   FlinkJobState = "reconciling"
	FlinkJobStateRestarting    FlinkJobState = "restarting"
	FlinkJobStateRunning       FlinkJobState = "running"
	FlinkJobStateSuspended     FlinkJobState = "suspended"
	FlinkJobStateUninitialized FlinkJobState = "uninitialized"
)

type FlinkJobStatus struct {
	State   FlinkJobState `json:"state"`
	Details string        `json:"details,omitempty"`
}

type FlinkJob struct {
	UUID         string           `json:"uuid"`
	Name         string           `json:"name"`
	OutputStream string           `json:"output_stream,omitempty"`
	Environment  EntityIdentifier `json:"environment,omitempty"`
	Status       FlinkJobStatus   `json:"status"`
	CreatedAt    time.Time        `json:"created_at"`
	UpdatedAt    time.Time        `json:"updated_at"`
}

type CreateFlinkJobInput struct {
	Name   string `json:"name"`
	JarURL string `json:"jar_url"`
}

func (c *client) GetFlinkJob(ctx context.Context, nameOrUUID string) (*FlinkJob, error) {
	resp, err := c.MakeRequest(ctx, http.MethodGet, fmt.Sprintf("%s/%s", flinkJobsBasePath, nameOrUUID), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	if err = handleAPIErrors(resp); err != nil {
		return nil, err
	}

	var fj *FlinkJob
	if err = json.NewDecoder(resp.Body).Decode(&fj); err != nil {
		return nil, err
	}

	return fj, nil
}

func (c *client) ListFlinkJobs(ctx context.Context) ([]*FlinkJob, error) {
	resp, err := c.MakeRequest(ctx, http.MethodGet, flinkJobsBasePath, nil, nil, nil)
	if err != nil {
		return nil, err
	}

	if err = handleAPIErrors(resp); err != nil {
		return nil, err
	}

	var fjs []*FlinkJob
	if err = json.NewDecoder(resp.Body).Decode(&fjs); err != nil {
		return nil, err
	}

	return fjs, nil
}

func (c *client) CreateFlinkJob(ctx context.Context, input *CreateFlinkJobInput) (*FlinkJob, error) {
	resp, err := c.MakeRequest(ctx, http.MethodPost, flinkJobsBasePath, input, nil, nil)
	if err != nil {
		return nil, err
	}

	if err = handleAPIErrors(resp); err != nil {
		return nil, err
	}

	var fj *FlinkJob
	if err = json.NewDecoder(resp.Body).Decode(&fj); err != nil {
		return nil, err
	}

	return fj, nil
}

func (c *client) DeleteFlinkJob(ctx context.Context, nameOrUUID string) error {
	resp, err := c.MakeRequest(ctx, http.MethodDelete, fmt.Sprintf("%s/%s", flinkJobsBasePath, nameOrUUID), nil, nil, nil)
	if err != nil {
		return err
	}

	return handleAPIErrors(resp)
}
