package auth

import (
	"net/http"
	"testing"
)

func TestGetBearerToken(t *testing.T) {
	tests := []struct {
		name          string
		headers       http.Header
		expectedToken string
		expectError   bool
	}{
		{
			name: "Valid Bearer Token",
			headers: http.Header{
				"Authorization": []string{"Bearer valid_token_string_123"},
			},
			expectedToken: "valid_token_string_123",
			expectError:   false,
		},
		{
			name:          "Lack Authorization Header",
			headers:       http.Header{},
			expectedToken: "",
			expectError:   true,
		},
		{
			name: "Prefix is not Bearer",
			headers: http.Header{
				"Authorization": []string{"Cat dXNlcjpwYXNzd29yZA=="},
			},
			expectedToken: "",
			expectError:   true,
		},
		{
			name: "Lack Prefix",
			headers: http.Header{
				"Authorization": []string{"dXNlcjpwYXNzd29yZA=="},
			},
			expectedToken: "",
			expectError:   true,
		},
		{
			name: "Miss whitespace between Bearer and Token",
			headers: http.Header{
				"Authorization": []string{"BearerdXNlcjpwYXNzd29yZA=="},
			},
			expectedToken: "",
			expectError:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, err := GetBearerToken(tt.headers)

			// 檢查是否符合預期錯誤
			if tt.expectError {
				if err == nil {
					t.Errorf("expected error, but got nil")
				}
				return
			}

			// 檢查非錯誤情況
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			// 檢查 token 是否真的正確提取
			if token != tt.expectedToken {
				t.Errorf("expected token %s, but got %s", tt.expectedToken, token)
			}
		})
	}
}
