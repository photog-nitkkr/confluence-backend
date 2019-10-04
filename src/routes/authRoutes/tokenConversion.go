package authRoutes

import (
	"common/structs"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func fromTokenInfoToMyCustomToken(token structs.TokenInfo) MyCustomToken {
	var myCustomToken MyCustomToken
	myCustomToken.Name = token.Name
	myCustomToken.Email = token.Email
	myCustomToken.College = token.College
	myCustomToken.ContactNumber = token.ContactNumber
	myCustomToken.OnBoard = token.OnBoard
	myCustomToken.ErrorDescription = token.ErrorDescription
	myCustomToken.ProfilePic = token.ProfilePic
	myCustomToken.Sub = token.Sub
	myCustomToken.Year = token.Year

	standardClaims := jwt.StandardClaims{
		ExpiresAt: (time.Now().UnixNano() + int64((time.Duration(2*time.Hour))/time.Millisecond) ) / int64(time.Millisecond),
		Issuer: "Mudit Jain",
	}
	myCustomToken.StandardClaims = standardClaims
	return myCustomToken
}

func fromCustomTokenToInfoToken(myCustomToken MyCustomToken) structs.TokenInfo {
	var tokenInfo structs.TokenInfo
	tokenInfo.Name = myCustomToken.Name
	tokenInfo.Email = myCustomToken.Email
	tokenInfo.College = myCustomToken.College
	tokenInfo.ContactNumber = myCustomToken.ContactNumber
	tokenInfo.OnBoard = myCustomToken.OnBoard
	tokenInfo.ErrorDescription = myCustomToken.ErrorDescription
	tokenInfo.ProfilePic = myCustomToken.ProfilePic
	tokenInfo.Sub = myCustomToken.Sub
	tokenInfo.Year = myCustomToken.Year

	return tokenInfo
}
