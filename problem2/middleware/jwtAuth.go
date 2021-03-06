package middleware

import (
	"problem2/config"
	"time"

	"github.com/golang-jwt/jwt"
)
func CreateToken(email , password string) (string, error) {
	claims := jwt.MapClaims{}
	claims["email"]= email
	claims["password"] = password
	claims["expired"] = time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.JwtSecret))
}

// func ExtractTokenUser(e echo.Context) int {
// 	user := e.Get("user").(*jwt.Token)
// 	if user.Valid {
// 		claims := user.Claims(jwt.MapClaims)
// 		password := claims["password"].(int)
// 		return password
// 	}
// 	return 0
// }