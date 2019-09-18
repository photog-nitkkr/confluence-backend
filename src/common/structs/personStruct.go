package structs

import (
	"cloud.google.com/go/firestore"
)

type Person struct {
	Name       string `json:"name"`
	ProfilePic string `json:"profilePic"`
	FbID       string `json:"fbID"`
	GithubID   string `json:"githubID"`
	Email      string `json:"email"`
	MobileNo   string `json:"mobileNo"`
}

func ConvertToPersonObject(firestoreDocument *firestore.DocumentSnapshot, person *Person) error {
	err := firestoreDocument.DataTo(person)

	if err != nil {
		return err
	}
	return nil
}
