package structs

import (
	"cloud.google.com/go/firestore"
)

type Developer struct {
	Name     string `json:"name"`
	FbID     string `json:"fbID"`
	GithubID string `json:"githubID"`
	Email    string `json:"email"`
	MobileNo string `json:"mobileNo"`
}

func ConvertToDeveloperObject(firestoreDocument *firestore.DocumentSnapshot, developer *Developer) error {
	err := firestoreDocument.DataTo(developer)

	if err != nil {
		return err
	}
	return nil
}
