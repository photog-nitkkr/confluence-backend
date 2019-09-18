package teamConfluRoutes

import (
	"errors"
	"net/http"

	"protocol"
)

func writeTeamConflu(w http.ResponseWriter, r *http.Request) {
	// err := addTeamConflu(r)
	err := errors.New("h")

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
		Message: "Added team member successfully",
		Request: protocol.GetRequestObject(r),
		Data:    nil,
	}
	protocol.WriteResponseObject(w, r, responseObject, http.StatusOK)
	return
}

// func addTeamConfluMember(r *http.Request) error {
// 	var memberVar structs.Person
// 	err := getPersonObject(r, &memberVar)
// 	if err != nil {
// 		return err
// 	}

// 	errMemberInvalid := isPersonValid(memberVar)
// 	if errMemberInvalid != nil {
// 		return errMemberInvalid
// 	}

// 	errInAdding := developer.AddDeveloperInFireStore(developerVar)
// 	return errInAdding
// }

// func isDeveloperValid(developer structs.Developer) error {
// 	if developer.Name == "" {
// 		err := errors.New("Error in adding developer: Developer name can not be undefined")
// 		return err
// 	} else if developer.Email == "" {
// 		err := errors.New("Error in adding developer: Developer email ID can not be undefined")
// 		return err
// 	}
// 	return nil
// }

// func getDeveloperObject(r *http.Request, developer *structs.Developer) error {
// 	decoder := json.NewDecoder(r.Body)
// 	err := decoder.Decode(&developer)

// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
