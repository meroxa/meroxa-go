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

func TestCreateApplication(t *testing.T) {
	input := CreateApplicationInput{
		Name:     "test",
		Language: "golang",
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
		json.NewEncoder(w).Encode(c)
	}))
	// Close the server when test finishes
	defer server.Close()

	c := testClient(server.Client(), server.URL)

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

	return Application{Name: name, UUID: uuid.NewString()}
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
		json.NewEncoder(w).Encode(app)
	}))
	// Close the server when test finishes
	defer server.Close()

	c := testClient(server.Client(), server.URL)

	resp, err := c.GetApplication(context.Background(), app.Name)
	if err != nil {
		t.Errorf("expected no error, got %+v", err)
	}

	if !reflect.DeepEqual(resp, &app) {
		t.Errorf("expected response same as application")
	}
}

func TestGetApplicationByUUID(t *testing.T) {
	app := generateApplication("")
	app.Functions = make([]EntityIdentifier, 0)
	app.Functions = append(app.Functions, EntityIdentifier{Name: null.StringFrom("fun1")})
	app.Connectors = make([]EntityIdentifier, 0)
	app.Connectors = append(app.Connectors, EntityIdentifier{Name: null.StringFrom("conn1")})
	app.Resources = make([]EntityIdentifier, 0)
	app.Resources = append(app.Resources, EntityIdentifier{Name: null.StringFrom("resource1")})

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if want, got := fmt.Sprintf("%s/%s", applicationsBasePath, app.UUID), req.URL.Path; want != got {
			t.Fatalf("mismatched of request path: want=%s got=%s", want, got)
		}

		defer req.Body.Close()

		// Return response to satisfy client and test response
		json.NewEncoder(w).Encode(app)
	}))
	// Close the server when test finishes
	defer server.Close()

	c := testClient(server.Client(), server.URL)

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
	list := []*Application{&a1, &a2}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if want, got := fmt.Sprintf("%s", applicationsBasePath), req.URL.Path; want != got {
			t.Fatalf("mismatched of request path: want=%s got=%s", want, got)
		}

		defer req.Body.Close()

		// Return response to satisfy client and test response
		json.NewEncoder(w).Encode(list)
	}))
	// Close the server when test finishes
	defer server.Close()

	c := testClient(server.Client(), server.URL)

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

		json.NewEncoder(w).Encode(app)
	}))
	defer server.Close()

	c := testClient(server.Client(), server.URL)

	err := c.DeleteApplication(context.Background(), app.UUID)
	if err != nil {
		t.Fatal(err)
	}
}
