package pkg

import (
	"thirawoot/in2forest_shop_backend/internal/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func SignToken(sub uint, role string, exp time.Time) (string, error) {
	env := config.LoadEnv()
	secret := []byte(env.Secret)

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  sub,
		"role": role,
		"iss":  "IN2FOREST",
		"exp":  exp.Unix(),
		"iat":  time.Now().Unix(),
	})

	return claims.SignedString(secret)
}
