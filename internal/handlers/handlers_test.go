package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type postData struct {
	key   string
	value string
}

var theTests = []struct {
	name               string
	url                string
	method             string
	params             []postData
	expectedStatusCode int
}{
	{"home", "/", "GET", []postData{}, http.StatusOK},
	{"signup", "/signup", "GET", []postData{}, http.StatusOK},
	{"login", "/login", "GET", []postData{}, http.StatusOK},
	{"dashboard", "/dashboard", "GET", []postData{}, http.StatusOK},
	{"signupPost", "/signup", "POST", []postData{
		{key: "name", value: "Rafael"},
		{key: "username", value: "rafaelusername"},
		{key: "email", value: "me@email.com"},
		{key: "password", value: "P@ssW0rd"},
	}, http.StatusOK},
	{"check-username-available", "/check-username-available", "POST", []postData{
		{key: "username", value: "rafaelusername"},
	}, http.StatusOK},
}

func TestHandlers(t *testing.T) {
	routes := GetRoutes()
	testServer := httptest.NewTLSServer(routes)
	defer testServer.Close()
	for _, e := range theTests {
		if e.method == "GET" {
			resp, err := testServer.Client().Get(testServer.URL + e.url)
			if err != nil {
				t.Fatal(err)
			}
			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("Test %s expected statusCode %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}
		} else {
			values := url.Values{}
			for _, x := range e.params {
				values.Add(x.key, x.value)
			}
			resp, err := testServer.Client().PostForm(testServer.URL+e.url, values)
			if err != nil {
				t.Fatal(err)
			}
			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("Test %s expected statusCode %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}
		}
	}
}
