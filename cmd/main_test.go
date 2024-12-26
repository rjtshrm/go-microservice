package main

import (
	"net/http"
	"testing"
)

func TestService(t *testing.T) {
	go main() // Start the server in a goroutine.
	resp, err := http.Get("http://localhost:8080/")

	if err != nil {
		t.Errorf("Could not send request: %v", err)
	}
	defer resp.Body.Close()

	// Check the status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK, got %v", resp.StatusCode)
	}
}
