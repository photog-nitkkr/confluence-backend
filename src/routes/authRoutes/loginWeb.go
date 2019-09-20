package authRoutes

import (
	"common/structs"
	"errors"
	"fmt"
	"io/ioutil"
	. "net/http"
)

func login(w ResponseWriter, r *Request) {
	handleLogin(r)
}

func handleLogin(r *Request) (interface{}, error) {
	tokenInfo, err := getUserTokenInfo(r)
	if err != nil {
		return nil, err
	}

	tokenInfoObject := *tokenInfo
	if tokenInfoObject.ErrorDescription != "" {
		return nil, errors.New("Error is: " + tokenInfoObject.ErrorDescription)
	}

	name := tokenInfoObject.Name
	email := tokenInfoObject.Email
	profilePic := tokenInfoObject.ProfilePic
	sub := tokenInfoObject.Sub
	fmt.Println(name, email, profilePic, sub)
	return nil, nil
}


func getUserTokenInfo(r *Request) (*structs.TokenInfo, error) {
	googleUrl := "https://www.googleapis.com/oauth2/v3/tokeninfo?id_token="
	idToken, err := structs.GetIdToken(r)

	if err != nil {
		return nil, err
	}
	finalUrl := googleUrl + idToken
	resp, err := Get(finalUrl)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	tokenInfo , err := structs.GetTokenInfo(body)
	if err != nil {
		return nil, err
	}
	fmt.Println(tokenInfo)
	return tokenInfo, nil
}