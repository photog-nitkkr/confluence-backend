package eventDescRoutes

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"../../common/structs"
	. "../../db"
	"../../protocol"
	"cloud.google.com/go/firestore"
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

	errEVentInvalid := isEventValid(event)
	if errEVentInvalid != nil {
		return errEVentInvalid
	}

	errInAdding := addEventInFireStore(event)
	if err != nil {
		return errInAdding
	}
	return nil
}

func isEventValid(event structs.Event) error {
	if event.Category == "" {
		err := errors.New("Error in adding event: Category value can not be undefined")
		return err
	} else if event.Name == "" {
		err := errors.New("Error in adding event: Event name can not be undefined")
		return err
	} else if event.DisplayName == "" {
		err := errors.New("Error in adding event: Event display name can not be undefined")
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

func addEventInFireStore(event structs.Event) error {
	fireStoreClient := GetFirestore()

	err := addEventInEventDesc(event, fireStoreClient)
	if err != nil {
		return err
	}
	err = addEventInEventName(event, fireStoreClient)
	return err
}

func addEventInEventDesc(event structs.Event, fireStoreClient *firestore.Client) error {
	_, err := fireStoreClient.Collection("events").Doc("eventDesc").Collection(event.Category).Doc(event.Name).Set(context.Background(), event)
	return err
}

func addEventInEventName(event structs.Event, fireStoreClient *firestore.Client) error {
	newEvent := structs.Event{
		Name:        event.Name,
		DisplayName: event.DisplayName,
		Category:    event.Category,
	}
	_, err := fireStoreClient.Collection("events").Doc("eventsName").Collection(event.Category).Doc(event.Name).Set(context.Background(), newEvent)
	return err
}
