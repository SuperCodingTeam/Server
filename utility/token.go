package utility

import (
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func JWTDecode(JWTToken string) string {
	parts := strings.Split(JWTToken, ".")
	if len(parts) != 3 {
		panic(fmt.Sprintln("Invalid JWT: Token must have 3 parts (header, payload, signature)"))
	}
	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(JWTToken, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			panic(fmt.Sprintf("unexpected signing method: %v\n", token.Header["alg"]))
		}
		return []byte("your-secret-key"), nil
	})

	if err != nil {
		panic(fmt.Sprintf("failed to parse token: %v\n", err))
	}

	if !token.Valid {
		panic("Invalid Token\n")
	}

	if value, ok := claims["userUUID"]; ok {
		if userUUID, ok := value.(string); ok {
			return userUUID
		}
		panic("userUUID claim is not a string\n")
	}

	panic("userUUID claim not found in token\n")
}
