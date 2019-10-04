package authRoutes

import (
	"common/structs"
	. "net/http"
	"protocol"
	getEvent "user"
)

func getRegisteredEvents(w ResponseWriter, r *Request) {
	data, err := getRegisteredEventsUtil(r)
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
		Message: "Returning All Events",
		Request: protocol.GetRequestObject(r),
		Data:    data,
	}
	protocol.WriteResponseObject(w, r, responseObject, StatusOK)
	return
}

func getRegisteredEventsUtil(r *Request) ([]structs.Category, error) {
	user, err := isAuthenticated(r)
	if err != nil {
		return nil, err
	}
	events, err := getEvent.GetUserEvents(user.Sub)
	if err != nil {
		return nil, err
	}
	return *events, nil
}
