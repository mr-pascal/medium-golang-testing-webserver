package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"

	"net/http"
	"net/http/httptest"
	"testing"
)

// Equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func Equal(a, b []interface{}) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

type AppHandlerFake struct {
	// method name -> call -> params
	Calls map[string][][]interface{}
}

func (a *AppHandlerFake) Sum(x, y int) (r Result) {
	b := a.Calls[`Sum`]
	c := []interface{}{x, y}
	a.Calls[`Sum`] = append(b, c)

	r.Value = 7
	return
}
func (a *AppHandlerFake) Multiply(x, y int) (r Result) {
	b := a.Calls[`Multiply`]
	c := []interface{}{x, y}

	a.Calls[`Multiply`] = append(b, c)

	r.Value = 10
	return
}

func TestApp(t *testing.T) {

	testTable := []struct {
		name                  string
		method                string
		path                  string
		statusCode            int
		body                  string
		requestBody           map[string]interface{}
		handlerMethodName     string
		handlerToBeCalledWith []interface{}
		requestHeaders        map[string]string
		headers               map[string]string
	}{
		{
			name:                  `GET endpoint to get a sum`,
			method:                http.MethodGet,
			path:                  `/sum?x=5&y=2`,
			statusCode:            200,
			requestBody:           nil,
			body:                  `{"value":7}`,
			handlerMethodName:     "Sum",
			handlerToBeCalledWith: []interface{}{5, 2},
			headers:               map[string]string{`Content-Type`: `application/json`},
		},
		{
			name:       `POST endpoint to multiply, wrong header`,
			method:     http.MethodPost,
			path:       `/multiply`,
			statusCode: 400,
			requestBody: map[string]interface{}{
				"x": 2,
				"y": 3,
			},
			body:           `Invalid payload`,
			requestHeaders: map[string]string{`Content-Type`: `application/text`},
			// headers:        map[string]string{`Content-Type`: `application/json`},
		},
		{
			name:       `POST endpoint to multiply`,
			method:     http.MethodPost,
			path:       `/multiply`,
			statusCode: 200,
			requestBody: map[string]interface{}{
				"x": 4,
				"y": 5,
			},
			body:                  `{"value":10}`,
			handlerMethodName:     "Multiply",
			handlerToBeCalledWith: []interface{}{4, 5},
			headers:               map[string]string{`Content-Type`: `application/json`},
		},
	}

	appHandler := &AppHandlerFake{}
	app := CreateApp(appHandler)

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {

			appHandler.Calls = map[string][][]interface{}{}

			// Create and send request
			rbody, _ := json.Marshal(tc.requestBody)
			request := httptest.NewRequest(tc.method, tc.path, bytes.NewReader(rbody))
			request.Header.Add(`Content-Type`, `application/json`)

			// Request Headers
			for k, v := range tc.requestHeaders {
				request.Header.Add(k, v)
			}

			response, _ := app.Test(request)

			// Status Code
			statusCode := response.StatusCode
			if statusCode != tc.statusCode {
				t.Errorf("StatusCode was incorrect, got: %d, want: %d.", statusCode, tc.statusCode)
			}

			// Headers
			for k, want := range tc.headers {
				headerValue := response.Header.Get(k)
				if headerValue != want {
					t.Errorf("Response header '%s' was incorrect, got: '%s', want: '%s'", k, headerValue, want)
				}
			}

			// Response Body
			body, _ := ioutil.ReadAll(response.Body)
			actual := string(body)
			if actual != tc.body {
				t.Errorf("Body was incorrect, got: %v, want: %v", actual, tc.body)
			}

			// Check if handler was called correctly
			if tc.handlerMethodName != "" {
				if !Equal(appHandler.Calls[tc.handlerMethodName][0], tc.handlerToBeCalledWith) {
					t.Errorf("Handler method '%s' wasn't called with the correct parameters. Got: '%v', want: '%v'", tc.handlerMethodName, appHandler.Calls[tc.handlerMethodName][0], tc.handlerToBeCalledWith)
				}
			}
		})
	}

}
