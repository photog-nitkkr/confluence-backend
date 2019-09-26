package authRoutes

import (
	"common/structs"
	"errors"
	"io/ioutil"
	. "net/http"
	"user"
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

	 user.HandleFirestoreUser(&tokenInfoObject)

	// db
	// sub - user
	// null = no user
	// create user in db {name, email, picture, sub, onBoard=false}
	// jwt details = {as above]

	// !null = user
	// user already get

	// details = { onBoard, college, phoneNo }

	// create jwt with details



	//encodedToken, err := JwtEncode(&tokenInfoObject)
	//if err != nil {
	//	return nil, err
	//}
	//return encodedToken, nil
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

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	tokenInfo , err := structs.GetTokenInfo(body)
	if err != nil {
		return nil, err
	}
	return tokenInfo, nil
}