package meroxa

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/google/uuid"
)

func TestCreateDeployment(t *testing.T) {
	appName := "test"
	specVersion := "latest"
	gitSha := "abc"
	input := CreateDeploymentInput{
		GitSha:      gitSha,
		Application: EntityIdentifier{Name: appName},
		SpecVersion: specVersion,
		Spec:        map[string]interface{}{},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// Test request
		var cr CreateDeploymentInput
		err := json.NewDecoder(req.Body).Decode(&cr)
		if err != nil {
			t.Errorf("expected no error, got %+v", err)
		}

		if cr.GitSha != input.GitSha {
			t.Errorf("expected git_sha %s, got %s", input.GitSha, cr.GitSha)
		}
		if reflect.DeepEqual(cr.Spec, input.Spec) {
			t.Errorf("expected spec %s, got %s", input.Spec, cr.Spec)
		}
		if cr.SpecVersion != input.SpecVersion {
			t.Errorf("expected spec_version %s, got %s", input.SpecVersion, cr.SpecVersion)
		}
		if cr.Application != input.Application {
			t.Errorf("expected application %v, got %v", input.Application, cr.Application)
		}

		defer req.Body.Close()

		// Return response to satisfy client and test response
		c := generateDeployment(appName, gitSha, specVersion)
		if err := json.NewEncoder(w).Encode(c); err != nil {
			t.Errorf("expected no error, got %+v", err)
		}
	}))
	// Close the server when test finishes
	defer server.Close()

	c := testClient(testRequester(server.Client(), server.URL))

	resp, err := c.CreateDeployment(context.Background(), &input)

	if err != nil {
		t.Errorf("expected no error, got %+v", err)
	}

	if resp.GitSha != input.GitSha {
		t.Errorf("expected git_sha %s, got %s", input.GitSha, resp.GitSha)
	}
}

func generateDeployment(appName, gitSha, specVersion string) Deployment {
	return Deployment{
		UUID:        uuid.NewString(),
		GitSha:      gitSha,
		Application: EntityIdentifier{Name: appName},
		SpecVersion: specVersion,
		Status:      DeploymentStatus{State: DeploymentStateDeploying}}
}

func TestGetLatestDeployment(t *testing.T) {
	appName := "test"
	deployment := generateDeployment(appName, "abc", "latest")

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if want, got := fmt.Sprintf("%s/%s/deployments/latest", applicationsBasePath, appName), req.URL.Path; want != got {
			t.Fatalf("mismatched of request path: want=%s got=%s", want, got)
		}

		defer req.Body.Close()

		// Return response to satisfy client and test response
		if err := json.NewEncoder(w).Encode(deployment); err != nil {
			t.Errorf("expected no error, got %+v", err)
		}
	}))
	// Close the server when test finishes
	defer server.Close()

	c := testClient(testRequester(server.Client(), server.URL))

	resp, err := c.GetLatestDeployment(context.Background(), appName)
	if err != nil {
		t.Errorf("expected no error, got %+v", err)
	}

	if !reflect.DeepEqual(resp, &deployment) {
		t.Errorf("expected response not same as deployment")
	}
}

func TestGetDeployment(t *testing.T) {
	appName := "test"
	deployment := generateDeployment(appName, "abc", "latest")

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if want, got := fmt.Sprintf("%s/%s/deployments/%s", applicationsBasePath, appName, deployment.UUID), req.URL.Path; want != got {
			t.Fatalf("mismatched of request path: want=%s got=%s", want, got)
		}

		defer req.Body.Close()

		// Return response to satisfy client and test response
		if err := json.NewEncoder(w).Encode(deployment); err != nil {
			t.Errorf("expected no error, got %+v", err)
		}
	}))
	// Close the server when test finishes
	defer server.Close()

	c := testClient(testRequester(server.Client(), server.URL))

	resp, err := c.GetDeployment(context.Background(), appName, deployment.UUID)
	if err != nil {
		t.Errorf("expected no error, got %+v", err)
	}

	if !reflect.DeepEqual(resp, &deployment) {
		t.Errorf("expected response not same as deployment")
	}
}
