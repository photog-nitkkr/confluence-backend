package developer

import (
	"context"

	. "db"

	. "common/structs"
)

func AddDeveloperInFireStore(developer Developer) error {
	firestoreClient := GetFirestore()

	_, err := firestoreClient.Collection("developers").Doc(developer.Email).Set(context.Background(), developer)
	return err
}
