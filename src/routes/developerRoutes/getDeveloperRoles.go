package developerRoutes

import (
	. "net/http"
	"protocol"

	. "person"
)

func getDeveloperRoles(w ResponseWriter, r *Request) {
	if len(r.URL.Query()) == 0 {
		returnDeveloperRoles(w, r)
		return
	} else {
		returnInvalidParamsError(w, r)
		return
	}
}

func returnDeveloperRoles(w ResponseWriter, r *Request) {
	developerRoles, err := GetNamesOfSubTeams("developers")

	if err != nil {
		responseObject := protocol.Response{
			Success: false,
			Message: "Error in Getting Developer Roles",
			Request: protocol.GetRequestObject(r),
			Data:    nil,
		}
		protocol.WriteResponseObject(w, r, responseObject, StatusInternalServerError)
		return
	}

	responseObject := protocol.Response{
		Success: true,
		Message: "Giving All Developer Roles",
		Request: protocol.GetRequestObject(r),
		Data:    *developerRoles,
	}
	protocol.WriteResponseObject(w, r, responseObject, StatusOK)
	return
}
