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

func TestGetEnvironments(t *testing.T) {
	env := &Environment{
		UUID:     "1234",
		Name:     "environment-1234",
		Provider: "aws",
		Region:   "us-east",
		Type:     "private",
		Status:   EnvironmentViewStatus{State: "provisioned"},
	}

	environments := []*Environment{env}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if want, got := environmentsBasePath, req.URL.Path; want != got {
			t.Fatalf("Path mismatched: want=%v got=%v", want, got)
		}

		if err := json.NewEncoder(w).Encode(environments); err != nil {
			t.Fatal(err)
		}

	}))
	defer server.Close()

	c := testClient(testRequester(server.Client(), server.URL))

	gotEnv, err := c.ListEnvironments(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if want, got := environments, gotEnv; !reflect.DeepEqual(want, got) {
		t.Fatalf("Environments mismatched: want=%v got=%v", want, got)
	}
}

func TestCreateEnvironment(t *testing.T) {
	environment := &CreateEnvironmentInput{}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if want, got := environmentsBasePath, req.URL.Path; want != got {
			t.Fatalf("mismatched of request path: want=%s got=%s", want, got)
		}

		var ee *CreateEnvironmentInput
		if err := json.NewDecoder(req.Body).Decode(&ee); err != nil {
			t.Errorf("expected no error, got %+v", err)
		}
		defer req.Body.Close()

		if environment.Type != ee.Type {
			t.Errorf("expected type %q, got %q", ee.Type, environment.Type)
		}

		if environment.Provider != ee.Provider {
			t.Errorf("expected provider %q, got %q", ee.Provider, environment.Provider)
		}

		if environment.Name != ee.Name {
			t.Errorf("expected name %q, got %q", ee.Name, environment.Name)
		}

		if environment.Region != ee.Region {
			t.Errorf("expected region %q, got %q", ee.Region, environment.Region)
		}

		if !reflect.DeepEqual(ee.Configuration, environment.Configuration) {
			t.Errorf("expected same configuration")
		}

		// Return response to satisfy client and test response
		c := generateEnvironment(environment.Type, environment.Provider, environment.Name, EnvironmentViewStatus{State: "private"})
		if err := json.NewEncoder(w).Encode(c); err != nil {
			t.Errorf("expected no error, got %+v", err)
		}
	}))
	// Close the server when test finishes
	defer server.Close()

	c := testClient(testRequester(server.Client(), server.URL))

	resp, err := c.CreateEnvironment(context.Background(), environment)

	if err != nil {
		t.Errorf("expected no error, got %+v", err)
	}

	if environment.Type != resp.Type {
		t.Errorf("expected type %q, got %q", resp.Type, environment.Type)
	}

	if environment.Provider != resp.Provider {
		t.Errorf("expected provider %q, got %q", resp.Provider, environment.Provider)
	}

	if environment.Name != resp.Name {
		t.Errorf("expected name %q, got %q", resp.Name, environment.Name)
	}
}

func TestCreateBadEnvironment(t *testing.T) {
	environment := &CreateEnvironmentInput{Type: "self_hosted",
		Name:     "badenv",
		Provider: "aws",
		Region:   "us-east-1",
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if want, got := environmentsBasePath, req.URL.Path; want != got {
			t.Fatalf("mismatched of request path: want=%s got=%s", want, got)
		}

		var ee *CreateEnvironmentInput
		if err := json.NewDecoder(req.Body).Decode(&ee); err != nil {
			t.Errorf("expected no error, got %+v", err)
		}
		defer req.Body.Close()

		if environment.Type != ee.Type {
			t.Errorf("expected type %q, got %q", ee.Type, environment.Type)
		}

		if environment.Provider != ee.Provider {
			t.Errorf("expected provider %q, got %q", ee.Provider, environment.Provider)
		}

		if environment.Name != ee.Name {
			t.Errorf("expected name %q, got %q", ee.Name, environment.Name)
		}

		if environment.Region != ee.Region {
			t.Errorf("expected region %q, got %q", ee.Region, environment.Region)
		}

		if !reflect.DeepEqual(ee.Configuration, environment.Configuration) {
			t.Errorf("expected same configuration")
		}

		// Return response to satisfy client and test response
		c := generateEnvironment(environment.Type, environment.Provider, environment.Name, EnvironmentViewStatus{State: "preflight_error", PreflightDetails: &PreflightDetails{PreflightPermissions: &PreflightPermissions{S3: []string{"missing read permission for S3", "missing write permissions for S3"}}}})
		if err := json.NewEncoder(w).Encode(c); err != nil {
			t.Errorf("expected no error, got %+v", err)
		}
	}))
	// Close the server when test finishes
	defer server.Close()

	c := testClient(testRequester(server.Client(), server.URL))

	resp, err := c.CreateEnvironment(context.Background(), environment)

	if err != nil {
		t.Errorf("expected no error, got %+v", err)
	}

	if environment.Type != resp.Type {
		t.Errorf("expected type %q, got %q", resp.Type, environment.Type)
	}

	if environment.Provider != resp.Provider {
		t.Errorf("expected provider %q, got %q", resp.Provider, environment.Provider)
	}

	if environment.Name != resp.Name {
		t.Errorf("expected name %q, got %q", resp.Name, environment.Name)
	}
}
func TestGetEnvironment(t *testing.T) {
	env := generateEnvironment("private", "environment-1234", "aws", EnvironmentViewStatus{State: "private"})

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		path := fmt.Sprintf("%s/%s", environmentsBasePath, env.UUID)
		if req.URL.Path != path {
			t.Fatalf("Path mismatched: want=%v got=%v", path, req.URL.Path)
		}

		if err := json.NewEncoder(w).Encode(env); err != nil {
			t.Fatal(err)
		}
	}))
	defer server.Close()

	c := testClient(testRequester(server.Client(), server.URL))

	resp, err := c.GetEnvironment(context.Background(), env.UUID)

	if err != nil {
		t.Errorf("expected no error, got %+v", err)
	}

	if env.Type != resp.Type {
		t.Errorf("expected type %q, got %q", resp.Type, env.Type)
	}

	if env.Provider != resp.Provider {
		t.Errorf("expected provider %q, got %q", resp.Provider, env.Provider)
	}

	if env.Name != resp.Name {
		t.Errorf("expected name %q, got %q", resp.Name, env.Name)
	}
}

func TestDeleteEnvironment(t *testing.T) {
	env := generateEnvironment("private", "environment-1234", "aws", EnvironmentViewStatus{State: "private"})
	deprovisioningState := "deprovisioning"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if want, got := fmt.Sprintf("%s/%s", environmentsBasePath, env.UUID), req.URL.Path; want != got {
			t.Fatalf("mismatched of request path: want=%s got=%s", want, got)
		}

		defer req.Body.Close()

		env.Status.State = EnvironmentState(deprovisioningState)
		if err := json.NewEncoder(w).Encode(env); err != nil {
			t.Errorf("expected no error, got %+v", err)
		}
	}))
	defer server.Close()

	c := testClient(testRequester(server.Client(), server.URL))

	resp, err := c.DeleteEnvironment(context.Background(), env.UUID)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(resp, &env) {
		t.Errorf("expected response same as environment")
	}

	if string(resp.Status.State) != deprovisioningState {
		t.Errorf("expected state %q, got %s", deprovisioningState, resp.Status.State)
	}
}

func TestUpdateEnvironment(t *testing.T) {
	env := generateEnvironment("private", "environment-1234", "aws", EnvironmentViewStatus{State: "private"})
	updatedName := "new-name"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if want, got := fmt.Sprintf("%s/%s", environmentsBasePath, env.UUID), req.URL.Path; want != got {
			t.Fatalf("mismatched of request path: want=%s got=%s", want, got)
		}

		defer req.Body.Close()

		env.Name = updatedName
		if err := json.NewEncoder(w).Encode(env); err != nil {
			t.Errorf("expected no error, got %+v", err)
		}
	}))
	defer server.Close()

	c := testClient(testRequester(server.Client(), server.URL))

	resp, err := c.UpdateEnvironment(context.Background(), env.UUID, &UpdateEnvironmentInput{Name: updatedName})
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(resp, &env) {
		t.Errorf("expected response to be updated environment")
	}
}

func TestRepairEnvironment(t *testing.T) {
	env := generateEnvironment("private", "environment-awaiting-repair", "aws", EnvironmentViewStatus{State: "private"})

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if expected, actual := fmt.Sprintf("%s/%s/%s", environmentsBasePath, env.UUID, "actions"), req.URL.Path; expected != actual {
			t.Fatalf("mismatched of request path: expected=%s actual=%s", expected, actual)
		}

		expectedActionName := EnvironmentActionRepair
		var re *RepairEnvironmentInput

		if err := json.NewDecoder(req.Body).Decode(&re); err != nil {
			t.Errorf("expected no error, got %+v", err)
		}

		if re.Action != expectedActionName {
			t.Errorf("expected action parameter %q, got %q", expectedActionName, re.Action)
		}

		defer req.Body.Close()

		env.Status.State = "repairing"
		if err := json.NewEncoder(w).Encode(env); err != nil {
			t.Errorf("expected no error, got %+v", err)
		}
	}))
	defer server.Close()

	c := testClient(testRequester(server.Client(), server.URL))

	resp, err := c.PerformActionOnEnvironment(context.Background(), env.UUID, &RepairEnvironmentInput{Action: EnvironmentActionRepair})
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(resp, &env) {
		t.Errorf("expected response to be environment in repairing state")
	}
}

func generateEnvironment(t EnvironmentType, p EnvironmentProvider, n string, status EnvironmentViewStatus) Environment {
	return Environment{
		Type:     t,
		Name:     n,
		Provider: p,
		Region:   "us-east-1",
		Status:   status,
		UUID:     "1a92d590-d59c-460b-94de-870f04ab35bf",
	}
}
