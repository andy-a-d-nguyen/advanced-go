package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

var expect = []byte("Don't communicate by sharing memory, share memory by communicating.")

func TestService(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(expect))
	rec := httptest.NewRecorder()

	handler(rec, req)

	res := rec.Result()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(body) != string(expect) {
		t.Fail()
	}
}

func TestServer(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(handler))
	defer s.Close()

	c := s.Client()
	res, err := c.Post(s.URL, "text/plain", bytes.NewBuffer(expect))
	if err != nil {
		t.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(body) != string(expect) {
		t.Fail()
	}

}
