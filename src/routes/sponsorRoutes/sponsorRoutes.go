package sponsorRoutes

import (
	"github.com/gorilla/mux"
)

func SponsorsHandler() *mux.Router {
	muxRouter := mux.NewRouter()

	addRoutes(muxRouter)

	return muxRouter
}

func addRoutes(muxRouter *mux.Router) {
	muxRouter.HandleFunc("/", writeSponsor).Methods("POST")
	muxRouter.HandleFunc("/", readSponsor).Methods("GET")
}
