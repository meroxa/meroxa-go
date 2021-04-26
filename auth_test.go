package meroxa

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"golang.org/x/oauth2"
)

func TestTokenSource_Token(t *testing.T) {
	clientID := "my-client-id"
	refreshToken := "my-refresh-token"

	server := httptest.NewServer(http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		var body map[string]interface{}
		err := json.NewDecoder(req.Body).Decode(&body)
		if err != nil {
			t.Fatalf("could not read request body: %v", err)
		}

		wantBody := map[string]interface{}{
			"client_id":     clientID,
			"refresh_token": refreshToken,
			"grant_type":    "refresh_token",
		}
		if diff := cmp.Diff(wantBody, body); diff != "" {
			t.Fatal(diff)
		}

		_, err = resp.Write([]byte(`{"access_token":"eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6ImRycTRuN0NMNEt0WmxFZGlBRU9uWiJ9.eyJodHRwczovL2FwaS5tZXJveGEuaW8vdjEvZW1haWwiOiJsb3Zyb0BtZXJveGEuaW8iLCJodHRwczovL2FwaS5tZXJveGEuaW8vdjEvdXNlcm5hbWUiOiJMb3ZybyBNYcW-Z29uIiwiaXNzIjoiaHR0cHM6Ly9hdXRoLm1lcm94YS5pby8iLCJzdWIiOiJnb29nbGUtb2F1dGgyfDEwOTk0NTQ3OTI5ODExMTczMTI2MyIsImF1ZCI6WyJodHRwczovL2FwaS5tZXJveGEuaW8vdjEiLCJodHRwczovL21lcm94YS5hdXRoMC5jb20vdXNlcmluZm8iXSwiaWF0IjoxNjE4OTE3MDU5LCJleHAiOjE2MTg5MjU2NTksImF6cCI6IjJWQzl6MFp4dHpUY1FMRE55Z2VFRUxWM2xZRlJad3BiIiwic2NvcGUiOiJvcGVuaWQgZW1haWwgdXNlciBvZmZsaW5lX2FjY2VzcyJ9.zp9trLjYsnD__Q4jlH1O7x9WQ62C0WHdi3VVrflbb8Q0BqKcmeJDhoIigItodee1MQJ7o08loSIIdCx-H-gHZe2n63Mi1DkZa-WYmutuKHVVfXZ9aX-jEI7Kf0f5bxWdq3hYGO4p4JODZnpa0Mx_G-Lk7BZ9Qu7p5OiEvr0LLhtp3c5A26ZhGz2lkag9LhUer8SDocFpRONiwgQg8c1BYQUu2oIunHccSgAgmWXL3Yt7ww2SXn4odHjBbY7EqxKIzodEHUE6cjhUnOC5e9Kx7ThFTHV_pma0y_BV_fDys0YhLPhjBo0hdvQ8wLxi4wIrmkt42qmvQoPiBdTulvNp6Q","id_token":"[my-id-token]","scope":"openid email user offline_access","expires_in":8600,"token_type":"Bearer"}`))
		if err != nil {
			t.Fatalf("could not write response: %v", err)
		}
	}))
	t.Cleanup(server.Close)

	ts := &tokenSource{
		client: http.DefaultClient,
		conf: &oauth2.Config{
			ClientID: clientID,
			Endpoint: oauth2.Endpoint{
				TokenURL: server.URL,
			},
		},
		refreshToken: refreshToken,
	}

	got, err := ts.Token()
	if err != nil {
		t.Fatalf("could not retrieve token: %v", err)
	}

	want := &oauth2.Token{
		AccessToken:  "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6ImRycTRuN0NMNEt0WmxFZGlBRU9uWiJ9.eyJodHRwczovL2FwaS5tZXJveGEuaW8vdjEvZW1haWwiOiJsb3Zyb0BtZXJveGEuaW8iLCJodHRwczovL2FwaS5tZXJveGEuaW8vdjEvdXNlcm5hbWUiOiJMb3ZybyBNYcW-Z29uIiwiaXNzIjoiaHR0cHM6Ly9hdXRoLm1lcm94YS5pby8iLCJzdWIiOiJnb29nbGUtb2F1dGgyfDEwOTk0NTQ3OTI5ODExMTczMTI2MyIsImF1ZCI6WyJodHRwczovL2FwaS5tZXJveGEuaW8vdjEiLCJodHRwczovL21lcm94YS5hdXRoMC5jb20vdXNlcmluZm8iXSwiaWF0IjoxNjE4OTE3MDU5LCJleHAiOjE2MTg5MjU2NTksImF6cCI6IjJWQzl6MFp4dHpUY1FMRE55Z2VFRUxWM2xZRlJad3BiIiwic2NvcGUiOiJvcGVuaWQgZW1haWwgdXNlciBvZmZsaW5lX2FjY2VzcyJ9.zp9trLjYsnD__Q4jlH1O7x9WQ62C0WHdi3VVrflbb8Q0BqKcmeJDhoIigItodee1MQJ7o08loSIIdCx-H-gHZe2n63Mi1DkZa-WYmutuKHVVfXZ9aX-jEI7Kf0f5bxWdq3hYGO4p4JODZnpa0Mx_G-Lk7BZ9Qu7p5OiEvr0LLhtp3c5A26ZhGz2lkag9LhUer8SDocFpRONiwgQg8c1BYQUu2oIunHccSgAgmWXL3Yt7ww2SXn4odHjBbY7EqxKIzodEHUE6cjhUnOC5e9Kx7ThFTHV_pma0y_BV_fDys0YhLPhjBo0hdvQ8wLxi4wIrmkt42qmvQoPiBdTulvNp6Q",
		TokenType:    "Bearer",
		RefreshToken: refreshToken,
		Expiry:       time.Date(2021, 4, 20, 13, 34, 19, 0, time.UTC),
	}
	if diff := cmp.Diff(want, got, cmpopts.IgnoreUnexported(oauth2.Token{})); diff != "" {
		t.Fatal(diff)
	}
}
