package utils

import (
	"os"
	"testing"

	"github.com/dgrijalva/jwt-go"
)

func TestGenerateJwt(t *testing.T) {
	userID := 123
	role := ROLE_STUDENT
	secretKey := "test-secret-key"

	os.Setenv("JWT_SECRET_KEY", secretKey)
	defer os.Unsetenv("JWT_SECRET_KEY")

	token, err := GenerateJwt(userID, role)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Parse the token to get the claims
	parsedToken, err := jwt.ParseWithClaims(token, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Check if the token is valid and has the expected claims
	claims, ok := parsedToken.Claims.(*JwtClaims)
	if !ok || !parsedToken.Valid || claims.UserID != userID || claims.Role != role {
		t.Errorf("Invalid token: %v", parsedToken)
	}
}

func TestValidateJwt(t *testing.T) {
	// Set the environment variable for JWT_SECRET_KEY
	os.Setenv("JWT_SECRET_KEY", "test-secret-key")
	defer os.Unsetenv("JWT_SECRET_KEY")

	// Generate a sample JWT token
	userID := 123
	role := ROLE_STUDENT
	tokenString, err := GenerateJwt(userID, role)
	if err != nil {
		t.Fatalf("Failed to generate JWT token: %v", err)
	}

	// Call the function to validate the JWT token
	claims, err := ValidateJwt(tokenString)
	if err != nil {
		t.Fatalf("Failed to validate JWT token: %v", err)
	}

	// Verify that the claims match the expected values
	if claims.UserID != userID {
		t.Errorf("UserID does not match. Expected: %d, Actual: %d", userID, claims.UserID)
	}
	if claims.Role != role {
		t.Errorf("role does not match. Expected: %s, Actual: %s", role, claims.Role)
	}
}
