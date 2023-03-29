package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNoSurve(t *testing.T) {
	var myH myHandler

	h := NoSurf(&myH)

	switch v := h.(type) {
	case http.Handler:
	//does nothing
	default:
		t.Error(fmt.Sprintf("Type is not http.handler, but is %T", v))
	}
}

func TestSessionLoad(t *testing.T) {
	var myH myHandler

	h := SessionLoad(&myH)

	switch v := h.(type) {
	case http.Handler:
	//does nothing
	default:
		t.Error(fmt.Sprintf("Type is not http.handler, but is %T", v))
	}
}
