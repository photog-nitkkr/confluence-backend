package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	firestore "cloud.google.com/go/firestore"
	"github.com/gorilla/mux"
)

func main() {
	client := firestoreInitialization()
	fmt.Println(client)

	http.Handle("/", muxRouterInitialization())
	http.ListenAndServe(":8080", nil)
}

func muxRouterInitialization() *mux.Router {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", MuxHandler)
	return muxRouter
}

func firestoreInitialization() *firestore.Client {
	ctx := context.Background()
	projectId := "confluence-backend"
	firebaseError := "Failed to create firebase client with err:%v"

	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalln(firebaseError, err)
	}
	return client
}

func MuxHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "It's working with mux")
}
