package meroxa

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"testing"
)

func TestListAccounts(t *testing.T) {
	testCases := []struct {
		desc             string
		requester        func() *mockRequester
		expectedAccounts []Account
		expectedError    error
	}{
		{
			desc: "get list of account",
			requester: func() *mockRequester {
				m := newMockRequester()
				m.response = &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(strings.NewReader(`[{"uuid":"123456","name":"TestAccount"}]`)),
				}
				return m
			},
			expectedAccounts: []Account{{UUID: "123456", Name: "TestAccount"}},
		},
		{
			desc: "MakeRequest can respond with something different than 200 OK",
			requester: func() *mockRequester {
				m := newMockRequester()
				m.response = &http.Response{
					StatusCode: http.StatusBadRequest,
					Body:       io.NopCloser(strings.NewReader(`{"code":"100","message":"test error"}`)),
				}
				return m
			},
			expectedError: &errResponse{Code: "100", Message: "test error"},
		},
		{
			desc: "MakeRequest can return a different error code",
			requester: func() *mockRequester {
				m := newMockRequester()
				m.response = nil
				m.err = errors.New("an error")
				return m
			},
			expectedError: errors.New("an error"),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			ctx := context.Background()
			client := testClient(tc.requester())

			accounts, err := client.ListAccounts(ctx)
			if err != nil {
				if err.Error() != tc.expectedError.Error() {
					t.Fatalf("expected error: %s, got: %s", tc.expectedError, err)
				}
			}

			if tc.expectedAccounts != nil {
				if want, got := &tc.expectedAccounts[0], accounts[0]; !reflect.DeepEqual(want, got) {
					t.Fatalf("Accounts mismatched:\nwant=%+v\ngot= %+v", want, got)
				}
			}
		})
	}
}

type mockRequester struct {
	response *http.Response
	err      error
}

func newMockRequester() *mockRequester {
	m := &mockRequester{}
	return m
}

func (m mockRequester) MakeRequest(ctx context.Context, method string, path string, body interface{}, params url.Values, headers http.Header) (*http.Response, error) {
	return m.response, m.err
}

func (m mockRequester) AddHeader(key, value string) {
	m.response.Header.Add(key, value)
}
