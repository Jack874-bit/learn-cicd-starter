package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	header := http.Header{}
	header.Set("Authorization", "ApiKey HelloWorld")
	api, err := GetAPIKey(header)
	if err != nil {
		t.Fatalf("Failed to get header: %v", err)
	}
	if api != "HelloWorld" {
		t.Errorf("got %v, Expected HelloWorld", api)
	}
}

func TestGetAPIKey_No_Auth(t *testing.T) {
	header := http.Header{}
	header.Set("Auth", "ApiKey HelloWorld")
	api, err := GetAPIKey(header)
	if err == nil {
		t.Errorf("got %v, expected err", api)
	}

}
