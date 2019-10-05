package event

import (
	"category"
	. "common/structs"
	"context"
	. "db"
	"errors"
	"strings"

	"google.golang.org/api/iterator"
)

func GetEventsForCategories(categoryNames []string, docName string) (*[]Category, error) {
	var categories []Category
	for _, categoryName := range categoryNames {
		category, err := GetEventsForCategory(categoryName, docName)

		if err != nil {
			return nil, err
		}
		categories = append(categories, *category)
	}
	return &categories, nil
}

func GetEventsForCategory(categoryName string, docName string) (*Category, error) {
	events, err := GetAllEventsForCategory(categoryName, docName)

	if err != nil {
		return nil, err
	}

	category, err := category.GetCategory(categoryName)
	if err != nil {
		return nil, err
	}

	category.Events = *events
	return category, nil
}

func GetAllEventsForCategory(categoryName string, docName string) (*[]Event, error) {
	firestoreClient := GetFirestore()

	eventIterator := firestoreClient.Collection("events").Doc(docName).Collection(strings.ToLower(categoryName)).Documents(context.Background())
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

		var event Event
		errInConversion := convertToEventObject(doc, &event)

		if errInConversion != nil {
			return nil, errInConversion
		}
		events = append(events, event)
	}

	return &events, nil //Unreachable Code
}

func GetEvents(categoryName string, eventNames []string, docName string) (*[]Event, error) {
	var events []Event

	for _, eventName := range eventNames {
		event, err := GetEvent(categoryName, eventName, docName)

		if err != nil {
			return nil, err
		}

		events = append(events, *event)
	}
	return &events, nil
}

func GetAllEventsForAllCategory(docName string) (*[]Category, error) {
	allCategory, err := category.GetAllCategory()

	if err != nil {
		return nil, err
	}

	var categories []Category
	for _, value := range *allCategory {
		category, err := GetEventsForCategory(value.Name, docName)
		if err != nil {
			return nil, err
		}
		categories = append(categories, *category)
	}
	return &categories, nil
}

func GetSpecificEvent(eventName string, docName string) (interface{}, error) {
	eventName = strings.ToLower(eventName)
	allEvents, err := GetAllEventsForAllCategory(docName)
	if err != nil {
		return nil, err
	}
	for _, category := range *allEvents {
		for _, event := range category.Events {
			if strings.ToLower(event.Name) == eventName {
				return event, nil
			}
		}
	}
	return nil, errors.New("Event Not Found")
}
