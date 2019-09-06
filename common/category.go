package common

type Category struct {
	Name string `json:"name"`
	DisplayName string `json:"displayName"`
	Events []Event `json:"events"`
}

