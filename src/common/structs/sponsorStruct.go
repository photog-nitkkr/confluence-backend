package structs

import (
	"cloud.google.com/go/firestore"
)

type Sponsor struct {
	Name     string `json:"name"`
	ImageURL string `json:"imageURL"`
}

func ConvertToSponsorObject(firestoreDocument *firestore.DocumentSnapshot, sponsor *Sponsor) error {
	err := firestoreDocument.DataTo(sponsor)

	if err != nil {
		return err
	}
	return nil
}
