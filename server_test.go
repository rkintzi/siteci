package main

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestHook(t *testing.T) {
	var h Hook
	s := httptest.NewServer(&h)
	c := s.Client()
	resp, err := c.Get(s.URL)
	if err != nil {
		t.Fatalf("Get should not err. Got: %v", err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("Expected \"200 OK\". Got: \"%s\"", resp.Status)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("ReadAll should not err. Got: %v", err)
	}
	if len(body) != 0 {
		t.Errorf("Expected zero lenght body. Got: \"%s\"", string(body))
	}
}
