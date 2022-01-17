package middlewares

import (
	"capstone/backend/config"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
)

func CreateToken(UserID uint, name string) (string, error) {
	claims := jwt.MapClaims()

	claims["userId"] = UserID
	claims["name"] = name
	claims["exp"] = time.Now().Add(time.Hour * 6).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(config.JWT_KEY))
}
func ExtractClaim(e echo.Context) (claims map[string]interface{}) {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims = user.Claims.(jwt.MapClaims)
	}
	return
}
