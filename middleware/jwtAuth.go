package middleware

import (
	"time"

	"github.com/dragranzer/capstone-BE-FGD/config"
	jwt "github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func CreateToken(userId int, name string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userId,
		"name":    name,
		"exp":     time.Now().Add(time.Hour * 2).Unix(),
	}
	tokenWithClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenWithClaims.SignedString([]byte(config.ENV.JWT_SECRET))
	return token, err
}

func ExtractClaim(e echo.Context) (claims map[string]interface{}) {
	user := e.Get("user").(*jwt.Token)

	if user.Valid {
		claims = user.Claims.(jwt.MapClaims)
	}

	return
}
