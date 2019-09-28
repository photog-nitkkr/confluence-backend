package authRoutes

import (
	"common/structs"
	"errors"
	"io/ioutil"
	. "net/http"
	"protocol"
	"user"
)

func login(w ResponseWriter, r *Request) {
	token, err := handleLogin(r)
	if err != nil {
		responseObject := protocol.Response{
			Success: false,
			Message: err.Error(),
			Request: protocol.GetRequestObject(r),
			Data:    nil,
		}
		protocol.WriteResponseObject(w, r, responseObject, StatusBadRequest)
		return
	}

	responseObject := protocol.Response{
		Success: true,
		Message: "User Logged In",
		Request: protocol.GetRequestObject(r),
		Data:    token,
	}
	protocol.WriteResponseObject(w, r, responseObject, StatusOK)
	return
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

	 user, err := user.HandleFirestoreUser(&tokenInfoObject)
	 if err != nil {
	 	return nil, err
	 }
	 encodedToken, err := JwtEncode(user)
	if err != nil {
		return nil, err
	}
	return encodedToken, nil
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