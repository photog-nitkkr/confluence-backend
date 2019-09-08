package category

import (
	"../db"
	"cloud.google.com/go/firestore"
	"context"
	"errors"
)

func GetCategory(categoryName string) (*Category , error) {
	firestoreClient := db.GetFirestore()

	doc, err := firestoreClient.Collection("categories").Doc(categoryName).Get(context.Background())

	if err != nil {
		return nil, err
	}

	if doc.Data() == nil {
		return nil, errors.New("Internal Server Error / No Category")
	}

	var category Category

	errInCustomObject := convertToCategoryObject(doc, &category)

	if errInCustomObject != nil {
		return nil, errInCustomObject
	}

	return &category, nil
}

func convertToCategoryObject(firestoreDocument *firestore.DocumentSnapshot, category *Category) error {
	err := firestoreDocument.DataTo(category)

	if err != nil {
		return err
	}
	return nil
}