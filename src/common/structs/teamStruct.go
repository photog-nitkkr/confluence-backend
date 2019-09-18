package structs

import (
	"cloud.google.com/go/firestore"
)

type Team struct {
	Name    string   `json:"name"`
	Members []Person `json:"members"`
}

func ConvertToTeamObject(firestoreDocument *firestore.DocumentSnapshot, team *Team) error {
	err := firestoreDocument.DataTo(team)

	if err != nil {
		return err
	}
	return nil
}
