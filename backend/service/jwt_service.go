package service

import (
	"errors"
	"os"
	"strconv"
	"strings"
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

func CreateToken(ID uint64) (string, error) {
	UserID := strconv.FormatUint(ID, 10)
	claims := &jwtCustomClaim{
		UserID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(30 * time.Minute).Unix(),
			Issuer:    "ISSUER BOOK API",
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	signature := os.Getenv("JWT_SECRET")

	t, err := token.SignedString([]byte(signature))
	return t, err
}

func ValidateToken(token string) (*jwt.Token, error) {
	token = strings.Replace(token, "Bearer ", "", -1)
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}

		err := godotenv.Load()
		if err != nil {
			panic("Error loading .env file")
		}
		signature := os.Getenv("JWT_SECRET")
		return []byte(signature), nil
	})
}
