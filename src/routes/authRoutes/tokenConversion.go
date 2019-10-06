package authRoutes

import (
	"common/structs"
	"fmt"
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

	expirationTime := time.Now().Add(2 * time.Hour)
	issueTime := time.Now()

	standardClaims := jwt.StandardClaims{
		IssuedAt: issueTime.Unix(),
		ExpiresAt: expirationTime.Unix(),
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
