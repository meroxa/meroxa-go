// +build integration

package meroxa

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"golang.org/x/oauth2"
)

func TestNewAuthClient_Integration_Token(t *testing.T) {
	refreshToken := os.Getenv("MEROXA_REFRESH_TOKEN")
	if refreshToken == "" {
		t.Skipf("MEROXA_REFRESH_TOKEN is not set, skipping integration test")
	}

	tokenChan := make(chan *oauth2.Token)
	c, err := newAuthClient(
		&http.Client{
			Timeout: 5 * time.Second,
		},
		DefaultOAuth2Config(),
		"",
		refreshToken,
		func(token *oauth2.Token) {
			tokenChan <- token
		},
	)
	if err != nil {
		t.Fatalf("could not create client: %v", err)
	}

	server := httptest.NewServer(http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		token := <-tokenChan
		wantAuth := fmt.Sprintf("%s %s", token.TokenType, token.AccessToken)

		gotAuth := req.Header.Get("Authorization")
		if gotAuth != wantAuth {
			t.Fatalf("expected %q, got %q", wantAuth, gotAuth)
		}

		resp.WriteHeader(200)
	}))

	_, err = c.Get(server.URL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
