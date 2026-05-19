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

	t.Logf("Starting test server at: %s\n", ts.URL)

	for _, e := range theTests {
		t.Run(e.name, func(t *testing.T) {
			t.Logf("[%s] Testing %s %s", e.name, e.method, e.url)

			if e.method == "GET" {
				resp, err := ts.Client().Get(ts.URL + e.url)
				if err != nil {
					t.Fatalf("request failed: %v", err)
				}
				defer resp.Body.Close()

				t.Logf("[%s] Response status: %d (expected: %d)", e.name, resp.StatusCode, e.expectedStatusCode)

				if resp.StatusCode != e.expectedStatusCode {
					t.Errorf("%s %s: expected status %d, got %d", e.method, e.url, e.expectedStatusCode, resp.StatusCode)
				} else {
					t.Logf("[%s] ✓ PASS", e.name)
				}

			} else if e.method == "POST" {
				// POST test cases can be implemented here if needed
				t.Skip("POST test cases not yet implemented")
			}
		})
	}

}
