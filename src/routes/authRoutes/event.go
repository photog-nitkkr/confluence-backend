package authRoutes

import (
	"errors"
	"fmt"
	. "net/http"
	"protocol"
	. "user"
)

func eventRegister(w ResponseWriter, r *Request) {
	err := eventRegisterUtil(r)
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

func eventRegisterUtil(r *Request) error {
	user, err := isAuthenticated(r)
	if err != nil {
		return err
	}

	err, category, event := parseEventCategoryAndName(r)
	if err != nil {
		return err
	}
	fmt.Println(category)
	fmt.Println(event)
	fmt.Println(user.Sub)
	err = AddUserEvent(category, event, user.Sub, *user)
	return err
}

func parseEventCategoryAndName(r *Request) (error, string, string) {
	category := r.URL.Query()["category"]
	if category == nil {
		return errors.New("nil Category"), "", ""
	}

	eventName := r.URL.Query()["event"]
	if eventName == nil {
		return errors.New("nil event"), "", ""
	}
	return nil, category[0], eventName[0]
}
