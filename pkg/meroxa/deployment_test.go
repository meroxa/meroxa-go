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
	"github.com/volatiletech/null/v8"
)

func TestCreateDeployment(t *testing.T) {
	appName := "test"
	specVersion := "latest"
	gitSha := "abc"
	input := CreateDeploymentInput{
		GitSha:      gitSha,
		Application: EntityIdentifier{Name: null.StringFrom(appName)},
		SpecVersion: null.StringFrom(specVersion),
		Spec:        null.StringFrom("{}"),
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
		if cr.Spec != input.Spec {
			t.Errorf("expected spec %s, got %s", input.Spec.String, cr.Spec.String)
		}
		if cr.SpecVersion != input.SpecVersion {
			t.Errorf("expected spec_version %s, got %s", input.SpecVersion.String, cr.SpecVersion.String)
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
		Application: EntityIdentifier{Name: null.StringFrom(appName)},
		SpecVersion: null.StringFrom(specVersion),
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
