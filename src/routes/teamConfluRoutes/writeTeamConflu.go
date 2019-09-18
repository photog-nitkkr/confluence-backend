package teamConfluRoutes

import (
	"net/http"

	"protocol"

	. "person"
)

func sendJSONAfterWrite(w http.ResponseWriter, r *http.Request, role string) {
	err := AddPerson(r, "teamConflu", role)

	if err != nil {
		responseObject := protocol.Response{
			Success: false,
			Message: err.Error(),
			Request: protocol.GetRequestObject(r),
			Data:    nil,
		}
		protocol.WriteResponseObject(w, r, responseObject, http.StatusInternalServerError)
		return
	}

	responseObject := protocol.Response{
		Success: true,
		Message: "Added team Conflu member successfully",
		Request: protocol.GetRequestObject(r),
		Data:    nil,
	}
	protocol.WriteResponseObject(w, r, responseObject, http.StatusOK)
	return
}

func writeTeamConflu(w http.ResponseWriter, r *http.Request) {
	roleParam := r.URL.Query()["role"]
	var role string
	if len(roleParam) == 1 {
		role = roleParam[0]
	} else {
		returnInvalidParamsError(w, r)
		return
	}
	sendJSONAfterWrite(w, r, role)
	return
}
