// +build integration

package auth

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"
)

func TestTokenSource_Integration_Token(t *testing.T) {
	refreshToken := os.Getenv("MEROXA_REFRESH_TOKEN")
	if refreshToken == "" {
		t.Skipf("MEROXA_REFRESH_TOKEN is not set, skipping integration test")
	}

	c, err := NewClient(
		&http.Client{
			Timeout: 5 * time.Second,
		},
		DefaultConfig(),
		"",
		refreshToken,
	)
	if err != nil {
		t.Fatalf("could not create client: %v", err)
	}

	server := httptest.NewServer(http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		auth := req.Header.Get("Authorization")
		if auth == "" ||
			!strings.HasPrefix(auth, "Bearer ") ||
			len(auth[7:]) == 0 {
			t.Fatalf("Unexpected Authorization header: %q", auth)
		}
		resp.WriteHeader(200)
	}))

	_, err = c.Get(server.URL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
