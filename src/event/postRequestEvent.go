package event

import (
	"common/structs"
	. "db"
	"context"
)

func AddEventInFireStoreUtil(event structs.Event) error {
	err := addEventInFirestore(event, "eventDesc")
	if err != nil {
		return err
	}

	err = addEventInFirestore(getEventForEventsName(event), "eventsName")
	return err
}

func getEventForEventsName(event structs.Event) structs.Event {
	return structs.Event{
		Name:         event.Name,
		Category:     event.Category,
	}
}

func addEventInFirestore(event structs.Event, docName string) error {
	firestoreClient := GetFirestore()

	_, err := firestoreClient.Collection("events").Doc(docName).Collection(event.Category).Doc(event.Name).Set(context.Background(), event)
	return err
}