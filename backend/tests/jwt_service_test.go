package tests

import (
	"os"
	"pawAPIbackend/service"
	"strconv"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

type jwtCustomClaim struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

func TestCreateToken(t *testing.T) {

	err := godotenv.Load("../../.env")

	userID := uint64(123)
	expectedIssuer := "ISSUER BOOK API"

	// Set the JWT_SECRET environment variable for testing
	err = os.Setenv("JWT_SECRET", "your-secret-key")
	if err != nil {
		t.Fatalf("Failed to set JWT_SECRET environment variable: %v", err)
	}

	// Call the CreateToken function to generate the token
	token, err := service.CreateToken(userID)

	// Verify that the token creation was successful
	if err != nil {
		t.Errorf("Failed to create token: %v", err)
	}

	// Parse the token to verify its claims
	parsedToken, err := jwt.ParseWithClaims(token, &jwtCustomClaim{}, func(token *jwt.Token) (interface{}, error) {
		// Retrieve the JWT_SECRET from the environment variable
		secret := os.Getenv("JWT_SECRET")
		return []byte(secret), nil
	})

	// Verify that the token parsing was successful
	if err != nil {
		t.Errorf("Failed to parse token: %v", err)
	}

	// Verify the token claims
	claims, ok := parsedToken.Claims.(*jwtCustomClaim)
	if !ok {
		t.Errorf("Failed to parse custom claims from token")
	}

	// Verify the token issuer
	if claims.Issuer != expectedIssuer {
		t.Errorf("Unexpected token issuer. Expected: %s, Got: %s", expectedIssuer, claims.Issuer)
	}

	// Verify the token expiration
	expirationTime := time.Unix(claims.StandardClaims.ExpiresAt, 0)
	expectedExpirationTime := time.Now().Add(30 * time.Minute)
	if expirationTime.After(expectedExpirationTime) {
		t.Errorf("Token expiration time is later than expected. Expected: %s, Got: %s", expectedExpirationTime, expirationTime)
	}

	// Verify the token user ID
	parsedUserID, err := strconv.ParseUint(claims.UserID, 10, 64)
	if err != nil {
		t.Errorf("Failed to parse user ID from token: %v", err)
	}
	if parsedUserID != userID {
		t.Errorf("Unexpected user ID in token. Expected: %d, Got: %d", userID, parsedUserID)
	}
}
