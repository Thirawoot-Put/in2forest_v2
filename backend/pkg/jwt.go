package pkg

import (
	"fmt"
	"thirawoot/in2forest_shop_backend/internal/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func SignToken(sub uint, role string, exp time.Time) (string, error) {
	fmt.Println("exp --> ", exp)
	fmt.Println("exp in unix --> ", exp.Unix())
	fmt.Println("now in unix --> ", time.Now().Unix())

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

func GetSubInToken(token interface{}) uint {
	jwtToken := token.(*jwt.Token)
	cliams := jwtToken.Claims.(jwt.MapClaims)
	id, ok := cliams["sub"].(float64)
	if !ok {
		return 0
	}

	return uint(id)
}
