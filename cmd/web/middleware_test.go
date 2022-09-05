package main

import (
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {
	var myH myHandler
	h := NoSurf(myH)
	switch h.(type) {
	case http.Handler:
		//do nothing
	default:
		t.Error("NoSurf return is not of type http.Handler")
	}
}

func TestSessionLoad(t *testing.T) {
	var myH myHandler
	h := SessionLoad(myH)
	switch h.(type) {
	case http.Handler:
		//do nothing
	default:
		t.Error("SessionLoad return is not of type http.Handler")
	}
}
