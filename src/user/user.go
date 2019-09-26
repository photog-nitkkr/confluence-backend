package user

import (
	"common/structs"
	"db"
	"fmt"
	"golang.org/x/net/context"
)

func HandleFirestoreUser(tokenInfo *structs.TokenInfo) interface{} {
	firestoreClient := db.GetFirestore()

	user, _ := firestoreClient.Collection("users").Doc(tokenInfo.Sub).Get(context.Background())

	if user != nil {
		fmt.Println("user found")
		return nil
	}
	fmt.Println("user found")
	return nil
}



