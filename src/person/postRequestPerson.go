package person

import (
	"context"

	. "db"

	. "common/structs"
)

func AddPersonInFireStore(person Person, team string, role string) error {
	firestoreClient := GetFirestore()

	_, err := firestoreClient.Collection("contacts").Doc(team).Collection(role).Doc(person.Email).Set(context.Background(), person)
	return err
}
