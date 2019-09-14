package categoryRoutes

import (
	"category"
	"protocol"
	. "net/http"
)

func readCategory(w ResponseWriter, r *Request) {
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

func returnAllCategory(w ResponseWriter, r *Request) {
	categories, err := category.GetAllCategory()

	if err != nil {
		responseObject := protocol.Response{
			Success: false,
			Message: "Error in Getting Categories",
			Request: protocol.GetRequestObject(r),
			Data:    nil,
		}
		protocol.WriteResponseObject(w, r, responseObject, StatusInternalServerError)
		return
	}

	responseObject := protocol.Response{
		Success: true,
		Message: "Giving All Categories",
		Request: protocol.GetRequestObject(r),
		Data:    *categories,
	}
	protocol.WriteResponseObject(w, r, responseObject, StatusOK)
	return
}

func returnCategories(w ResponseWriter, r *Request, categoryArray []string) {
	categories, err := category.GetCategories(categoryArray)

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
		Message: "Giving data for all required categories",
		Request: protocol.GetRequestObject(r),
		Data:    categories,
	}
	protocol.WriteResponseObject(w, r, responseObject, StatusOK)
	return
}
