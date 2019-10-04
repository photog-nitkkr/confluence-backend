package user

import (
	. "common/structs"
	"context"
	"db"
)

func AddUserEvent(category string, event string, team string, tokenInfo TokenInfo) error {
	firestoreClient := db.GetFirestore()

	_, err := firestoreClient.Collection("events").Doc("eventParticipants").Collection(category).Doc(event).Collection(team).Doc(tokenInfo.Sub).Set(context.Background(), tokenInfo)
	if err != nil {
		return err
	}
	var dummy DummyStruct
	_, err = firestoreClient.Collection("userEvents").Doc(tokenInfo.Sub).Collection(category).Doc(event).Set(context.Background(), dummy)
	if err != nil {
		return err
	}
	return nil
}
