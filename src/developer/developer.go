package developer

import (
	"context"

	. "common/structs"
	"errors"

	"google.golang.org/api/iterator"

	"db"
)

func GetAllDevelopers() (*[]Developer, error) {
	firestoreClient := db.GetFirestore()

	developerIterator := firestoreClient.Collection("developers").Documents(context.Background())

	if developerIterator == nil {
		return nil, errors.New("Internal Server Error / No Developers")
	}

	var developers []Developer

	for {
		doc, err := developerIterator.Next()

		if err == iterator.Done {
			return &developers, nil
		}

		if err != nil {
			return nil, err
		}

		var developer Developer
		errInConversion := ConvertToDeveloperObject(doc, &developer)

		if errInConversion != nil {
			return nil, errInConversion
		}
		developers = append(developers, developer)
	}
}
