package db

import (
	firestore "cloud.google.com/go/firestore"
	"context"
	"log"
	"sync"
)

type db struct {
	client *firestore.Client
}

var singletonFireStore *db
var once sync.Once

func GetFirestore() *firestore.Client {
	once.Do(func() {
		singletonFireStore = &db{firestoreInitialization()}
	})

	return singletonFireStore.client
}

func firestoreInitialization() *firestore.Client {
	ctx := context.Background()
	projectId := "confluence-backend"
	firebaseError := "Failed to create firebase client with err: "

	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalln(firebaseError, err)
	}
	return client
}
