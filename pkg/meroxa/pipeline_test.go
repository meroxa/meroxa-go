package meroxa

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestUpdatePipelineStatus(t *testing.T) {
	name := "test"
	pipelineID := 1234567
	state := "pause"
	newState := "healthy"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if want, got := fmt.Sprintf("%s/%d/status", pipelinesBasePath, pipelineID), req.URL.Path; want != got {
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
		p := generatePipeline(name, pipelineID, newState, nil)
		json.NewEncoder(w).Encode(p)
	}))
	// Close the server when test finishes
	defer server.Close()

	c := testClient(server.Client(), server.URL)

	resp, err := c.UpdatePipelineStatus(context.Background(), pipelineID, Action(state))
	if err != nil {
		t.Errorf("expected no error, got %+v", err)
	}

	if string(resp.State) != newState {
		t.Errorf("expected state %s, got %s", state, resp.State)
	}
}

func TestUpdatePipeline(t *testing.T) {
	var pipelineUpdate UpdatePipelineInput
	var pipeline = generatePipeline("", 0, "", nil)

	pipelineUpdate.Name = pipeline.Name

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if want, got := fmt.Sprintf("%s/%d", pipelinesBasePath, pipeline.ID), req.URL.Path; want != got {
			t.Fatalf("mismatched of request path: want=%s got=%s", want, got)
		}

		var pi UpdatePipelineInput

		if err := json.NewDecoder(req.Body).Decode(&pi); err != nil {
			t.Errorf("expected no error, got %+v", err)
		}
		defer req.Body.Close()

		if pi.Name != pipeline.Name {
			t.Errorf("expected name %s, got %s", pipeline.Name, pi.Name)
		}

		// Return response to satisfy client and test response
		json.NewEncoder(w).Encode(pipeline)
	}))
	// Close the server when test finishes
	defer server.Close()

	c := testClient(server.Client(), server.URL)

	resp, err := c.UpdatePipeline(context.Background(), pipeline.ID, &pipelineUpdate)
	if err != nil {
		t.Errorf("expected no error, got %+v", err)
	}

	if !reflect.DeepEqual(resp, &pipeline) {
		t.Errorf("expected response same as pipeline")
	}
}

func generatePipeline(name string, id int, state string, metadata map[string]interface{}) Pipeline {
	if name == "" {
		name = "test"
	}

	if state == "" {
		state = "healthy"
	}

	if id == 0 {
		id = rand.Intn(10000)
	}

	if metadata == nil {
		metadata = map[string]interface{}{
			"custom_metadata": true,
		}
	}

	return Pipeline{
		ID:       id,
		Name:     name,
		State:    PipelineState(state),
	}
}
