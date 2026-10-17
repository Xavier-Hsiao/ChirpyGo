package auth

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestMakeAndValidateJWT_Valid(t *testing.T) {
	userID := uuid.New()
	secret := "my-secret-key"
	expiresIn := time.Hour

	tokenString, err := MakeJWT(userID, secret, expiresIn)
	if err != nil {
		t.Fatalf("unexpected error making JWT: %v", err)
	}

	if tokenString == "" {
		t.Fatalf("expected non-empty token string")
	}

	parsedUserID, err := ValidateJWT(tokenString, secret)
	if err != nil {
		t.Fatalf("unexpected error validating JWT: %v", err)
	}

	if parsedUserID != userID {
		t.Errorf("expected user ID %v, but got %v", userID, parsedUserID)
	}
}

func TestValidateJWT_Expired(t *testing.T) {
	userID := uuid.New()
	secret := "my-secret-key"
	// 負數代表過去時間 = 過期
	expiresIn := -time.Hour

	tokenString, err := MakeJWT(userID, secret, expiresIn)
	if err != nil {
		t.Fatalf("unexpected error making expired JWT: %v", err)
	}

	// 驗證過期的 Token，預期報錯
	_, err = ValidateJWT(tokenString, secret)
	if err == nil {
		t.Errorf("expected error validating expired token, but got nil")
	}
}

func TestValidateJWT_WrongSecret(t *testing.T) {
	userID := uuid.New()
	secret := "my-secret-key"
	wrongSecret := "wrong-secret"
	expiresIn := time.Hour

	tokenString, err := MakeJWT(userID, secret, expiresIn)
	if err != nil {
		t.Fatalf("unexpected error making JWT: %v", err)
	}

	// 故意用錯誤的密鑰驗證，預期會報錯
	_, err = ValidateJWT(tokenString, wrongSecret)
	if err == nil {
		t.Errorf("expected error validating token with wrong secret, but got nil")
	}
}

func TestValidateJWT_InvalidTokenString(t *testing.T) {
	invalidToken := "this.is.an.invalid.token"
	secret := "my-secret-key"

	// 預期報錯
	_, err := ValidateJWT(invalidToken, secret)
	if err == nil {
		t.Errorf("expected error validating malformed token, but got nil")
	}
}
