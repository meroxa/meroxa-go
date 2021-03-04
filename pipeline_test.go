package meroxa

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
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

	resp, err := c.UpdatePipelineStatus(context.Background(), pipelineID, state)
	if err != nil {
		t.Errorf("expected no error, got %+v", err)
	}

	if resp.State != newState {
		t.Errorf("expected state %s, got %s", state, resp.State)
	}
}

func generatePipeline(name string, id int, state string, metadata map[string]string) Pipeline {
	if name == "" {
		name = "test"
	}

	if id == 0 {
		id = rand.Intn(10000)
	}

	return Pipeline{
		ID:       id,
		Name:     name,
		State:    state,
		Metadata: metadata,
	}
}
