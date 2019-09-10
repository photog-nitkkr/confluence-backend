package event

import (
	"../category"
	. "../common/structs"
	. "../db"
	"context"
	"fmt"
	"google.golang.org/api/iterator"
)

func GetEventsForCategories(categoryNames []string) (*[]Category, error) {
	var categories []Category
	for _, categoryName := range categoryNames {
		category, err := GetEventsForCategory(categoryName)
		fmt.Println(categoryName)

		if err != nil {
			return nil, err
		}
		categories = append(categories, *category)
	}
	return &categories, nil
}

func GetEventsForCategory(categoryName string) (*Category, error) {
	events, err := GetAllEventsForCategory(categoryName)

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

func GetAllEventsForCategory(categoryName string) (*[]Event, error) {
	firestoreClient := GetFirestore()
	fmt.Println(categoryName)
	eventIterator := firestoreClient.Collection("events").Doc("eventsName").Collection(categoryName).Documents(context.Background())

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
			return nil, nil
		}

		var event Event
		errInConversion := convertToEventObject(doc, &event)

		if errInConversion != nil {
			return nil, nil
		}
		events = append(events, event)
	}

	return &events, nil
}

func GetEventsName(categoryName string, eventNames []string) (*[]Event, error) {
	var events []Event

	for _, eventName := range eventNames {
		event, err := GetEventName(categoryName, eventName)

		if err != nil {
			return nil, err
		}

		events = append(events, *event)
	}
	return &events, nil
}

func GetAllEventsForAllCategory() (*[]Category, error) {
	allCategory, err := category.GetAllCategory()

	if err != nil {
		return nil, err
	}

	var categories []Category
	for _, value := range *allCategory {
		category, err := GetEventsForCategory(value.Name)
		if err != nil {
			return nil, err
		}
		categories = append(categories, *category)
	}
	return &categories, nil
}
