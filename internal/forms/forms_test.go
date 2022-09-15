package forms

import (
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	request := httptest.NewRequest("POST", "/url", nil)
	form := New(request.PostForm)
	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid form when it should be valid")
	}
}

func TestForm_Required(t *testing.T) {
	request := httptest.NewRequest("POST", "/url", nil)
	form := New(request.PostForm)
	form.Required("key1", "key2")
	if form.Valid() {
		t.Error("form evaluated as valid when it should be invalid")
	}
	postedData := url.Values{}
	postedData.Add("key1", "value1")
	postedData.Add("key2", "value2")
	request.PostForm = postedData
	form = New(postedData)
	form.Required("key1", "key2")
	if !form.Valid() {
		t.Error("form evaluated as invalid when it should be valid")
	}
}

func TestForm_Has(t *testing.T) {
	// TODO
}

func TestForm_IsEmail(t *testing.T) {
	// TODO
}
