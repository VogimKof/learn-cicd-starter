package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name        string
		header      string
		wantKey     string
		wantErr     bool
		expectedErr error
	}{
		{
			name:    "valid api key",
			header:  "ApiKey abc123",
			wantKey: "abc123",
		},
		{
			name:        "missing authorization header",
			header:      "",
			wantErr:     true,
			expectedErr: ErrNoAuthHeaderIncluded,
		},
		{
			name:    "wrong authorization type",
			header:  "Bearer abc123",
			wantErr: true,
		},
		{
			name:    "missing api key",
			header:  "ApiKey",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			headers := http.Header{}
				
			// fmt.Println("Running test ", tt.name)	

			if tt.header != "" {
				headers.Set("Authorization", tt.header)
			}

			key, err := GetAPIKey(headers)

			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected an error, got nil")
				}

				if tt.expectedErr != nil && err != tt.expectedErr {
					t.Fatalf("expected %v, got %v", tt.expectedErr, err)
				}

				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if key != tt.wantKey {
				t.Fatalf("expected key %q, got %q", tt.wantKey, key)
			}
		})
	}
}
