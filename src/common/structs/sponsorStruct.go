package structs

import (
	"cloud.google.com/go/firestore"
)

type Sponsor struct {
	Name     string `json:"name,omitempty"`
	ImageURL string `json:"imageURL,omitempty"`
}

func ConvertToSponsorObject(firestoreDocument *firestore.DocumentSnapshot, sponsor *Sponsor) error {
	err := firestoreDocument.DataTo(sponsor)

	if err != nil {
		return err
	}
	return nil
}
