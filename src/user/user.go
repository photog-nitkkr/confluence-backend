package user

import (
	"common/structs"
	"db"
	"golang.org/x/net/context"
)

func HandleFirestoreUser(tokenInfo *structs.TokenInfo) (*structs.TokenInfo, error) {
	firestoreClient := db.GetFirestore()

	userDocument, err := firestoreClient.Collection("users").Doc(tokenInfo.Sub).Get(context.Background())

	if userDocument != nil && err == nil {
		var user structs.TokenInfo
		structs.ConvertToUserObject(userDocument, &user)
		return &user, nil
	}
	tokenInfo.OnBoard = false
	err = addUser(tokenInfo)
	if err != nil {
		return nil, err
	}
	return tokenInfo, nil
}

func addUser(tokenInfo *structs.TokenInfo) error {
	firestoreClient := db.GetFirestore()

	_, err := firestoreClient.Collection("users").Doc(tokenInfo.Sub).Set(context.Background(), tokenInfo)

	return err
}



