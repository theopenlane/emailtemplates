package emailtemplates

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddTokenToURL(t *testing.T) {
	tests := []struct {
		name          string
		baseURL       string
		token         string
		expectedURL   string
		expectedError bool
	}{
		{
			name:          "valid token",
			baseURL:       "https://example.com/verify",
			token:         "validtoken",
			expectedURL:   "https://example.com/verify?token=validtoken",
			expectedError: false,
		},
		{
			name:          "empty token",
			baseURL:       "https://example.com/verify",
			token:         "",
			expectedURL:   "",
			expectedError: true,
		},
		{
			name:          "invalid base URL",
			baseURL:       "://invalid-url",
			token:         "validtoken",
			expectedURL:   "",
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url, err := addTokenToURL(tt.baseURL, tt.token)
			if tt.expectedError {
				require.Error(t, err)
				require.Empty(t, url)

				return
			}

			require.NoError(t, err)
			require.Equal(t, tt.expectedURL, url)
		})
	}
}
