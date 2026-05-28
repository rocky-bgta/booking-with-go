package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {

	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.Form)

	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid when should have been valid")
	}
}

func TestForm_Redirect(t *testing.T) {

	r := httptest.NewRequest("GET", "/whatever", nil)
	form := New(r.Form)

	form.Required("a", "b", "c")

	if form.Valid() {
		t.Error("form shows valid when required field is missing")
	}

	postData := url.Values{}
	postData.Add("a", "a")
	postData.Add("b", "a")
	postData.Add("c", "c")

	r, _ = http.NewRequest("POST", "/whatever", nil)

	r.PostForm = postData

	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("shows does not have required fields when it dose")
	}

}
