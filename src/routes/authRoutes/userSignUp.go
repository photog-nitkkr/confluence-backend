package authRoutes

import (
	"common/structs"
	"encoding/json"
	"io/ioutil"
	. "net/http"
	"protocol"
	addUser "user"
)

func userSignUp(w ResponseWriter, r *Request) {
	err := userSignUpUtil(r)
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
		Message: "Registered Event Successfully",
		Request: protocol.GetRequestObject(r),
		Data:    nil,
	}
	protocol.WriteResponseObject(w, r, responseObject, StatusOK)
	return
}

func userSignUpUtil(r *Request) error {
	user, err := isAuthenticated(r)
	if err != nil {
		return err
	}

	err = parseUser(r, user)
	if err != nil {
		return err
	}

	err = addUser.AddUser(user)
	return err
}

func parseUser(r *Request, user *structs.TokenInfo) (error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	userInfo, err := getUserInfo(body)
	if err != nil {
		return err
	}
	updateUserInfo(user, userInfo)
	return nil
}

func getUserInfo(body []byte) (*structs.TokenInfo, error) {
	var tokenInfo structs.TokenInfo
	err := json.Unmarshal(body, &tokenInfo)

	if err != nil {
		return nil, err
	}
	return &tokenInfo, nil
}

func updateUserInfo(user *structs.TokenInfo, userInfo *structs.TokenInfo) {
	user.College = userInfo.College
	user.Year = userInfo.Year
	user.ContactNumber = userInfo.ContactNumber
	user.OnBoard = true
}
