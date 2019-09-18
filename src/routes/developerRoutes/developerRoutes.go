package developerRoutes

import (
	"github.com/gorilla/mux"
)

func DeveloperHandler() *mux.Router {
	muxRouter := mux.NewRouter()

	addRoutes(muxRouter)

	return muxRouter
}

func addRoutes(muxRouter *mux.Router) {
	muxRouter.HandleFunc("/", writeDeveloper).Methods("POST")
	muxRouter.HandleFunc("/", readDeveloper).Methods("GET")
}
