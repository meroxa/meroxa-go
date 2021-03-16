package meroxa

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
)

var httpNoErrorResponse = &http.Response{
	StatusCode: 200,
}

// create a new reader with a JSON response
var errorJSONResponse = ioutil.NopCloser(bytes.NewReader([]byte(`{ "error": "api error" }`)))
var errorJSONResponseCodeMessage = ioutil.NopCloser(bytes.NewReader([]byte(`{ "code": "already_exists", "message": "resource with name test already exists"}`)))
var errorJSONResponseDetails = ioutil.NopCloser(bytes.NewReader([]byte(
	`{ "code": "already_exists", "message": "resource with name test already exists", "details": { "name": ["too long", "invalid"], "type": ["invalid"] } }`,
)))

// create a new reader with an HTML response
var errorHTMLResponse = ioutil.NopCloser(bytes.NewReader([]byte(`<h1>Error!</h1>`)))

var http503JSONResponse = &http.Response{
	StatusCode: 503,
	Status:     "503 Service Unavailable",
	Body:       errorJSONResponse,
	Proto:      "HTTP/1.0",
	Header:     make(http.Header),
}

var http422JSONResponse = func(body io.ReadCloser) *http.Response {
	return &http.Response{
		StatusCode: 422,
		Status:     "422 Unprocessable Entity",
		Body:       body,
		Proto:      "HTTP/1.0",
		Header:     make(http.Header),
	}
}

var http503HTMLResponse = &http.Response{
	Status:     "503 Service Unavailable",
	StatusCode: 503,
	Body:       errorHTMLResponse,
	Proto:      "HTTP/1.0",
	Header:     make(http.Header),
}

func TestHandleAPIErrors(t *testing.T) {
	tests := []struct {
		in  *http.Response
		err error
	}{
		{httpNoErrorResponse, nil},
		{http503JSONResponse, &errResponse{ErrorDeprecated: "api error"}},
		{http503HTMLResponse, errors.New("HTTP/1.0 503 Service Unavailable")},
		{http422JSONResponse(errorJSONResponseCodeMessage),
			&errResponse{
				Code:    "already_exists",
				Message: "resource with name test already exists",
			}},
		{http422JSONResponse(errorJSONResponseDetails),
			&errResponse{
				Code:    "already_exists",
				Message: "resource with name test already exists",
				Details: map[string][]string{
					"name": {"too long", "invalid"},
					"type": {"invalid"},
				},
			}},
	}

	for _, tt := range tests {
		err := handleAPIErrors(tt.in)

		if !reflect.DeepEqual(err, tt.err) {
			t.Errorf("expected %+v, got %+v", tt.err, err)
		}
	}
}
