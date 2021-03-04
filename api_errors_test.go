package meroxa

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
)

var httpNoErrorResponse = &http.Response{
	StatusCode: 200,
}

// create a new reader with a JSON response
var  errorJSONResponse = ioutil.NopCloser(bytes.NewReader([]byte(`{ "error": "api error" }`)))

// create a new reader with an HTML response
var errorHTMLResponse = ioutil.NopCloser(bytes.NewReader([]byte(`<h1>Error!</h1>`)))

var http503JSONResponse = &http.Response{
	StatusCode: 503,
	Status: "503 Service Unavailable",
	Body: errorJSONResponse,
	Proto: "HTTP/1.0",
	Header: make(http.Header),
}

var http503HTMLResponse = &http.Response{
	Status: "503 Service Unavailable",
	StatusCode: 503,
	Body: errorHTMLResponse,
	Proto: "HTTP/1.0",
	Header: make(http.Header),
}

func TestHandleAPIErrors(t *testing.T) {
	tests := []struct {
		in  *http.Response
		err error
	}{
		{httpNoErrorResponse, nil},
		{http503JSONResponse ,errors.New("api error")},
		{http503HTMLResponse ,errors.New("error: HTTP/1.0 503 Service Unavailable")},
	}

	for _, tt := range tests {
		err := handleAPIErrors(tt.in)

		if !reflect.DeepEqual(err, tt.err) {
			t.Errorf("expected %+v, got %+v", tt.err, err)
		}
	}
}
