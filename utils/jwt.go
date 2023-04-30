package utils

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	ROLE_STUDENT   = "student"
	ROLE_PROFESSOR = "professor"
)

type JwtClaims struct {
	UserID int    `json:"sub"`
	Role   string `json:"role"`

	jwt.StandardClaims
}

func GenerateJwt(userID int, role string) (string, error) {
	// Set the expiration time of the token
	expiresAt := time.Now().Add(time.Hour * 24 * 30 * 6).Unix()
	// expiresAt := time.Now().Add(-time.Hour).Unix()

	// Set the claims for the JWT token
	claims := JwtClaims{
		UserID: userID,
		Role:   role,
		StandardClaims: jwt.StandardClaims{
			// Subject:   fmt.Sprint(id),
			ExpiresAt: expiresAt,
		},
	}

	// Create a new JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)

	// Sign the token using the secret key
	secretKey := os.Getenv("JWT_SECRET_KEY")
	signedToken, _ := token.SignedString([]byte(secretKey))

	return signedToken, nil
}

func ValidateJwt(tokenString string) (JwtClaims, error) {
	claims := JwtClaims{}

	key := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	}

	_, err := jwt.ParseWithClaims(tokenString, &claims, key)
	return claims, err
}
