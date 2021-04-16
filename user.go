package meroxa

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const usersPath = "/v1/users"

type User struct {
	UUID       string `json:"uuid"`
	Username   string `json:"preferred_username,omitempty"`
	Email      string `json:"email,omitempty"`
	GivenName  string `json:"given_name,omitempty"`
	FamilyName string `json:"family_name,omitempty"`
	Verified   bool   `json:"email_verified,omitempty"`
}

// GetUser returns a User with
func (c *Client) GetUser(ctx context.Context) (*User, error) {
	path := fmt.Sprintf("%s/me", usersPath)

	resp, err := c.makeRequest(ctx, http.MethodGet, path, nil, nil)
	if err != nil {
		return nil, err
	}

	err = handleAPIErrors(resp)
	if err != nil {
		return nil, err
	}

	var u User
	err = json.NewDecoder(resp.Body).Decode(&u)
	if err != nil {
		return nil, err
	}

	return &u, nil
}
