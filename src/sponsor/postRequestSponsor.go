package sponsor

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	. "db"

	. "common/structs"
)

func AddSponsorInFireStore(sponsor Sponsor) error {
	firestoreClient := GetFirestore()

	_, err := firestoreClient.Collection("sponsors").Doc(sponsor.Name).Set(context.Background(), sponsor)
	return err
}

func AddSponsor(r *http.Request) error {
	var sponsorVar Sponsor
	err := getSponsorObject(r, &sponsorVar)
	if err != nil {
		return err
	}

	errSponsorInvalid := isSponsorValid(sponsorVar)
	if errSponsorInvalid != nil {
		return errSponsorInvalid
	}

	errInAdding := AddSponsorInFireStore(sponsorVar)
	return errInAdding
}

func isSponsorValid(sponsor Sponsor) error {
	if sponsor.Name == "" {
		err := errors.New("Error in adding sponsor: Sponsor's name can not be undefined")
		return err
	}
	return nil
}

func getSponsorObject(r *http.Request, sponsor *Sponsor) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&sponsor)

	if err != nil {
		return err
	}
	return nil
}
