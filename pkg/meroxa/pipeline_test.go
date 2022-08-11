package meroxa

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/volatiletech/null/v8"
)

func TestUpdatePipelineStatus(t *testing.T) {
	name := "test"
	pipelineName := "pipeline-1234567"
	state := "pause"
	newState := "healthy"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if want, got := fmt.Sprintf("%s/%s/status", pipelinesBasePath, pipelineName), req.URL.Path; want != got {
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
		p := generatePipeline(name, newState)
		if err := json.NewEncoder(w).Encode(p); err != nil {
			t.Errorf("expected no error, got %+v", err)
		}
	}))
	// Close the server when test finishes
	defer server.Close()

	c := testClient(server.Client(), server.URL)

	resp, err := c.UpdatePipelineStatus(context.Background(), pipelineName, Action(state))
	if err != nil {
		t.Errorf("expected no error, got %+v", err)
	}

	if string(resp.State) != newState {
		t.Errorf("expected state %s, got %s", state, resp.State)
	}
}

func TestUpdatePipeline(t *testing.T) {
	var pipelineUpdate UpdatePipelineInput
	var pipeline = generatePipeline("", "")

	pipelineUpdate.Name = pipeline.Name

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if want, got := fmt.Sprintf("%s/%s", pipelinesBasePath, pipeline.Name), req.URL.Path; want != got {
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
		if err := json.NewEncoder(w).Encode(pipeline); err != nil {
			t.Errorf("expected no error, got %+v", err)
		}
	}))
	// Close the server when test finishes
	defer server.Close()

	c := testClient(server.Client(), server.URL)

	resp, err := c.UpdatePipeline(context.Background(), pipeline.Name, &pipelineUpdate)
	if err != nil {
		t.Errorf("expected no error, got %+v", err)
	}

	if !reflect.DeepEqual(resp, &pipeline) {
		t.Errorf("expected response same as pipeline")
	}
}

func TestGetPipelines(t *testing.T) {
	pBase := generatePipeline("without-env", "")
	pWithEnv := generatePipelineWithEnvironment("with-env")

	pipelines := []*Pipeline{&pBase, &pWithEnv}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if want, got := pipelinesBasePath, req.URL.Path; want != got {
			t.Fatalf("Path mismatched: want=%v got=%v", want, got)
		}

		if err := json.NewEncoder(w).Encode(pipelines); err != nil {
			t.Fatal(err)
		}

	}))
	defer server.Close()

	c := testClient(server.Client(), server.URL)

	gotEnv, err := c.ListPipelines(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if want, got := pipelines, gotEnv; !reflect.DeepEqual(want, got) {
		t.Fatalf("Pipelines mismatched: want=%v got=%v", want, got)
	}
}

func TestGetPipeline(t *testing.T) {
	p := generatePipelineWithEnvironment("my-pipeline-with-env")

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		path := fmt.Sprintf("%s?name=%s", pipelinesBasePath, p.Name)
		if req.RequestURI != path {
			t.Fatalf("Path mismatched: want=%v got=%v", path, req.RequestURI)
		}

		if err := json.NewEncoder(w).Encode(p); err != nil {
			t.Fatal(err)
		}
	}))
	defer server.Close()

	c := testClient(server.Client(), server.URL)

	resp, err := c.GetPipelineByName(context.Background(), p.Name)

	if err != nil {
		t.Errorf("expected no error, got %+v", err)
	}

	if p.Name != resp.Name {
		t.Errorf("expected name %q, got %q", p.Name, resp.Name)
	}

	if !reflect.DeepEqual(p.Environment, resp.Environment) {
		t.Errorf("expected environment %v, got %v", p.Environment, resp.Environment)
	}
}

func TestCreatePipeline(t *testing.T) {
	p := generatePipelineWithEnvironment("my-pipeline-with-env")

	pi := CreatePipelineInput{
		Name:        p.Name,
		Environment: &EntityIdentifier{Name: p.Environment.Name},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if want, got := pipelinesBasePath, req.URL.Path; want != got {
			t.Fatalf("mismatched of request path: want=%s got=%s", want, got)
		}

		var cPi *CreatePipelineInput
		if err := json.NewDecoder(req.Body).Decode(&cPi); err != nil {
			t.Errorf("expected no error, got %+v", err)
		}
		defer req.Body.Close()

		if pi.Name != cPi.Name {
			t.Errorf("expected name %q, got %q", pi.Name, cPi.Name)
		}

		if !reflect.DeepEqual(pi.Environment, cPi.Environment) {
			t.Errorf("expected same environment")
		}

		rP := generatePipeline(pi.Name, "")
		rP.Environment = pi.Environment
		rP.Environment.UUID = null.StringFrom("067fc522-7f3c-4c71-8749-68f3698c2c68")

		// Return response to satisfy client and test response
		if err := json.NewEncoder(w).Encode(rP); err != nil {
			t.Errorf("expected no error, got %+v", err)
		}
	}))
	// Close the server when test finishes
	defer server.Close()

	c := testClient(server.Client(), server.URL)

	resp, err := c.CreatePipeline(context.Background(), &pi)

	if err != nil {
		t.Errorf("expected no error, got %+v", err)
	}

	if pi.Name != resp.Name {
		t.Errorf("expected name %q, got %q", pi.Name, resp.Name)
	}

	if !reflect.DeepEqual(pi.Environment, resp.Environment) {
		t.Errorf("expected environment %v, got %v", pi.Environment, resp.Environment)
	}
}

func generatePipeline(name string, state string) Pipeline {
	if name == "" {
		name = "test"
	}

	if state == "" {
		state = "healthy"
	}

	return Pipeline{
		Name:  name,
		State: PipelineState(state),
	}
}

func generatePipelineWithEnvironment(name string) Pipeline {
	p := generatePipeline(name, "")

	p.Environment = &EntityIdentifier{
		UUID: null.StringFrom("9c73bbc5-75c2-400d-a270-d8aefe727c15"),
		Name: null.StringFrom("my-environment"),
	}
	return p
}
