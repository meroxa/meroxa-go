package meroxa

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

const (
	address    = "test@host.com"
	privateKey = "1234"
)

func Test_Resource_PerformActions(t *testing.T) {
	for k, v := range map[string]func(c Client, id string) (*Resource, error){
		"validate": func(c Client, name string) (*Resource, error) {
			return c.ValidateResource(context.Background(), name)
		},
		"rotate_keys": func(c Client, name string) (*Resource, error) {
			return c.RotateTunnelKeyForResource(context.Background(), name)
		},
	} {
		action := k
		f := v

		t.Run(action, func(t *testing.T) {
			resName := "resource-1"
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				if want, got := fmt.Sprintf("%s/%s/actions", ResourcesBasePath, resName), req.URL.Path; want != got {
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
				c := generateResource(resName, "", nil)
				if err := json.NewEncoder(w).Encode(c); err != nil {
					t.Errorf("expected no error, got %+v", err)
				}
			}))

			// Close the server when test finishes
			defer server.Close()

			c := testClient(testRequester(server.Client(), server.URL))

			resp, err := f(c, resName)
			if err != nil {
				t.Errorf("expected no error, got %+v", err)
			}

			if want, got := resName, resp.Name; want != got {
				t.Errorf("unexpected resource name: want=%s got %s", want, got)
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
		{in: "s3://KAHDKJKSA:askkshe+skje/fhds@us-east-1/bucket", want: "s3://KAHDKJKSA:askkshe+skje%2Ffhds@us-east-1/bucket", err: nil},
		{in: "s3://KAHDKJKSA:secretsecret@us-east-1/bucket", want: "s3://KAHDKJKSA:secretsecret@us-east-1/bucket", err: nil},
		{in: "s3://us-east-1/bucket", want: "s3://us-east-1/bucket", err: nil},
		{in: "s3://:apassword@us-east-1/bucket", want: "s3://:apassword@us-east-1/bucket", err: nil},
		{in: "s3://foo@bar:/barfoo/+@us-east-1/bucket", want: "s3://foo%40bar:%2Fbarfoo%2F+@us-east-1/bucket", err: nil},
		{in: "s3://foo@us-east-1/bucket", want: "s3://foo:@us-east-1/bucket", err: nil},
		{in: "s3://foo:@us-east-1/bucket", want: "s3://foo:@us-east-1/bucket", err: nil},
		{in: "s3://:bar@us-east-1/bucket", want: "s3://:bar@us-east-1/bucket", err: nil},
		{in: "not a URL", want: "", err: ErrMissingScheme},
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
	tests := []struct {
		desc  string
		input func() CreateResourceInput
	}{
		{
			desc: "resource without an environment",
			input: func() CreateResourceInput {
				var resource CreateResourceInput

				resource.Name = "resource-name"
				resource.URL = "http://foo.com"
				resource.Metadata = map[string]interface{}{
					"key": "value",
				}
				resource.SSHTunnel = &ResourceSSHTunnelInput{
					Address:    address,
					PrivateKey: privateKey,
				}

				return resource
			},
		},
		{
			desc: "resource with an environment",
			input: func() CreateResourceInput {
				var resource CreateResourceInput
				var env = &EntityIdentifier{
					Name: "my-environment",
				}

				resource.Environment = env
				resource.Name = "resource-name"
				resource.URL = "http://foo.com"
				resource.Metadata = map[string]interface{}{
					"key": "value",
				}
				resource.SSHTunnel = &ResourceSSHTunnelInput{
					Address:    address,
					PrivateKey: privateKey,
				}
				return resource
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			var resource = tc.input()

			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				if want, got := ResourcesBasePath, req.URL.Path; want != got {
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
				c := generateResource(resource.Name, "", nil)
				c.URL = resource.URL
				c.Metadata = resource.Metadata
				c.SSHTunnel = &ResourceSSHTunnel{
					Address:   resource.SSHTunnel.Address,
					PublicKey: privateKey,
				}

				if resource.Environment != nil {
					c.Environment = &EntityIdentifier{Name: resource.Environment.Name}
				}

				if err := json.NewEncoder(w).Encode(c); err != nil {
					t.Errorf("expected no error, got %+v", err)
				}
			}))

			// Close the server when test finishes
			defer server.Close()

			c := testClient(testRequester(server.Client(), server.URL))

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

			if want, got := privateKey, resp.SSHTunnel.PublicKey; want != got {
				t.Errorf("unexpected ssh tunnel public key: want=%s got=%s", want, got)
			}

			if want, got := ResourceStateReady, resp.Status.State; want != got {
				t.Errorf("unexpected status state: want=%s got=%s", want, got)
			}

			if want, got := "your resource is ready to use", resp.Status.Details; want != got {
				t.Errorf("unexpected status details: want=%s got=%s", want, got)
			}

			if resp.Status.LastUpdatedAt.IsZero() {
				t.Errorf("expected time to not be null: got=%s", resp.Status.LastUpdatedAt)
			}

			if resource.Environment != nil {
				if want, got := resource.Environment.Name, resp.Environment.Name; want != got {
					t.Errorf("unexpected environment name: want=%s got=%s", want, got)
				}
			}
		})
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
		c := generateResource(resource.Name, "", nil)
		c.URL = resource.URL
		c.Metadata = resource.Metadata
		c.SSHTunnel = &ResourceSSHTunnel{
			Address:   resource.SSHTunnel.Address,
			PublicKey: privateKey,
		}
		if err := json.NewEncoder(w).Encode(c); err != nil {
			t.Errorf("expected no error, got %+v", err)
		}
	}))

	// Close the server when test finishes
	defer server.Close()

	c := testClient(testRequester(server.Client(), server.URL))

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

	if want, got := privateKey, resp.SSHTunnel.PublicKey; want != got {
		t.Errorf("unexpected ssh tunnel public key: want=%s got=%s", want, got)
	}

	if want, got := ResourceStateReady, resp.Status.State; want != got {
		t.Errorf("unexpected status state: want=%s got=%s", want, got)
	}

	if want, got := "your resource is ready to use", resp.Status.Details; want != got {
		t.Errorf("unexpected status details: want=%s got=%s", want, got)
	}

	if resp.Status.LastUpdatedAt.IsZero() {
		t.Errorf("expected time to not be null: got=%s", resp.Status.LastUpdatedAt)
	}
}

func TestIntrospectResource(t *testing.T) {
	output := generateResource("test", "postgres://example.com:5432/mydb", nil)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if diff := cmp.Diff(http.MethodGet, req.Method); diff != "" {
			t.Fatalf("mismatched of request method (-want +got): %s", diff)
		}

		if diff := cmp.Diff(fmt.Sprintf("%s/%s/introspection", ResourcesBasePath, output.Name), req.URL.Path); diff != "" {
			t.Fatalf("mismatched of request path (-want +got): %s", diff)
		}

		introspectionResponse := ResourceIntrospection{
			ID:              0,
			AccountID:       110,
			UUID:            output.UUID,
			ResourceID:      100,
			Collections:     []string{"users", "events"},
			Schemas:         map[string]string{},
			Capabilities:    nil,
			Samples:         nil,
			ResourceVersion: "9.6",
			IntrospectedAt:  time.Now().UTC(),
		}

		err := json.NewEncoder(w).Encode(introspectionResponse)
		if err != nil {
			t.Errorf("expected no error, got %+v", err)
		}
	}))
	defer server.Close()

	c := testClient(testRequester(server.Client(), server.URL))

	_, err := c.IntrospectResource(context.Background(), output.Name)
	if err != nil {
		t.Fatalf("expected no error, got %+v", err)
	}
}

func generateResource(name string, url string, metadata map[string]interface{}) Resource {
	if name == "" {
		name = "test"
	}

	return Resource{
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
