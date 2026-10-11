package main

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {
	rec := httptest.NewRecorder()
	handler(rec, httptest.NewRequest("GET", "/", nil))

	if rec.Code != 200 {
		t.Fatalf("status = %d, quero 200", rec.Code)
	}
	if !strings.HasPrefix(rec.Body.String(), "hello world from ") {
		t.Fatalf("corpo inesperado: %q", rec.Body.String())
	}
}
