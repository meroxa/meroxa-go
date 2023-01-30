package meroxa

import (
	"context"
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListResourceTypesV2(t *testing.T) {
	r1 := generateResourceType("r1")
	r2 := generateResourceType("r2")

	list := []ResourceType{r1, r2}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if want, got := V2ResourcesTypeBasePath, req.URL.Path; want != got {
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

	resp, err := c.ListResourceTypesV2(context.Background())
	if err != nil {
		t.Errorf("expected no error, got %+v", err)
	}

	if diff := cmp.Diff(resp, list); diff != "" {
		t.Fatalf("mismatch of function output (-want +got): %s", diff)
	}
}

func generateResourceType(name string) ResourceType {
	if name == "" {
		name = "test"
	}

	return ResourceType{
		UUID:         uuid.NewString(),
		Name:         name,
		ReleaseStage: ResourceTypeReleaseStageGA,
		OptedIn:      true,
		HasAccess:    true,
		CliOnly:      true,
	}

}
