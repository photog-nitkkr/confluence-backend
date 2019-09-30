package structs

import (
	"cloud.google.com/go/firestore"
)

type Person struct {
	Name       string `json:"name,omitempty"`
	ProfilePic string `json:"profilePic,omitempty"`
	FbID       string `json:"fbID,omitempty"`
	GithubID   string `json:"githubID,omitempty"`
	Email      string `json:"email,omitempty"`
	MobileNo   string `json:"mobileNo,omitempty"`
}

func ConvertToPersonObject(firestoreDocument *firestore.DocumentSnapshot, person *Person) error {
	err := firestoreDocument.DataTo(person)

	if err != nil {
		return err
	}
	return nil
}
