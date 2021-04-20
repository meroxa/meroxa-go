package auth

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/cristalhq/jwt/v3"
	"golang.org/x/oauth2"
)

type tokenSource struct {
	client       *http.Client
	conf         *oauth2.Config
	refreshToken string
}

func (ts *tokenSource) Token() (*oauth2.Token, error) {
	if ts.refreshToken == "" {
		return nil, errors.New("oauth2: token expired and refresh token is not set")
	}

	tk, err := retrieveToken(ts.client, ts.conf, ts.refreshToken)
	if err != nil {
		return nil, err
	}

	if tk.RefreshToken == "" {
		tk.RefreshToken = ts.refreshToken
	}
	if ts.refreshToken != tk.RefreshToken {
		ts.refreshToken = tk.RefreshToken
	}

	return tk, err
}

func retrieveToken(client *http.Client, conf *oauth2.Config, refreshToken string) (*oauth2.Token, error) {
	tmp := make(map[string]interface{})
	tmp["client_id"] = conf.ClientID
	tmp["grant_type"] = "refresh_token"
	tmp["refresh_token"] = refreshToken
	requestBody, err := json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	resp, err := client.Post(conf.Endpoint.TokenURL, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("oauth2: cannot fetch token: %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(io.LimitReader(resp.Body, 1<<20))
	if err != nil {
		return nil, fmt.Errorf("oauth2: cannot fetch token: %v", err)
	}

	if c := resp.StatusCode; c < 200 || c > 299 {
		return nil, &oauth2.RetrieveError{
			Response: resp,
			Body:     body,
		}
	}

	// tokenRes is the JSON response body.
	var tokenRes struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		ExpiresIn   int64  `json:"expires_in"` // relative seconds from now
		// Ignored fields
		// Scope       string `json:"scope"`
		// IDToken     string `json:"id_token"`
	}
	if err := json.Unmarshal(body, &tokenRes); err != nil {
		return nil, fmt.Errorf("oauth2: cannot fetch token: %v", err)
	}
	token := &oauth2.Token{
		AccessToken: tokenRes.AccessToken,
		TokenType:   tokenRes.TokenType,
	}
	raw := make(map[string]interface{})
	json.Unmarshal(body, &raw) // no error checks for optional fields
	token = token.WithExtra(raw)

	expiry, err := getTokenExpiry(token.AccessToken)
	if err != nil {
		// fallback to calculate expiry
		expiry = time.Unix(tokenRes.ExpiresIn, 0).UTC()
	}
	token.Expiry = expiry

	return token, nil
}

func getTokenExpiry(token string) (time.Time, error) {
	jwtToken, err := jwt.ParseString(token)
	if err != nil {
		return time.Time{}, fmt.Errorf("could not parse access token as JWT: %w", err)
	}

	var claims jwt.StandardClaims
	err = json.Unmarshal(jwtToken.RawClaims(), &claims)
	if err != nil {
		return time.Time{}, fmt.Errorf("could not parse access token JWT claims: %w", err)
	}

	return claims.ExpiresAt.Time.UTC(), nil
}
