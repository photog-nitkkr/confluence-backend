package eventDescRoutes

import (
	"../../common/structs"
	. "../../db"
	"../../protocol"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
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
	var event structs.Event
	err := getEventObject(r, &event)
	if err != nil {
		return err
	}
	
	errInAdding := addEventInFireStore(event)
	if err != nil {
		return errInAdding
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

func addEventInFireStore(event structs.Event) error {
	fireStoreClient := GetFirestore()

	_, err := fireStoreClient.Collection("events").Doc("eventDesc").Collection(event.Category).Doc(event.Name).Set(context.Background(), event)

	return err
}

