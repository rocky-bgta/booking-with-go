package handlers

import (
	"net/http"
	"net/http/httptest"
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
	{"about", "/about", "GET", []postData{}, http.StatusOK},
	{"gq", "/generals-quarters", "GET", []postData{}, http.StatusOK},
	{"ms", "/majors-suite", "GET", []postData{}, http.StatusOK},
	{"sa", "/search-availability", "GET", []postData{}, http.StatusOK},
	{"contact", "/contact", "GET", []postData{}, http.StatusOK},
	{"res", "/make-reservation", "GET", []postData{}, http.StatusOK},
}

func TestNewHandlers(t *testing.T) {
	routes := getRoutes()
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _, e := range theTests {
		t.Run(e.name, func(t *testing.T) {
			if e.method == "GET" {
				resp, err := ts.Client().Get(ts.URL + e.url)
				if err != nil {
					t.Fatalf("request failed: %v", err)
				}
				defer resp.Body.Close()

				if resp.StatusCode != e.expectedStatusCode {
					t.Errorf("%s %s: expected status %d, got %d", e.method, e.url, e.expectedStatusCode, resp.StatusCode)
				}

			} else if e.method == "POST" {
				// POST test cases can be implemented here if needed
				t.Skip("POST test cases not yet implemented")
			}
		})
	}

}
