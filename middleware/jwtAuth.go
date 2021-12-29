package middleware

import (
	"time"

	"github.com/dragranzer/capstone-BE-FGD/config"
	jwt "github.com/golang-jwt/jwt"
)

func CreateToken(userId int, name string) (string, error) {
	claims := jwt.MapClaims{
		"userid": userId,
		"name":   name,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	}
	tokenWithClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenWithClaims.SignedString([]byte(config.ENV.JWT_SECRET))
	return token, err
}
