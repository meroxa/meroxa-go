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

func TestCreateApplication(t *testing.T) {
	input := CreateApplicationInput{
		Name:     "test",
		Language: "golang",
		GitSha:   "abc",
		Pipeline: EntityIdentifier{UUID: "def"},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// Test request
		type connectionRequest struct {
			Name string `json:"name"`
		}

		var cr connectionRequest
		err := json.NewDecoder(req.Body).Decode(&cr)
		if err != nil {
			t.Errorf("expected no error, got %+v", err)
		}

		if cr.Name != input.Name {
			t.Errorf("expected name %s, got %s", input.Name, cr.Name)
		}

		defer req.Body.Close()

		// Return response to satisfy client and test response
		c := generateApplication(input.Name)
		if err := json.NewEncoder(w).Encode(c); err != nil {
			t.Errorf("expected no error, got %+v", err)
		}
	}))
	// Close the server when test finishes
	defer server.Close()

	c := testClient(testRequester(server.Client(), server.URL))

	resp, err := c.CreateApplication(context.Background(), &input)

	if err != nil {
		t.Errorf("expected no error, got %+v", err)
	}

	if resp.Name != input.Name {
		t.Errorf("expected name %s, got %s", input.Name, resp.Name)
	}
}

func generateApplication(name string) Application {
	if name == "" {
		name = "test"
	}

	return Application{Name: name, UUID: uuid.NewString(), Language: "golang", GitSha: "abc", Status: ApplicationStatus{State: ApplicationStateRunning}}
}

func generateApplicationWithEnvironment(name string) Application {
	a := Application{Name: name, UUID: uuid.NewString(), Language: "golang", GitSha: "abc", Status: ApplicationStatus{State: ApplicationStateRunning}}
	a.Environment = generateApplicationEnvironment("private", "env-1234", "aws")
	return a
}

func TestCreateApplicationV2(t *testing.T) {

	tests := []struct {
		desc  string
		input func() CreateApplicationInput
	}{
		{
			desc: "An application without an environment",
			input: func() CreateApplicationInput {
				return CreateApplicationInput{
					Name:     "test",
					Language: "golang",
					GitSha:   "abc",
				}
			},
		},
		{
			desc: "An application with an environment",
			input: func() CreateApplicationInput {
				return CreateApplicationInput{
					Name:        "test",
					Language:    "golang",
					GitSha:      "abc",
					Environment: EntityIdentifier{Name: "my-env"},
				}
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			input := tc.input()
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				// Test request
				type connectionRequest struct {
					Name string `json:"name"`
				}

				var cr connectionRequest
				err := json.NewDecoder(req.Body).Decode(&cr)
				if err != nil {
					t.Errorf("expected no error, got %+v", err)
				}

				if cr.Name != input.Name {
					t.Errorf("expected name %s, got %s", input.Name, cr.Name)
				}

				defer req.Body.Close()

				// Return response to satisfy client and test response
				c := generateApplication(input.Name)
				if err := json.NewEncoder(w).Encode(c); err != nil {
					t.Errorf("expected no error, got %+v", err)
				}
			}))
			// Close the server when test finishes
			defer server.Close()

			c := testClient(testRequester(server.Client(), server.URL))

			resp, err := c.CreateApplicationV2(context.Background(), &input)

			if err != nil {
				t.Errorf("expected no error, got %+v", err)
			}

			if resp.Name != input.Name {
				t.Errorf("expected name %s, got %s", input.Name, resp.Name)
			}
		})
	}
}

func TestGetApplicationByName(t *testing.T) {
	name := "my-app"
	app := generateApplication(name)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if want, got := fmt.Sprintf("%s/%s", applicationsBasePath, name), req.URL.Path; want != got {
			t.Fatalf("mismatched of request path: want=%s got=%s", want, got)
		}

		defer req.Body.Close()

		// Return response to satisfy client and test response
		if err := json.NewEncoder(w).Encode(app); err != nil {
			t.Errorf("expected no error, got %+v", err)
		}
	}))
	// Close the server when test finishes
	defer server.Close()

	c := testClient(testRequester(server.Client(), server.URL))

	resp, err := c.GetApplication(context.Background(), app.Name)
	if err != nil {
		t.Errorf("expected no error, got %+v", err)
	}

	if !reflect.DeepEqual(resp, &app) {
		t.Errorf("expected response same as application")
	}
}

func TestGetApplicationByUUID(t *testing.T) {
	app := generateApplicationWithEnvironment("")
	app.Functions = make([]EntityDetails, 0)
	app.Functions = append(app.Functions, EntityDetails{EntityIdentifier: EntityIdentifier{Name: "fun1"}})
	app.Connectors = make([]EntityDetails, 0)
	app.Connectors = append(
		app.Connectors,
		EntityDetails{EntityIdentifier: EntityIdentifier{Name: "conn1"}},
		EntityDetails{EntityIdentifier: EntityIdentifier{Name: "conn2"}})
	app.Resources = make([]ApplicationResource, 0)
	app.Resources = append(
		app.Resources,
		ApplicationResource{
			EntityIdentifier: EntityIdentifier{
				Name: "resource1",
			},
			Collection: ResourceCollection{
				Name:   "table",
				Source: "true",
			},
		},
		ApplicationResource{
			EntityIdentifier: EntityIdentifier{
				Name: "resource1",
			},
			Collection: ResourceCollection{
				Name:        "table_out",
				Destination: "true",
			},
		})

	app.Environment = generateApplicationEnvironment("private", "aws", "env-1234")

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if want, got := fmt.Sprintf("%s/%s", applicationsBasePath, app.UUID), req.URL.Path; want != got {
			t.Fatalf("mismatched of request path: want=%s got=%s", want, got)
		}

		defer req.Body.Close()

		// Return response to satisfy client and test response
		if err := json.NewEncoder(w).Encode(app); err != nil {
			t.Errorf("expected no error, got %+v", err)
		}
	}))
	// Close the server when test finishes
	defer server.Close()

	c := testClient(testRequester(server.Client(), server.URL))

	resp, err := c.GetApplication(context.Background(), fmt.Sprint(app.UUID))
	if err != nil {
		t.Errorf("expected no error, got %+v", err)
	}

	if !reflect.DeepEqual(resp, &app) {
		t.Errorf("expected response same as application")
	}
}

func TestListApplications(t *testing.T) {
	a1 := generateApplication("app1")
	a2 := generateApplication("app2")
	a3 := generateApplicationWithEnvironment("app3")

	list := []*Application{&a1, &a2, &a3}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if want, got := applicationsBasePath, req.URL.Path; want != got {
			t.Fatalf("mismatched of request path: want=%s got=%s", want, got)
		}

		defer req.Body.Close()

		// Return response to satisfy client and test response
		if err := json.NewEncoder(w).Encode(list); err != nil {
			t.Errorf("expected no error, got %+v", err)
		}
	}))
	// Close the server when test finishes
	defer server.Close()

	c := testClient(testRequester(server.Client(), server.URL))

	resp, err := c.ListApplications(context.Background())
	if err != nil {
		t.Errorf("expected no error, got %+v", err)
	}

	if !reflect.DeepEqual(resp, list) {
		t.Errorf("expected response same as list")
	}
}

func TestDeleteApplication(t *testing.T) {
	app := generateApplication("another-app")

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if want, got := fmt.Sprintf("%s/%s", applicationsBasePath, app.UUID), req.URL.Path; want != got {
			t.Fatalf("mismatched of request path: want=%s got=%s", want, got)
		}

		defer req.Body.Close()

		if err := json.NewEncoder(w).Encode(app); err != nil {
			t.Errorf("expected no error, got %+v", err)
		}
	}))
	defer server.Close()

	c := testClient(testRequester(server.Client(), server.URL))

	err := c.DeleteApplication(context.Background(), app.UUID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteApplicationEntitiesWithAppNotFound(t *testing.T) {
	appName := "test"
	pipelineName := fmt.Sprintf("turbine-pipeline-%s", appName)
	pipeline := generatePipeline(pipelineName, "")

	connectorSrc := generateConnector("src-connector", nil, nil)
	connectorSrc.PipelineName = pipeline.Name

	function := generateFunction()
	function.Pipeline = PipelineIdentifier{Name: pipeline.Name}

	mux := http.NewServeMux()

	mux.HandleFunc(fmt.Sprintf("%s/%s", applicationsBasePath, appName), func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusNotFound)
		defer req.Body.Close()
	})

	mux.HandleFunc(pipelinesBasePath, func(res http.ResponseWriter, req *http.Request) {
		if want, got := fmt.Sprintf("name=%s", pipeline.Name), req.URL.RawQuery; want != got {
			t.Fatalf("mismatched of request query parameter: want=%s got=%s", want, got)
		}
		if err := json.NewEncoder(res).Encode(pipeline); err != nil {
			t.Errorf("expected no error, got %+v", err)
		}
		defer req.Body.Close()
	})
	mux.HandleFunc(fmt.Sprintf("%s/%s/connectors", pipelinesBasePath, pipelineName), func(res http.ResponseWriter, req *http.Request) {
		list := []*Connector{&connectorSrc}
		if err := json.NewEncoder(res).Encode(list); err != nil {
			t.Errorf("expected no error, got %+v", err)
		}
	})
	mux.HandleFunc(functionsBasePath, func(res http.ResponseWriter, req *http.Request) {
		list := []*Function{function}
		if err := json.NewEncoder(res).Encode(list); err != nil {
			t.Errorf("expected no error, got %+v", err)
		}
	})
	mux.HandleFunc(fmt.Sprintf("%s/%s", connectorsBasePath, connectorSrc.Name), func(res http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodDelete {
			res.WriteHeader(http.StatusOK)
		}
		defer req.Body.Close()
	})
	mux.HandleFunc(fmt.Sprintf("%s/%s", functionsBasePath, function.Name), func(res http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodDelete {
			res.WriteHeader(http.StatusOK)
		}
		defer req.Body.Close()
	})
	mux.HandleFunc(fmt.Sprintf("%s/%s", pipelinesBasePath, pipeline.Name), func(res http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodDelete {
			res.WriteHeader(http.StatusOK)
		}
		defer req.Body.Close()
	})

	server := httptest.NewServer(mux)
	defer server.Close()

	c := testClient(testRequester(server.Client(), server.URL))

	resp, err := c.DeleteApplicationEntities(context.Background(), appName)
	if err != nil {
		t.Errorf("expected no error, got %+v", err)
	}

	if resp.StatusCode != http.StatusNoContent {
		t.Errorf("got %v, expected %v", resp.StatusCode, http.StatusNoContent)
	}
}

func TestDeleteApplicationEntitiesWithAppFound(t *testing.T) {
	app := generateApplication("another-app")

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if want, got := fmt.Sprintf("%s/%s", applicationsBasePath, app.Name), req.URL.Path; want != got {
			t.Fatalf("mismatched of request path: want=%s got=%s", want, got)
		}

		defer req.Body.Close()

		if err := json.NewEncoder(w).Encode(app); err != nil {
			t.Errorf("expected no error, got %+v", err)
		}
	}))
	defer server.Close()

	c := testClient(testRequester(server.Client(), server.URL))

	resp, err := c.DeleteApplicationEntities(context.Background(), app.Name)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("got %v, expected %v", resp.StatusCode, http.StatusOK)
	}
}

func TestGetApplicationLogs(t *testing.T) {
	name := "my-app"
	appLogs := ApplicationLogs{
		FunctionLogs: map[string]string{
			"func1": "success",
		},
		ConnectorLogs: map[string]string{
			"conn1": "success",
		},
		DeploymentLogs: map[string]string{
			"ab-cd-ef": "success",
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if want, got := fmt.Sprintf("%s/%s/logs", applicationsBasePath, name), req.URL.Path; want != got {
			t.Fatalf("mismatched of request path: want=%s got=%s", want, got)
		}

		defer req.Body.Close()

		// Return response to satisfy client and test response
		if err := json.NewEncoder(w).Encode(appLogs); err != nil {
			t.Errorf("expected no error, got %+v", err)
		}
	}))
	// Close the server when test finishes
	defer server.Close()

	c := testClient(testRequester(server.Client(), server.URL))

	resp, err := c.GetApplicationLogs(context.Background(), name)
	if err != nil {
		t.Errorf("expected no error, got %+v", err)
	}

	if !reflect.DeepEqual(resp, &appLogs) {
		t.Errorf("expected response same as application logs")
	}
}
