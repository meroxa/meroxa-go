package meroxa

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCreateFunction(t *testing.T) {
	input := &CreateFunctionInput{
		Name:         "my_func",
		InputStream:  "input_stream",
		OutputStream: "output_stream",
		PipelineName: "pipeline_name",
		Image:        "meroxa/image",
		Command:      []string{"echo", "hello"},
		Args:         []string{"arg"},
		EnvVars:      map[string]string{"key": "val"},
	}
	output := &Function{
		UUID:         "1234",
		Name:         "my_func",
		InputStream:  "input_stream",
		OutputStream: "output_stream",
		Image:        "meroxa/image",
		Command:      []string{"echo", "hello"},
		Args:         []string{"arg"},
		EnvVars:      map[string]string{"key": "val"},
		Status: FunctionStatus{
			State:   "RUNNING",
			Details: "Details",
		},
		Pipeline: PipelineIdentifier{
			Name: "my_pipeline",
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		var i *CreateFunctionInput
		if err := json.NewDecoder(req.Body).Decode(&i); err != nil {
			t.Fatalf("expected no error, got %+v", err)
		}

		if diff := cmp.Diff(input, i); diff != "" {
			t.Fatalf("mismatch of function input (-want +got): %s", diff)
		}

		json.NewEncoder(w).Encode(output)
	}))
	defer server.Close()

	c := testClient(server.Client(), server.URL)

	gotOutput, err := c.CreateFunction(context.Background(), input)
	if err != nil {
		t.Fatalf("expected no error, got %+v", err)
	}

	if diff := cmp.Diff(output, gotOutput); diff != "" {
		t.Fatalf("mismatch of function output (-want +got): %s", diff)
	}
}
