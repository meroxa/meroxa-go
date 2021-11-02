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
	"time"
)

func Test_Resource_PerformActions(t *testing.T) {
	for k, v := range map[string]func(c Client, id int) (*Resource, error){
		"validate": func(c Client, id int) (*Resource, error) {
			return c.ValidateResource(context.Background(), id)
		},
		"rotate_keys": func(c Client, id int) (*Resource, error) {
			return c.RotateTunnelKeyForResource(context.Background(), id)
		},
	} {
		action := k
		f := v

		t.Run(action, func(t *testing.T) {
			resID := 1
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				if want, got := fmt.Sprintf("%s/%d/actions", ResourcesBasePath, resID), req.URL.Path; want != got {
					t.Fatalf("mismatched of request path: want=%s got=%s", want, got)
				}

				var body = struct {
					Action string `json:"action"`
				}{}

				if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
					t.Errorf("expected no error, got %+v", err)
				}
				defer req.Body.Close()

				if want, got := action, body.Action; want != got {
					t.Errorf("unexpected action: want=%s got=%s", want, got)
				}

				// Return response to satisfy client and test response
				c := generateResource("test", resID, "", nil)
				json.NewEncoder(w).Encode(c)
			}))

			// Close the server when test finishes
			defer server.Close()

			c := testClient(server.Client(), server.URL)

			resp, err := f(c, resID)
			if err != nil {
				t.Errorf("expected no error, got %+v", err)
			}

			if want, got := resID, resp.ID; want != got {
				t.Errorf("unexpected resource ID: want=%d got %d", want, got)
			}

		})

	}

}

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
		{"s3://foo@bar:/barfoo/+@us-east-1/bucket", "s3://foo%40bar:%2Fbarfoo%2F+@us-east-1/bucket", nil},
		{"s3://foo@us-east-1/bucket", "s3://foo:@us-east-1/bucket", nil},
		{"s3://foo:@us-east-1/bucket", "s3://foo:@us-east-1/bucket", nil},
		{"s3://:bar@us-east-1/bucket", "s3://:bar@us-east-1/bucket", nil},
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

func TestCreateResource(t *testing.T) {
	var resource CreateResourceInput

	resource.Name = "resource-name"
	resource.URL = "http://foo.com"
	resource.Metadata = map[string]interface{}{
		"key": "value",
	}
	resource.SSHTunnel = &ResourceSSHTunnelInput{
		Address:    "test@host.com",
		PrivateKey: "1234",
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if want, got := fmt.Sprintf("%s", ResourcesBasePath), req.URL.Path; want != got {
			t.Fatalf("mismatched of request path: want=%s got=%s", want, got)
		}

		var rr *CreateResourceInput
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

		if !reflect.DeepEqual(rr.SSHTunnel, resource.SSHTunnel) {
			t.Errorf("expected same ssh tunnel")
		}

		// Return response to satisfy client and test response
		c := generateResource(resource.Name, 0, "", nil)
		c.URL = resource.URL
		c.Metadata = resource.Metadata
		c.SSHTunnel = &ResourceSSHTunnel{
			Address:   resource.SSHTunnel.Address,
			PublicKey: "1234",
		}
		json.NewEncoder(w).Encode(c)
	}))

	// Close the server when test finishes
	defer server.Close()

	c := testClient(server.Client(), server.URL)

	resp, err := c.CreateResource(context.Background(), &resource)
	if err != nil {
		t.Errorf("expected no error, got %+v", err)
	}

	if resp.URL != resource.URL {
		t.Errorf("expected url %s, got %s", resource.URL, resp.URL)
	}

	if want, got := resource.SSHTunnel.Address, resp.SSHTunnel.Address; want != got {
		t.Errorf("unexpected ssh tunnel address: want=%s got=%s", want, got)
	}

	if want, got := "1234", resp.SSHTunnel.PublicKey; want != got {
		t.Errorf("unexpected ssh tunnel public key: want=%s got=%s", want, got)
	}

	if want, got := ResourceStateReady, string(resp.Status.State); want != got {
		t.Errorf("unexpected status state: want=%s got=%s", want, got)
	}

	if want, got := "your resource is ready to use", resp.Status.Details; want != got {
		t.Errorf("unexpected status details: want=%s got=%s", want, got)
	}

	if resp.Status.LastUpdatedAt.IsZero() {
		t.Errorf("expected time to not be null: got=%s", resp.Status.LastUpdatedAt)
	}
}

func TestUpdateResource(t *testing.T) {
	var resource UpdateResourceInput

	resource.Name = "resource-name"
	resource.URL = "http://foo.com"
	resource.Metadata = map[string]interface{}{
		"key": "value",
	}
	resource.SSHTunnel = &ResourceSSHTunnelInput{
		Address: "test@host.com",
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

		if !reflect.DeepEqual(rr.SSHTunnel, resource.SSHTunnel) {
			t.Errorf("expected same ssh tunnel")
		}

		// Return response to satisfy client and test response
		c := generateResource(resource.Name, 0, "", nil)
		c.URL = resource.URL
		c.Metadata = resource.Metadata
		c.SSHTunnel = &ResourceSSHTunnel{
			Address:   resource.SSHTunnel.Address,
			PublicKey: "1234",
		}
		json.NewEncoder(w).Encode(c)
	}))

	// Close the server when test finishes
	defer server.Close()

	c := testClient(server.Client(), server.URL)

	resp, err := c.UpdateResource(context.Background(), resource.Name, &resource)
	if err != nil {
		t.Errorf("expected no error, got %+v", err)
	}

	if resp.URL != resource.URL {
		t.Errorf("expected url %s, got %s", resource.URL, resp.URL)
	}

	if want, got := resource.SSHTunnel.Address, resp.SSHTunnel.Address; want != got {
		t.Errorf("unexpected ssh tunnel address: want=%s got=%s", want, got)
	}

	if want, got := "1234", resp.SSHTunnel.PublicKey; want != got {
		t.Errorf("unexpected ssh tunnel public key: want=%s got=%s", want, got)
	}

	if want, got := ResourceStateReady, string(resp.Status.State); want != got {
		t.Errorf("unexpected status state: want=%s got=%s", want, got)
	}

	if want, got := "your resource is ready to use", resp.Status.Details; want != got {
		t.Errorf("unexpected status details: want=%s got=%s", want, got)
	}

	if resp.Status.LastUpdatedAt.IsZero() {
		t.Errorf("expected time to not be null: got=%s", resp.Status.LastUpdatedAt)
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
		Status: ResourceStatus{
			State:         "ready",
			Details:       "your resource is ready to use",
			LastUpdatedAt: time.Now(),
		},
	}
}
