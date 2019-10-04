package authRoutes

import (
	. "net/http"
	"protocol"
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

func getRegisteredEventsUtil(r *Request) (interface{}, error) {
	return nil, nil
}
