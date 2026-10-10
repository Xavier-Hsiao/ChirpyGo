package auth

import (
	"testing"
)

func TestHashPasswordAndCheck(t *testing.T) {
	password := "mySecretPassword123"

	// 1. Test hashing
	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("unexpected error hashing password: %v", err)
	}
	if hash == "" {
		t.Fatal("expected hash to be non-empty string")
	}

	// 2. Test correct password match
	match, err := CheckPasswordHash(password, hash)
	if err != nil {
		t.Fatalf("unexpected error checking password hash: %v", err)
	}
	if !match {
		t.Errorf("expected password to match hash, but got false")
	}

	// 3. Test wrong password
	wrongPassword := "wrongPassword"
	match, err = CheckPasswordHash(wrongPassword, hash)
	if err != nil {
		t.Fatalf("unexpected error checking wrong password: %v", err)
	}
	if match {
		t.Errorf("expected wrong password to not match, but got true")
	}
}

func TestPasswordSalting(t *testing.T) {
	password := "samePassword"

	// Argon2id generates a random salt each time, so hashes should be unique
	hash1, err := HashPassword(password)
	if err != nil {
		t.Fatalf("unexpected error hashing password (1): %v", err)
	}

	hash2, err := HashPassword(password)
	if err != nil {
		t.Fatalf("unexpected error hashing password (2): %v", err)
	}

	if hash1 == hash2 {
		t.Errorf("expected hashes for the same password to be different due to random salt, but got identical hashes")
	}

	// Both hashes should still validate the same password
	match1, _ := CheckPasswordHash(password, hash1)
	match2, _ := CheckPasswordHash(password, hash2)
	if !match1 || !match2 {
		t.Errorf("expected both unique hashes to validate the password")
	}
}

func TestCheckPasswordHash_InvalidHash(t *testing.T) {
	// Passing a malformed hash string should return an error
	invalidHash := "not-a-valid-argon2-hash"
	match, err := CheckPasswordHash("password", invalidHash)
	if err == nil {
		t.Errorf("expected error when verifying against an invalid hash string, got nil")
	}
	if match {
		t.Errorf("expected match to be false on error, got true")
	}
}

func TestEmptyPassword(t *testing.T) {
	// Empty password should still be hashable and verifiable
	hash, err := HashPassword("")
	if err != nil {
		t.Fatalf("unexpected error hashing empty password: %v", err)
	}

	match, err := CheckPasswordHash("", hash)
	if err != nil {
		t.Fatalf("unexpected error checking empty password: %v", err)
	}
	if !match {
		t.Errorf("expected empty password to match hash")
	}
}
