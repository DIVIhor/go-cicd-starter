package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	headerName := "Authorization"
	key := "someToken87Gu21E"

	tests := []struct {
		name        string
		headers     http.Header
		expectedKey string
		expectedErr bool
	}{
		{
			name:        "No auth header",
			headers:     http.Header{"Auth": []string{key}},
			expectedKey: "",
			expectedErr: true,
		},
		{
			name:        "No bearer",
			headers:     http.Header{headerName: []string{key}},
			expectedKey: "",
			expectedErr: true,
		},
		{
			name:        "Wrong bearer",
			headers:     http.Header{headerName: []string{"Bearer " + key}},
			expectedKey: "",
			expectedErr: true,
		},
		{
			name:        "Clean API key",
			headers:     http.Header{headerName: []string{"ApiKey " + key}},
			expectedKey: key,
			expectedErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			apiKey, err := GetAPIKey(test.headers)
			if (err != nil) != test.expectedErr {
				t.Errorf("GetAPIKey() error = %v, expected error: %v", err, test.expectedErr)
				return
			}
			if apiKey != test.expectedKey {
				t.Errorf("GetAPIKey() key = %v, expected key: %v", err, test.expectedKey)
			}
		})
	}
}
