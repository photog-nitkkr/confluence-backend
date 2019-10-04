package structs

import (
	"cloud.google.com/go/firestore"
)

type Sponsor struct {
	Name     string `json:"name,omitempty"`
	ImageURL string `json:"imageURL,omitempty"`
	TargetUrl string `json:"targetUrl,omitempty"`
	TagLine string `json:"tagLine,omitempty"`
}

func ConvertToSponsorObject(firestoreDocument *firestore.DocumentSnapshot, sponsor *Sponsor) error {
	err := firestoreDocument.DataTo(sponsor)

	if err != nil {
		return err
	}
	return nil
}
