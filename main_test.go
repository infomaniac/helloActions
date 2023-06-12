package main

import (
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	helloHandler(w, req)

	if w.Body.String() != "Hello, world!\n" {
		t.Errorf("unexpected body: %s", w.Body.String())
	}
}
