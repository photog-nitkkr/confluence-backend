package categoryRoutes

import (
	"../../category"
	"../../protocol"
	"fmt"
	"net/http"
)

func readCategory(w http.ResponseWriter, r *http.Request) {
	category := r.URL.Query()["category"]
	if category == nil {
		returnAllCategory(w, r)
		return
	} else {
		fmt.Println(r.URL.Query()["category"][0])
		return
	}
}

func returnAllCategory(w http.ResponseWriter, r *http.Request) {
	categories, err := category.GetAllCategory()

	if err != nil {
		responseObject := protocol.Response{
			Success: false,
			Message: "Error in Getting Categories",
			Request: protocol.GetRequestObject(r),
			Data:    nil,
		}
		protocol.WriteResponseObject(w, r, responseObject, http.StatusInternalServerError)
		return
	}

	responseObject := protocol.Response{
		Success: true,
		Message: "Giving All Categories",
		Request: protocol.GetRequestObject(r),
		Data:    *categories,
	}
	protocol.WriteResponseObject(w, r, responseObject, http.StatusOK)
	return
}
