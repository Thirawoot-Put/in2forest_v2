package authjwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(sub uint, role string, exp time.Time) (string, error) {
	var secret = []byte("your-secret-key")

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  sub,
		"role": role,
		"iss":  "IN2FOREST",
		"exp":  exp.Unix(),
		"iat":  time.Now().Unix(),
	})

	tokenString, err := claims.SignedString(secret)
	if err != nil {
		return "", fmt.Errorf("Sign token error: %v", err.Error())
	}

	return tokenString, nil
}
