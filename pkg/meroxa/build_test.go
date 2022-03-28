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

func TestCreateBuild(t *testing.T) {
	input := CreateBuildInput{
		SourceBlob: SourceBlob{
			Url: "https://meroxa-test.url",
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// Test request
		type connectionRequest struct {
			SourceBlob SourceBlob `json:"source_blob"`
		}

		var cr connectionRequest
		err := json.NewDecoder(req.Body).Decode(&cr)
		if err != nil {
			t.Errorf("expected no error, got %+v", err)
		}

		if cr.SourceBlob.Url != input.SourceBlob.Url {
			t.Errorf("expected url %s, got %s", input.SourceBlob.Url, cr.SourceBlob.Url)
		}

		defer req.Body.Close()

		// Return response to satisfy client and test response
		c := generateBuild("", input.SourceBlob.Url)
		json.NewEncoder(w).Encode(c)
	}))
	// Close the server when test finishes
	defer server.Close()

	c := testClient(server.Client(), server.URL)

	resp, err := c.CreateBuild(context.Background(), &input)

	if err != nil {
		t.Errorf("expected no error, got %+v", err)
	}

	if resp.SourceBlob.Url != input.SourceBlob.Url {
		t.Errorf("expected name %s, got %s", input.SourceBlob.Url, resp.SourceBlob.Url)
	}
}

func TestGetBuild(t *testing.T) {
	uuid := "31152ef2-16e0-4c8e-8bd1-9e2181c4974a"
	build := generateBuild(uuid, "")

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if want, got := fmt.Sprintf("%s/%s", buildsBasePath, uuid), req.URL.Path; want != got {
			t.Fatalf("mismatched of request path: want=%s got=%s", want, got)
		}

		defer req.Body.Close()

		// Return response to satisfy client and test response
		json.NewEncoder(w).Encode(build)
	}))
	// Close the server when test finishes
	defer server.Close()

	c := testClient(server.Client(), server.URL)

	resp, err := c.GetBuild(context.Background(), build.Uuid)
	if err != nil {
		t.Errorf("expected no error, got %+v", err)
	}

	if !reflect.DeepEqual(resp, &build) {
		t.Errorf("expected response same as build")
	}
}

func generateBuild(uuid, url string) Build {
	if uuid == "" {
		uuid = "31152ef2-16e0-4c8e-8bd1-9e2181c4974a"
	}

	if url == "" {
		url = "https://meroxa-test.url"
	}
	return Build{Uuid: uuid, SourceBlob: SourceBlob{Url: url}}
}