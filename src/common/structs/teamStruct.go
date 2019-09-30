package structs

import (
	"cloud.google.com/go/firestore"
)

type Team struct {
	Name    string   `json:"name,omitempty"`
	Members []Person `json:"members,omitempty"`
}

func ConvertToTeamObject(firestoreDocument *firestore.DocumentSnapshot, team *Team) error {
	err := firestoreDocument.DataTo(team)

	if err != nil {
		return err
	}
	return nil
}
