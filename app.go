package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"./events"
	firestore "cloud.google.com/go/firestore"
	"github.com/gorilla/mux"
)

func main() {
	client := firestoreInitialization()
	fmt.Println(client)

	startListening()
}

func muxRouterInitializer() *mux.Router {
	muxRouter := mux.NewRouter()
	muxRouter.PathPrefix("/activities").Handler(http.StripPrefix("/activities", events.Handler()))
	muxRouter.HandleFunc("/", MuxHandler)
	return muxRouter
}

func startListening() {
	muxRouter := muxRouterInitializer()

	err := http.ListenAndServe(":8080", muxRouter)
	if err != nil {
		listeningError := "Error in Listening: "
		log.Fatalln(listeningError, err)
	}
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

func MuxHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := fmt.Fprintf(w, "It's working with mux")

	if err != nil {
		log.Fatalln(err)
	}
}
