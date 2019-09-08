package categoryRoutes

import (
	"../../category"
	"../../protocol"
	"fmt"
	"net/http"
)

func readCategory(w http.ResponseWriter, r *http.Request) {
	category := r.URL.Query()["category"]
	protocol.GetQueryParams(r)
	if category == nil {
		returnAllCategory(w, r)
	} else {
		fmt.Println(r.URL.Query()["category"][0])
	}
}

func returnAllCategory(w http.ResponseWriter, r *http.Request) {
	categories, err := category.GetAllCategory()

	if err != nil {
		return
	}

	responseObject := protocol.Response{
		Success: true,
		Message: "Giving All Categories",
		Request: protocol.GetRequestObject(r),
		Data:    *categories,
	}

	protocol.WriteResponseObject(w, responseObject, http.StatusOK, r)
}
