package categoryRoutes

import (
	"net/http"

	"../../category"
	"../../protocol"
)

func readCategory(w http.ResponseWriter, r *http.Request) {
	category := r.URL.Query()["category"]
	if len(r.URL.Query()) == 0 {
		returnAllCategory(w, r)
		return
	} else if category != nil {
		returnCategories(w, r, category)
		return
	} else {
		returnInvalidParamsError(w, r)
		return
	}
}

func returnInvalidParamsError(w http.ResponseWriter, r *http.Request) {
	responseObject := protocol.Response{
		Success: false,
		Message: "Invalid Parameters",
		Request: protocol.GetRequestObject(r),
		Data:    nil,
	}
	protocol.WriteResponseObject(w, r, responseObject, http.StatusBadGateway)
	return
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

func returnCategories(w http.ResponseWriter, r *http.Request, categoryArray []string) {
	categories, err := category.GetCategories(categoryArray)

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
		Message: "Giving data for all required categories",
		Request: protocol.GetRequestObject(r),
		Data:    categories,
	}
	protocol.WriteResponseObject(w, r, responseObject, http.StatusOK)
	return
}
