package main

import (
	"fmt"
	"log"
	"net/http"

	"./db"
	"./events"
	"github.com/gorilla/mux"
)

func main() {
	client := db.GetFirestore()
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

func MuxHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := fmt.Fprintf(w, "It's working with mux")

	if err != nil {
		log.Fatalln(err)
	}
}
