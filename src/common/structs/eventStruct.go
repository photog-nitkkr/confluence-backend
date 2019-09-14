package structs

type Event struct {
	Name string `json:"name"`
	DisplayName string `json:"displayName"`
	Category string `json:"category"`
	Coordinators []string `json:"coordinators"`
}

