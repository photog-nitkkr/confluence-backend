package category

import (
	"context"

	"../common/structs"
	. "../db"
)

func AddCategoryInFireStore(category structs.Category) error {
	firestoreClient := GetFirestore()

	_, err := firestoreClient.Collection("categories").Doc(category.Name).Set(context.Background(), category)
	return err
}
