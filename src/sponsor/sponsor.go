package sponsor

import (
	"context"

	. "common/structs"
	"errors"

	"google.golang.org/api/iterator"

	"db"
)

func GetSponsorsFromFirestore() (*[]Sponsor, error) {
	firestoreClient := db.GetFirestore()

	sponsorIterator := firestoreClient.Collection("sponsors").Documents(context.Background())

	if sponsorIterator == nil {
		return nil, errors.New("Internal Server Error / No Sponsors")
	}

	var sponsors []Sponsor

	for {
		doc, err := sponsorIterator.Next()

		if err == iterator.Done {
			return &sponsors, nil
		}

		if err != nil {
			return nil, err
		}

		var sponsor Sponsor
		errInConversion := ConvertToSponsorObject(doc, &sponsor)

		if errInConversion != nil {
			return nil, errInConversion
		}
		sponsors = append(sponsors, sponsor)
	}
}
