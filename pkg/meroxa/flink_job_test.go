package meroxa

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestCreateFlinkJob(t *testing.T) {
	name := "test"
	jarURL := "https://hell.lo"
	input := CreateFlinkJobInput{
		Name:   name,
		JarURL: jarURL,
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// Test request
		var cr CreateFlinkJobInput
		err := json.NewDecoder(req.Body).Decode(&cr)
		if err != nil {
			t.Errorf("expected no error, got %+v", err)
		}

		if cr.Name != input.Name {
			t.Errorf("expected name %q, got %q", input.Name, cr.Name)
		}
		if cr.JarURL != input.JarURL {
			t.Errorf("expected jar URL %q, got %q", input.JarURL, cr.JarURL)
		}

		defer req.Body.Close()

		// Return response to satisfy client and test response
		c := generateFlinkJob(name)
		if err := json.NewEncoder(w).Encode(c); err != nil {
			t.Errorf("expected no error, got %+v", err)
		}
	}))
	// Close the server when test finishes
	defer server.Close()

	c := testClient(testRequester(server.Client(), server.URL))

	resp, err := c.CreateFlinkJob(context.Background(), &input)

	if err != nil {
		t.Errorf("expected no error, got %+v", err)
	}

	if resp.Name != input.Name {
		t.Errorf("expected name %q, got %q", input.Name, resp.Name)
	}
}

func generateFlinkJob(name string) *FlinkJob {
	return &FlinkJob{Name: name}
}

func TestListFlinkJobs(t *testing.T) {
	flinkJobs := []*FlinkJob{generateFlinkJob("test1"), generateFlinkJob("test2")}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if want, got := flinkJobsBasePath, req.URL.Path; want != got {
			t.Fatalf("mismatched of request path: want=%s got=%s", want, got)
		}

		defer req.Body.Close()

		// Return response to satisfy client and test response
		if err := json.NewEncoder(w).Encode(flinkJobs); err != nil {
			t.Errorf("expected no error, got %+v", err)
		}
	}))
	// Close the server when test finishes
	defer server.Close()

	c := testClient(testRequester(server.Client(), server.URL))

	resp, err := c.ListFlinkJobs(context.Background())
	if err != nil {
		t.Errorf("expected no error, got %+v", err)
	}

	if !reflect.DeepEqual(resp, flinkJobs) {
		t.Errorf("expected response not same as flink job \n%v\n%v", resp, flinkJobs)
	}
}

func TestGetFlinkJob(t *testing.T) {
	name := "test"
	flinkJob := generateFlinkJob(name)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if want, got := fmt.Sprintf("%s/%s", flinkJobsBasePath, name), req.URL.Path; want != got {
			t.Fatalf("mismatched of request path: want=%s got=%s", want, got)
		}

		defer req.Body.Close()

		// Return response to satisfy client and test response
		if err := json.NewEncoder(w).Encode(flinkJob); err != nil {
			t.Errorf("expected no error, got %+v", err)
		}
	}))
	// Close the server when test finishes
	defer server.Close()

	c := testClient(testRequester(server.Client(), server.URL))

	resp, err := c.GetFlinkJob(context.Background(), name)
	if err != nil {
		t.Errorf("expected no error, got %+v", err)
	}

	if !reflect.DeepEqual(resp, flinkJob) {
		t.Errorf("expected response not same as flink job")
	}
}

func TestDeleteFlinkJob(t *testing.T) {
	name := "test"
	flinkJob := generateFlinkJob(name)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if want, got := fmt.Sprintf("%s/%s", flinkJobsBasePath, name), req.URL.Path; want != got {
			t.Fatalf("mismatched of request path: want=%s got=%s", want, got)
		}

		defer req.Body.Close()

		// Return response to satisfy client and test response
		if err := json.NewEncoder(w).Encode(flinkJob); err != nil {
			t.Errorf("expected no error, got %+v", err)
		}
	}))
	// Close the server when test finishes
	defer server.Close()

	c := testClient(testRequester(server.Client(), server.URL))

	err := c.DeleteFlinkJob(context.Background(), name)
	if err != nil {
		t.Errorf("expected no error, got %+v", err)
	}
}
