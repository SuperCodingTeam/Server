package utility

import (
	"github.com/golang-jwt/jwt/v5"
)

func JWTDecode(JWTToken string) string {
	claims := jwt.MapClaims{}
	jwt.ParseWithClaims(JWTToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("your-secret-key"), nil
	})

	for key, value := range claims {
		if key == "userUUID" {
			return value.(string)
		}
	}

	return "Invalid Token"
}
