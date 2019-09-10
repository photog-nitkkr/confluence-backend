package category

import (
	"../event"
)

type Category struct {
	Name string `json:"name"`
	DisplayName string `json:"displayName"`
	Events []event.Event `json:"events"`
}

