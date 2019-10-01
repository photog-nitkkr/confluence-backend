package category

import (
	"context"
	"strings"

	"common/structs"
	. "db"
)

func AddCategoryInFireStore(category structs.Category) error {
	firestoreClient := GetFirestore()

	_, err := firestoreClient.Collection("categories").Doc(strings.ToLower(category.Name)).Set(context.Background(), category)
	return err
}
