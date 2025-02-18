package authjwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func SignToken(sub string, role string) (string, error) {
	exp := time.Now().Add(time.Duration(24 * time.Hour))
	fmt.Println("expired at ----> ", exp)
	secret := []byte("supersecret")
	claims := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"sub":  sub,
		"role": role,
		"iss":  "IN2FOREST",
		"exp":  exp.Unix(),
		"iat":  time.Now().Unix(),
	})

	tokenString, err := claims.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
