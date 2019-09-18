package developerRoutes

import (
	. "net/http"
	"protocol"

	// . "developer"
	. "person"
)

func readDeveloper(w ResponseWriter, r *Request) {
	role := r.URL.Query()["role"]
	if len(r.URL.Query()) == 0 && role == nil {
		returnAllDevelopers(w, r)
		return
	} else if role != nil {
		returnDevelopersForARole(w, r, role)
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
	developers, err := GetAllSubTeams("developers")

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

func returnDevelopersForARole(w ResponseWriter, r *Request, role []string) {
	teams, err := GetSubTeams("developers", role)
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
		Message: "Giving All Developers for given roles",
		Request: protocol.GetRequestObject(r),
		Data:    *teams,
	}
	protocol.WriteResponseObject(w, r, responseObject, StatusOK)
	return
}
