package category

import (
	"../db"
	"errors"
	"golang.org/x/net/context"
	"google.golang.org/api/iterator"
)

func GetAllCategory() (*[]Category , error) {
	var categories []Category

	fireStoreClient := db.GetFirestore()

	categoriesIterator := fireStoreClient.Collection("categories").Documents(context.Background())

	if categoriesIterator == nil {
		return nil, errors.New("Internal Server Error / No Categories")
	}

	for {
		doc, err := categoriesIterator.Next()

		if err == iterator.Done {
			return &categories, nil
		}

		if err != nil {
			return nil, err
		}

		var category Category
		errInConversion := convertToCategoryObject(doc, &category)

		if errInConversion != nil {
			return nil, errInConversion
		}
		categories = append(categories, category)
	}

	return &categories, nil
}


