package eventNames

import (
	"../../db"
	"context"
	"errors"
	"google.golang.org/api/iterator"
)

type eventNames struct {
	Category string `json:"category"`
	Events []string `json:"events"`
}

func GetEventByCategory(category string) (*eventNames, error) {
	firestoreClient := db.GetFirestore()

	docIterator := firestoreClient.Collection("events").Doc("eventsName").Collection(category).Documents(context.Background())
	var arrayEvents []string

	for {
		doc, err := docIterator.Next()

		if err == iterator.Done {
			allEvents := eventNames{Category:category, Events: arrayEvents}
			return &allEvents, nil
		}
		if err != nil {
			return nil, err
		}else if doc!=nil && doc.Data() != nil && doc.Data()["name"] != nil {
			data := doc.Data()
			arrayEvents = append(arrayEvents, data["name"].(string))
		} else {
			return nil, errors.New("server internal error: no doc data for: " + doc.Ref.ID)
		}
	}
}
