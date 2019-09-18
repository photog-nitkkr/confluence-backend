package teamConfluRoutes

import (
	. "net/http"
	"protocol"

	. "person"
)

func getTeamConfluRoles(w ResponseWriter, r *Request) {
	if len(r.URL.Query()) == 0 {
		returnTeamConfluRoles(w, r)
		return
	} else {
		returnInvalidParamsError(w, r)
		return
	}
}

func returnTeamConfluRoles(w ResponseWriter, r *Request) {
	teamRoles, err := GetNamesOfSubTeams("teamConflu")

	if err != nil {
		responseObject := protocol.Response{
			Success: false,
			Message: "Error in Getting Team COnflu Roles",
			Request: protocol.GetRequestObject(r),
			Data:    nil,
		}
		protocol.WriteResponseObject(w, r, responseObject, StatusInternalServerError)
		return
	}

	responseObject := protocol.Response{
		Success: true,
		Message: "Giving All Team Conflu Roles",
		Request: protocol.GetRequestObject(r),
		Data:    *teamRoles,
	}
	protocol.WriteResponseObject(w, r, responseObject, StatusOK)
	return
}
