package user

import (
	. "common/structs"
	"context"
	"db"

	"google.golang.org/api/iterator"
)

func checkIfRegistered(category string, event string, sub string) (bool, error) {
	firestoreClient := db.GetFirestore()

	doc, err := firestoreClient.Collection("userEvents").Doc(sub).Collection(category).Doc(event).Get(context.Background())
	if err != nil {
		return false, err
	}
	if doc == nil {
		return false, nil
	}
	return true, nil
}

func getUserEvents(sub string) (*[]Category, error) {
	var categories []Category

	firestoreClient := db.GetFirestore()

	collectionIter := firestoreClient.Collection("userEvents").Doc(sub).Collections(context.Background())
	for {
		collRef, err := collectionIter.Next()
		if err == iterator.Done {
			return &categories, nil
		}
		if err != nil {
			return nil, err
		}

		category, err := GetEventsForCategory(collRef.ID, sub)
		if err != nil {
			return nil, err
		}
		categories = append(categories, *category)

		// var events []Event
		// eventIterator := fireStoreClient.Collection("userEvents").Doc(sub).Collection(collRef.ID).Documents(context.Background())

		// for {
		// 	doc, err := eventIterator.Next()

		// 	if err == iterator.Done {
		// 		return &events, nil
		// 	}
		// 	if err != nil {
		// 		return nil, err
		// 	}

		// 	var event Event
		// 	errInConversion := ConvertToEventObject(doc, &event)

		// 	if errInConversion != nil {
		// 		return nil, errInConversion
		// 	}
		// 	events = append(events, event)
		// }
	}

}

func GetEventsForCategory(categoryName string, sub string) (*Category, error) {
	events, err := GetAllEventsForCategory(categoryName, sub)

	if err != nil {
		return nil, err
	}

	category := Category{
		Name:   categoryName,
		Events: *events,
	}
	return &category, nil
}

func GetAllEventsForCategory(categoryName string, sub string) (*[]Event, error) {
	firestoreClient := db.GetFirestore()

	eventIterator := firestoreClient.Collection("userEvents").Doc(sub).Collection(categoryName).Documents(context.Background())
	if eventIterator == nil {
		return nil, nil
	}

	var events []Event

	for {
		doc, err := eventIterator.Next()

		if err == iterator.Done {
			return &events, nil
		}

		if err != nil {
			return nil, err
		}

		event := Event{Name: doc.Data()}
		events = append(events, event)
	}

	return &events, nil //Unreachable Code
}
