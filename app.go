package main

import (
	"fmt"
	"log"
	"net/http"

	"./db"
	"./routes/categoryRoutes"
	"github.com/gorilla/mux"
)

func main() {
	client := db.GetFirestore()
	fmt.Println(client)

	startListening()
}

func muxRouterInitializer() *mux.Router {
	muxRouter := mux.NewRouter()
	//muxRouter.PathPrefix("/events").Handler(http.StripPrefix("/events", events.EventHandler()))
	muxRouter.PathPrefix("/category").Handler(http.StripPrefix("/category", categoryRoutes.CategoryHandler()))
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
