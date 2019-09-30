package structs

import (
	"cloud.google.com/go/firestore"
)

type Category struct {
	Name string `json:"name,omitempty"`
	Events []Event `json:"events,omitempty"`
}

func ConvertToCategoryObject(firestoreDocument *firestore.DocumentSnapshot, category *Category) error {
	err := firestoreDocument.DataTo(category)

	if err != nil {
		return err
	}
	return nil
}