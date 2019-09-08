package category

import (
	"../common/structs"
)

type Category struct {
	Name string `json:"name"`
	DisplayName string `json:"displayName"`
	Events []structs.Event `json:"events"`
}

