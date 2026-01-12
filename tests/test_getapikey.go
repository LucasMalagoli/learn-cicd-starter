package tests

import (
	"net/http"
	"testing"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name           string
		headers        http.Header
		expectedAPIKey string
		expectError    bool
		errorMsg       string
	}{
		{
			name:           "valid api key",
			headers:        http.Header{"Authorization": []string{"ApiKey 1234567890"}},
			expectedAPIKey: "1234567890",
			expectError:    false,
		},
		{
			name:           "missing authorization header",
			headers:        http.Header{},
			expectedAPIKey: "",
			expectError:    true,
			errorMsg:       auth.ErrNoAuthHeaderIncluded.Error(),
		},
		{
			name:           "malformed authorization header - wrong prefix",
			headers:        http.Header{"Authorization": []string{"Bearer 1234567890"}},
			expectedAPIKey: "",
			expectError:    true,
			errorMsg:       "malformed authorization header",
		},
		{
			name:           "malformed authorization header - missing key",
			headers:        http.Header{"Authorization": []string{"ApiKey"}},
			expectedAPIKey: "",
			expectError:    true,
			errorMsg:       "malformed authorization header",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			apiKey, err := auth.GetAPIKey(test.headers)
			if test.expectError {
				if err == nil {
					t.Errorf("expected error, got nil")
				} else if err.Error() != test.errorMsg {
					t.Errorf("expected error message %q, got %q", test.errorMsg, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if apiKey != test.expectedAPIKey {
					t.Errorf("expected api key %q, got %q", test.expectedAPIKey, apiKey)
				}
			}
		})
	}
}
