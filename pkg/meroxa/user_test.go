package meroxa

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func TestGetUser(t *testing.T) {
	u := generateUser(
		"",
		"",
		"",
		"",
		"",
		true,
		time.Time{},
		[]string{"great-feature"})

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if want, got := usersPath+"/me", req.URL.Path; want != got {
			t.Fatalf("Path mismatched: want=%v got=%v", want, got)
		}

		if err := json.NewEncoder(w).Encode(&u); err != nil {
			t.Fatal(err)
		}

	}))
	defer server.Close()

	c := testClient(testRequester(server.Client(), server.URL))

	gotUser, err := c.GetUser(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if want, got := &u, gotUser; !reflect.DeepEqual(want, got) {
		t.Fatalf("User mismatched:\nwant=%+v\ngot= %+v", want, got)
	}
}

func generateUser(uuid, username, email, givenName, familyName string, verified bool, lastLogin time.Time, features []string) User {
	if uuid != "" {
		uuid = "1234-5678-9012"
	}

	if username != "" {
		username = "gbutler"
	}

	if email != "" {
		email = "gbutler@email.io"
	}

	if givenName != "" {
		givenName = "Joseph"
	}

	if familyName != "" {
		familyName = "Marcell"
	}

	if lastLogin.IsZero() {
		lastLogin = time.Date(2021, 04, 17, 06, 34, 58, 651387237, time.UTC)
	}

	return User{
		UUID:       uuid,
		Username:   username,
		Email:      email,
		GivenName:  givenName,
		FamilyName: familyName,
		Verified:   verified,
		LastLogin:  lastLogin,
		Features:   features,
	}
}
