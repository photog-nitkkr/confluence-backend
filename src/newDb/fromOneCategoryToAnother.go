package newDb

import (
	"event"
	"fmt"
)
//handling of error will not be done properly
//just used for one time purpose
//and checked manually if everything is cool


func FromOneCategoryToAnother() {
	categoryFrom := "music"
	categoryTo := "Music And Dance"

	fromOneCategoryToAnotherUtil(categoryFrom, categoryTo)
}

func fromOneCategoryToAnotherUtil( categoryFrom string, categoryTo string) {
	events, err := event.GetAllEventsForCategory(categoryFrom, "eventDesc")
	if err != nil {
		fmt.Println(err)
	}
	for _, eventStruct := range *events {
		eventStruct.Category = categoryTo
		err := event.AddEventInFireStoreUtil(eventStruct)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(eventStruct.Name)
		}
	}
}
