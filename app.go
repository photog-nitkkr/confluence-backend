package main

import (
	"db"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"routes"
)

func main() {
	client := db.GetFirestore()
	fmt.Println(client)

	startListening()
}

func muxRouterInitializer() *mux.Router {
	muxRouter := mux.NewRouter()
	muxRouter.PathPrefix("/api").Handler(http.StripPrefix("/api", routes.ApiMuxRouterInitializer()))
	return muxRouter
}

func startListening() {
	muxRouter := muxRouterInitializer()
	err := http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(muxRouter))
	if err != nil {
		fmt.Println(err)
		return
	}

}

func MuxHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := fmt.Fprintf(w, "It's working with mux")

	if err != nil {
		log.Fatalln(err)
	}
}
