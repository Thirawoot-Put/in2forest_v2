package pkg

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func SignToken(sub uint, role string, exp time.Time) (string, error) {
	var secret = []byte("your-secret-key")

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  sub,
		"role": role,
		"iss":  "IN2FOREST",
		"exp":  exp.Unix(),
		"iat":  time.Now().Unix(),
	})

	return claims.SignedString(secret)
}
