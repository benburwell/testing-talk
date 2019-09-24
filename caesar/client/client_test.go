package client

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestEncode(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "response")
	}))
	defer server.Close()
	client := CaesarClient{Endpoint: server.URL} // HL
	var b bytes.Buffer
	if err := client.EncodeMessage(&b, strings.NewReader("secret message")); err != nil {
		t.Errorf("should not have gotten error: %v", err)
	}
	resp := b.String()
	if resp != "response" {
		t.Errorf("expected 'response' but got '%s'", resp)
	}
}

func TestUserAgent(t *testing.T) {
	var ua string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ua = r.Header.Get("user-agent")
	}))
	defer server.Close()

	c := CaesarClient{Endpoint: server.URL}
	if err := c.EncodeMessage(ioutil.Discard, strings.NewReader("secret message")); err != nil {
		t.Errorf("expected no error but got %v", err)
	}

	expected := "caesar-client/1.0"
	if ua != expected {
		t.Errorf("expected user agent header %q but got %q", expected, ua)
	}
}
