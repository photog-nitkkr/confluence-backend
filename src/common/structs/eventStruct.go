package structs

type Event struct {
	Name string `json:"name,omitempty"`
	Category string `json:"category,omitempty"`
	Rules []string `json:"rules,omitempty"`
	Prize string `json:"prize,omitempty"`
	Venue string `json:"venue,omitempty"`
	Coordinators []string `json:"coordinators,omitempty"`
	Description string `json:"description,omitempty"`
	ImageUrl string `json:"imageUrl,omitempty"`
}

