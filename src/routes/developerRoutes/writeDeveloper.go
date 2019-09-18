package developerRoutes

import (
	"net/http"

	. "person"
	"protocol"
)

func writeDeveloper(w http.ResponseWriter, r *http.Request) {
	roleParam := r.URL.Query()["role"]
	var role string
	if roleParam == nil {
		role = "webapp"
	} else if len(roleParam) == 1 {
		role = roleParam[0]
	} else {
		returnInvalidParamsError(w, r)
		return
	}

	err := AddPerson(r, "developers", role)

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
		Message: "Added developer successfully",
		Request: protocol.GetRequestObject(r),
		Data:    nil,
	}
	protocol.WriteResponseObject(w, r, responseObject, http.StatusOK)
	return
}
