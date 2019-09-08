package categoryRoutes

import (
	"fmt"
	"net/http"
	"../../category"
	"../../protocol"
	"../../common"
)

func getCategory(w http.ResponseWriter, r *http.Request) {
	category := r.URL.Query()["category"]

	if category == nil {
		returnAllCategory(w)
	} else {
		fmt.Println(r.URL.Query()["category"][0])
	}
}

func returnAllCategory(w http.ResponseWriter) {
	categories, err := category.GetAllCategory()

	if err != nil {
		return
	}

	responseObject := protocol.ResponseProtocol{
		Success: true,
		Message: "Giving All Categories",
		Data:    *categories,
	}

	common.WriteResponseObject(w, responseObject, http.StatusOK)
}
