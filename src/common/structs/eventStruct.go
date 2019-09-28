package structs

type Event struct {
	Name string `json:"name"`
	Category string `json:"category"`
	Rules []string `json:"rules"`
	Prize string `json:"prize"`
	Venue string `json:"venue"`
	Coordinators []string `json:"coordinators"`
	Description string `json:"description"`
}

