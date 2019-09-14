package categoryRoutes

import (
	"encoding/json"
	"errors"
	"net/http"

	"../../category"
	"../../common/structs"
	"../../protocol"
)

func writeCategory(w http.ResponseWriter, r *http.Request) {
	err := addCategory(r)

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
		Message: "Added category successfully",
		Request: protocol.GetRequestObject(r),
		Data:    nil,
	}
	protocol.WriteResponseObject(w, r, responseObject, http.StatusOK)
	return
}

func addCategory(r *http.Request) error {
	var categoryVar structs.Category
	err := getCategoryObject(r, &categoryVar)
	if err != nil {
		return err
	}

	errCategoryInvalid := isCategoryValid(categoryVar)
	if errCategoryInvalid != nil {
		return errCategoryInvalid
	}

	errInAdding := category.AddCategoryInFireStore(categoryVar)
	return errInAdding
}

func isCategoryValid(category structs.Category) error {
	if category.Name == "" {
		err := errors.New("Error in adding category: Category name can not be undefined")
		return err
	} else if category.DisplayName == "" {
		err := errors.New("Error in adding category: Category display name can not be undefined")
		return err
	}
	return nil
}

func getCategoryObject(r *http.Request, category *structs.Category) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&category)

	if err != nil {
		return err
	}
	return nil
}
