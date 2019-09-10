package category

import (
	. "../common/structs"
	"../db"
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

	errInCustomObject := ConvertToCategoryObject(doc, &category)

	if errInCustomObject != nil {
		return nil, errInCustomObject
	}

	return &category, nil
}