package developerRoutes

import (
	"encoding/json"
	"errors"
	"net/http"

	"protocol"

	"common/structs"

	"developer"
)

func writeDeveloper(w http.ResponseWriter, r *http.Request) {
	err := addDeveloper(r)

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

func addDeveloper(r *http.Request) error {
	var developerVar structs.Developer
	err := getDeveloperObject(r, &developerVar)
	if err != nil {
		return err
	}

	errDeveloperInvalid := isDeveloperValid(developerVar)
	if errDeveloperInvalid != nil {
		return errDeveloperInvalid
	}

	errInAdding := developer.AddDeveloperInFireStore(developerVar)
	return errInAdding
}

func isDeveloperValid(developer structs.Developer) error {
	if developer.Name == "" {
		err := errors.New("Error in adding developer: Developer name can not be undefined")
		return err
	} else if developer.Email == "" {
		err := errors.New("Error in adding developer: Developer email ID can not be undefined")
		return err
	}
	return nil
}

func getDeveloperObject(r *http.Request, developer *structs.Developer) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&developer)

	if err != nil {
		return err
	}
	return nil
}
