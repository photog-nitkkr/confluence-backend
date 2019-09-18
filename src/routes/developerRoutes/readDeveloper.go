package developerRoutes

import (
	. "net/http"
	"protocol"

	. "developer"
)

func readDeveloper(w ResponseWriter, r *Request) {
	if len(r.URL.Query()) == 0 {
		returnAllDevelopers(w, r)
		return
	} else {
		returnInvalidParamsError(w, r)
		return
	}
}

func returnInvalidParamsError(w ResponseWriter, r *Request) {
	responseObject := protocol.Response{
		Success: false,
		Message: "Invalid Parameters",
		Request: protocol.GetRequestObject(r),
		Data:    nil,
	}
	protocol.WriteResponseObject(w, r, responseObject, StatusBadGateway)
	return
}

func returnAllDevelopers(w ResponseWriter, r *Request) {
	developers, err := GetAllDevelopers()

	if err != nil {
		responseObject := protocol.Response{
			Success: false,
			Message: "Error in Getting Developers",
			Request: protocol.GetRequestObject(r),
			Data:    nil,
		}
		protocol.WriteResponseObject(w, r, responseObject, StatusInternalServerError)
		return
	}

	responseObject := protocol.Response{
		Success: true,
		Message: "Giving All Developers",
		Request: protocol.GetRequestObject(r),
		Data:    *developers,
	}
	protocol.WriteResponseObject(w, r, responseObject, StatusOK)
	return
}
