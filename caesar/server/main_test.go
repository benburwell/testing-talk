package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandleCaesar(t *testing.T) {
	bodyReader := strings.NewReader("hello world")
	req := httptest.NewRequest("POST", "/", bodyReader)
	resp := httptest.NewRecorder()

	handler := handleCaesar()
	handler(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("expected OK status %d but got %d", http.StatusOK, resp.Code)
	}

	body := string(resp.Body.Bytes())
	if body != "khoor zruog" {
		t.Errorf("got unexpected body %s", body)
	}
}

func TestCaesarBadRequest(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	resp := httptest.NewRecorder()
	handler := handleCaesar()
	handler(resp, req)
	if resp.Code != http.StatusBadRequest {
		t.Errorf("status should have been bad request, got %d", resp.Code)
	}
	if len(resp.Body.Bytes()) != 0 {
		t.Errorf("should not have gotten a response body")
	}
}
