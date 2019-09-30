package eventDescRoutes

import (
	"encoding/json"
	"errors"
	"net/http"

	"common/structs"
	"event"
	"protocol"
)

func writeEvent(w http.ResponseWriter, r *http.Request) {
	err := addEvent(r)

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
		Message: "Added event successfully",
		Request: protocol.GetRequestObject(r),
		Data:    nil,
	}
	protocol.WriteResponseObject(w, r, responseObject, http.StatusOK)
	return
}

func addEvent(r *http.Request) error {
	var eventVar structs.Event
	err := getEventObject(r, &eventVar)
	if err != nil {
		return err
	}

	errEVentInvalid := isEventValid(eventVar)
	if errEVentInvalid != nil {
		return errEVentInvalid
	}

	errInAdding := event.AddEventInFireStoreUtil(eventVar)
	return errInAdding
}

func isEventValid(event structs.Event) error {
	if event.Category == "" {
		err := errors.New("Error in adding event: Category value can not be undefined")
		return err
	} else if event.Name == "" {
		err := errors.New("Error in adding event: Event name can not be undefined")
		return err
	}
	return nil
}

func getEventObject(r *http.Request, event *structs.Event) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&event)

	if err != nil {
		return err
	}
	return nil
}