package auth

import (
	"net/http"
	"testing"
)

func TestMissingAuthorzationHeaderFails(t *testing.T) {
	testHeader := http.Header{}
	testHeader.Set("Invalid", "")

	_, err := GetAPIKey(testHeader)
	if err == nil {
		t.Fatalf("Expected error: %v. Got nil", ErrNoAuthHeaderIncluded.Error())
	}
}

func TestMalformedHeaderFails(t *testing.T) {
	testHeader := http.Header{}
	testHeader.Set("Authorization", "invalid")

	_, err := GetAPIKey(testHeader)
	if err == nil {
		t.Fatalf("Expected error. Got %v", err)
	}

	testHeader.Set("Authorization", "ApiKey")
	_, err = GetAPIKey(testHeader)
	if err == nil {
		t.Fatalf("Expected error. Got %v", err)
	}
}

func TestValidAuthHeader(t *testing.T) {
	testHeader := http.Header{}
	testHeader.Set("Authorization", "ApiKey testkey")

	_, err := GetAPIKey(testHeader)
	if err != nil {
		t.Fatal("Didn't expect error")
	}
}
