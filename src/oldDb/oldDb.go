package oldDb

import (
	"category"
	"common/structs"
	"encoding/json"
	"event"
	"fmt"
	"io/ioutil"
	"net/http"
)

type OldEvent struct {
	Name string `json:"eventName,omitempty"`
	Category string `json:"category,omitempty"`
	Rules []string `json:"rules,omitempty"`
	Prize string `json:"prize,omitempty"`
	Venue string `json:"venue,omitempty"`
	Coordinators []string `json:"coordinators,omitempty"`
	Description string `json:"description,omitempty"`
	ImageUrl string `json:"imageUrl,omitempty"`
}



type oldDbData struct {
	Success bool
	Message string
	Data data
}

type data struct {
	Events []OldEvent
}


func GetOld(cat string) {
	resp, err := http.Get("https://us-central1-confluence19.cloudfunctions.net/api/events?category="+cat)
	if err != nil {
		fmt.Println(err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	var data oldDbData
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println(err)
		return
	}
	var categoryData structs.Category
	categoryData.Name = data.Data.Events[0].Category
	err = category.AddCategoryInFireStore(categoryData)
	if err != nil {
		fmt.Println(err)
	}
	for _,eventData := range data.Data.Events {
		newEvent := fromOldEventToNew(eventData)
		err := event.AddEventInFireStoreUtil(newEvent)
		fmt.Println(newEvent.Name)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func fromOldEventToNew(oldEvent OldEvent) structs.Event {
	var newEvent structs.Event
	newEvent = structs.Event(oldEvent)
	return newEvent
}
