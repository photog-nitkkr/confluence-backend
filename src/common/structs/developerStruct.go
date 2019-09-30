package structs

import (
	"cloud.google.com/go/firestore"
)

type Developer struct {
	Name     string `json:"name,omitempty"`
	ProfilePic string `json:"profilePic,omitempty"`
	FbID     string `json:"fbID,omitempty"`
	GithubID string `json:"githubID,omitempty"`
	Email    string `json:"email,omitempty"`
	MobileNo string `json:"mobileNo,omitempty"`
}

func ConvertToDeveloperObject(firestoreDocument *firestore.DocumentSnapshot, developer *Developer) error {
	err := firestoreDocument.DataTo(developer)

	if err != nil {
		return err
	}
	return nil
}
