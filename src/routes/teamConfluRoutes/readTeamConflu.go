package teamConfluRoutes

import (
	. "net/http"
	"protocol"

	. "developer"
)

func readTeamConflu(w ResponseWriter, r *Request) {
	role := r.URL.Query()["role"]
	if len(r.URL.Query()) == 0 && role == nil {
		returnTeamConflu(w, r)
		return
	} else if len(role) != 0 {
		returnTeamConfluForRole(w, r, role)
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

func returnTeamConflu(w ResponseWriter, r *Request) {
	teamMembers, err := GetAllTeam("teamConflu")

	if err != nil {
		responseObject := protocol.Response{
			Success: false,
			Message: "Error in Getting Team Conflu",
			Request: protocol.GetRequestObject(r),
			Data:    nil,
		}
		protocol.WriteResponseObject(w, r, responseObject, StatusInternalServerError)
		return
	}

	responseObject := protocol.Response{
		Success: true,
		Message: "Giving Team Conflu",
		Request: protocol.GetRequestObject(r),
		Data:    *teamMembers,
	}
	protocol.WriteResponseObject(w, r, responseObject, StatusOK)
	return
}

func returnTeamConfluForRole(w ResponseWriter, r *Request, role []string) {
	teamMembers, err := event.GetMembersForARole(role, "teamConflu")

	if err != nil {
		responseObject := protocol.Response{
			Success: false,
			Message: err.Error(),
			Request: protocol.GetRequestObject(r),
			Data:    nil,
		}
		protocol.WriteResponseObject(w, r, responseObject, StatusInternalServerError)
		return
	}

	responseObject := protocol.Response{
		Success: true,
		Message: "Giving members for given roles",
		Request: protocol.GetRequestObject(r),
		Data:    categories,
	}
	protocol.WriteResponseObject(w, r, responseObject, StatusOK)
	return
}
