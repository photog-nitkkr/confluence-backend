package authRoutes

import (
	"common/structs"
	"errors"
	jwt "github.com/dgrijalva/jwt-go"
)

type MyCustomToken struct {
	Email string `json:"email,omitempty"`
	Sub string `json:"sub,omitempty"`
	Name string `json:"name,omitempty"`
	ProfilePic string `json:"picture,omitempty"`
	ErrorDescription string `json:"errorDescription,omitempty"`
	College string `json:"college,omitempty"`
	ContactNumber string `json:"contactNumber,omitempty"`
	Year string `json:"year,omitempty"`
	OnBoard bool `json:"onBoard"`
	jwt.StandardClaims
}

var secret = "Mudit Is Good Boy"

func JwtEncode(tokenInfo *structs.TokenInfo) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, fromTokenInfoToMyCustomToken(*tokenInfo))

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func JWTVerify(tokenString string) (*structs.TokenInfo, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomToken{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}
	if token.Valid {

		claims, ok := token.Claims.(*MyCustomToken)
		if !ok {
			return nil, errors.New("problem in converting to custom Token")
		}
		token := fromCustomTokenToInfoToken(*claims)
		return &token, nil

	} else if ve, ok := err.(*jwt.ValidationError); ok {

		if ve.Errors & jwt.ValidationErrorMalformed != 0 {
			return nil, errors.New("that's not even a token")
		} else if ve.Errors & ( jwt.ValidationErrorExpired | jwt.ValidationErrorNotValidYet) != 0 {
			return nil, errors.New("token Expired")
		} else {
			return nil, err
		}

	} else {
		return nil, err
	}
}