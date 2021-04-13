package meroxa

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestEncodeURLCreds(t *testing.T) {
	tests := []struct {
		in   string
		want string
		err  error
	}{
		{"s3://KAHDKJKSA:askkshe+skje/fhds@us-east-1/bucket", "s3://KAHDKJKSA:askkshe+skje%2Ffhds@us-east-1/bucket", nil},
		{"s3://KAHDKJKSA:secretsecret@us-east-1/bucket", "s3://KAHDKJKSA:secretsecret@us-east-1/bucket", nil},
		{"s3://us-east-1/bucket", "s3://us-east-1/bucket", nil},
		{"s3://:apassword@us-east-1/bucket", "s3://:apassword@us-east-1/bucket", nil},
		{"not a URL", "", ErrMissingScheme},
	}

	for _, tt := range tests {
		got, err := encodeURLCreds(tt.in)
		if err != tt.err {
			t.Errorf("expected %+v, got %+v", tt.err, err)
		}
		if got != tt.want {
			t.Errorf("expected %+v, got %+v", tt.want, got)
		}
	}
}

func TestUpdateResource(t *testing.T) {
	var resource UpdateResourceInput

	resource.Name = "resource-name"
	resource.URL = "http://foo.com"
	resource.Metadata = map[string]interface{}{
		"key": "value",
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if want, got := fmt.Sprintf("%s/%s", ResourcesBasePath, resource.Name), req.URL.Path; want != got {
			t.Fatalf("mismatched of request path: want=%s got=%s", want, got)
		}

		var rr *UpdateResourceInput
		if err := json.NewDecoder(req.Body).Decode(&rr); err != nil {
			t.Errorf("expected no error, got %+v", err)
		}
		defer req.Body.Close()

		if rr.URL != resource.URL {
			t.Errorf("expected URL %s, got %s", resource.URL, rr.URL)
		}

		if !reflect.DeepEqual(rr.Metadata, resource.Metadata) {
			t.Errorf("expected same metadata")
		}

		// Return response to satisfy client and test response
		c := generateResource(resource.Name, 0, "", nil)
		c.URL = resource.URL
		c.Metadata = resource.Metadata
		json.NewEncoder(w).Encode(c)
	}))

	// Close the server when test finishes
	defer server.Close()

	c := testClient(server.Client(), server.URL)

	resp, err := c.UpdateResource(context.Background(), resource.Name, resource)
	if err != nil {
		t.Errorf("expected no error, got %+v", err)
	}

	if resp.URL != resource.URL {
		t.Errorf("expected url %s, got %s", resource.URL, resp.URL)
	}
}

func generateResource(name string, id int, url string, metadata map[string]interface{}) Resource {
	if name == "" {
		name = "test"
	}

	if id == 0 {
		id = rand.Intn(10000)
	}

	return Resource{
		ID:       id,
		Type:     "postgres",
		Name:     name,
		URL:      url,
		Metadata: metadata,
	}
}
