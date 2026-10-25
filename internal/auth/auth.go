package auth

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
	"strings"
)

func GetBearerToken(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", fmt.Errorf("Authorization header is missing")
	}

	const bearerPrefix = "Bearer "
	if !strings.HasPrefix(authHeader, bearerPrefix) {
		return "", fmt.Errorf("Invalid authorization format")
	}

	token := strings.TrimPrefix(authHeader, bearerPrefix)

	return token, nil
}

// generate a random 256-bit hex-encoded string
func MakeRefreshToken() (string, error) {
	randomData := make([]byte, 32)
	_, err := rand.Read(randomData)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(randomData), nil
}
