package authRoutes

import (
	"common/structs"
	"github.com/dgrijalva/jwt-go"
)

func JwtEncode(tokenInfo *structs.TokenInfo) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"token": "Hi",
	})

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}