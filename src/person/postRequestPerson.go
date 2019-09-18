package person

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	. "db"

	. "common/structs"
)

func AddPersonInFireStore(person Person, team string, role string) error {
	firestoreClient := GetFirestore()

	_, err := firestoreClient.Collection("contacts").Doc(team).Collection(role).Doc(person.Email).Set(context.Background(), person)
	return err
}

func AddPerson(r *http.Request, team string, role string) error {
	var personVar Person
	err := getPersonObject(r, &personVar)
	if err != nil {
		return err
	}

	errPersonInvalid := isPersonValid(personVar)
	if errPersonInvalid != nil {
		return errPersonInvalid
	}

	errInAdding := AddPersonInFireStore(personVar, team, role)
	return errInAdding
}

func isPersonValid(person Person) error {
	if person.Name == "" {
		err := errors.New("Error in adding person: Person's name can not be undefined")
		return err
	} else if person.Email == "" {
		err := errors.New("Error in adding person: Person's email ID can not be undefined")
		return err
	}
	return nil
}

func getPersonObject(r *http.Request, person *Person) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&person)

	if err != nil {
		return err
	}
	return nil
}
