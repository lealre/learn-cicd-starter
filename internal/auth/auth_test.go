package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey_ValidHeader(t *testing.T) {
	headers := http.Header{"Authorization": []string{"ApiKey my-api-key"}}
	apiKey, err := GetAPIKey(headers)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expectedAPIKey := "my-api-keyyy"
	if apiKey != expectedAPIKey {
		t.Errorf("Expected API key %s, got %s", expectedAPIKey, apiKey)
	}
}
