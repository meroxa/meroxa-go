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
	tests := []struct {
		desc            string
		input           func(string, string, string, string) CreateDeploymentInput
		deployment      func(string, string, string, string) Deployment
		withEnvironment bool
	}{
		{
			desc: "Deployment without environment",
			input: func(appName, gitSha, specVersion, _ string) CreateDeploymentInput {
				return CreateDeploymentInput{
					Application: EntityIdentifier{Name: appName},
					GitSha:      gitSha,
					SpecVersion: specVersion,
					Spec:        map[string]interface{}{},
				}
			},
			deployment: func(appName, gitSha, specVersion, _ string) Deployment {
				return generateDeployment(appName, gitSha, specVersion)
			},
			withEnvironment: false,
		},
		{
			desc: "Deployment with environment",
			input: func(appName, gitSha, specVersion, environmentName string) CreateDeploymentInput {
				return CreateDeploymentInput{
					Application: EntityIdentifier{Name: appName},
					GitSha:      gitSha,
					SpecVersion: specVersion,
					Spec:        map[string]interface{}{},
					Environment: EntityIdentifier{Name: environmentName},
				}
			},
			deployment: func(appName, gitSha, specVersion, envName string) Deployment {
				return generateDeploymentWithEnv(appName, gitSha, specVersion, envName)
			},
			withEnvironment: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			var input CreateDeploymentInput

			if tc.withEnvironment {
				input = tc.input("app-name", "fc89d80e-2dea-4dab-9a34-3095aa2a8958", "latest", "self-hosted")
			} else {
				input = tc.input("app-name", "fc89d80e-2dea-4dab-9a34-3095aa2a8958", "latest", "")
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

				if cr.Environment != input.Environment {
					t.Errorf("expected environment %v, got %v", input.Environment, cr.Environment)
				}

				defer req.Body.Close()

				// Return response to satisfy client and test response
				var d Deployment
				if tc.withEnvironment {
					d = tc.deployment(input.Application.Name, input.GitSha, input.SpecVersion, input.Environment.Name)
				} else {
					d = tc.deployment(input.Application.Name, input.GitSha, input.SpecVersion, "")
				}

				if err := json.NewEncoder(w).Encode(d); err != nil {
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
		})
	}
}

func generateAppDeploymentWithEnv(environmentName string) ApplicationDeployment {
	return ApplicationDeployment{
		EntityIdentifier: EntityIdentifier{
			UUID: uuid.NewString(),
		},
		Environment: EntityIdentifier{
			Name: environmentName,
		},
	}
}

func generateDeployment(appName, gitSha, specVersion string) Deployment {
	return Deployment{
		UUID:        uuid.NewString(),
		GitSha:      gitSha,
		Application: EntityIdentifier{Name: appName},
		SpecVersion: specVersion,
		Status:      DeploymentStatus{State: DeploymentStateDeploying},
	}
}

func generateDeploymentWithEnv(appName, gitSha, specVersion, envName string) Deployment {
	deployment := generateDeployment(appName, gitSha, specVersion)
	deployment.Environment = EntityIdentifier{Name: envName}
	return deployment
}

func TestGetLatestDeployment(t *testing.T) {
	appName := "test"
	gitSha := "9d856c23-82ad-4f01-8901-b1fe09c800e6"
	envName := "self-hosted"
	deployment := generateDeploymentWithEnv(appName, gitSha, "latest", envName)

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
	gitSha := "6adb1fc3-9cdb-4f56-a60b-98e0f96a63dd"
	envName := "self-hosted"
	deployment := generateDeploymentWithEnv(appName, gitSha, "latest", envName)

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
