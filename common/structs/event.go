package structs

type Event struct {
	Name string `json:"name"`
	DisplayName string `json:"displayName"`
	Category string `json:"category"`
	Coordinators []Person `json:"coordinators"`
	Description string `json:description`
}
