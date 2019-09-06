package category

import (
	"../../common/structs"
	"../../db"
	"cloud.google.com/go/firestore"
	"context"
	"errors"
)

func GetCategory(categoryName string) (*structs.Category , error) {
	firestoreClient := db.GetFirestore()

	doc, err := firestoreClient.Collection("categories").Doc(categoryName).Get(context.Background())

	if err != nil {
		return nil, err
	}

	if doc.Data() == nil {
		return nil, errors.New("Internal Server Error / No Category")
	}

	var category structs.Category

	errInCustomObject := convertToCategoryObject(doc, &category)

	if errInCustomObject != nil {
		return nil, errInCustomObject
	}

	return &category, nil
}

func convertToCategoryObject(firestoreDocument *firestore.DocumentSnapshot, category *structs.Category) error {
	err := firestoreDocument.DataTo(category)

	if err != nil {
		return err
	}
	return nil
}