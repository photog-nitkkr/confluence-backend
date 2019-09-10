package event

import (
	. "../common/structs"
	"../db"
	"cloud.google.com/go/firestore"
	"context"
	"errors"
)

func GetEventName(categoryName string, eventName string) (*Event, error) {
	firestoreClient := db.GetFirestore()

	doc, err := firestoreClient.Collection("events").Doc("eventsName").Collection(categoryName).Doc(eventName).Get(context.Background())

	if err != nil {
		return nil, err
	}

	if doc.Data() == nil {
		return nil, errors.New("Internal Server Error / No Category")
	}

	var event Event

	errInCustomObject := convertToEventObject(doc, &event)

	if errInCustomObject != nil {
		return nil, errInCustomObject
	}

	return &event, nil
}

func convertToEventObject(firestoreDocument *firestore.DocumentSnapshot, event *Event) error {
	err := firestoreDocument.DataTo(event)

	if err != nil {
		return err
	}
	return nil
}
