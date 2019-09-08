package categoryRoutes

import (
	"github.com/gorilla/mux"
)

func CategoryHandler() *mux.Router {
	muxRouter := mux.NewRouter()

	addRoutes(muxRouter)

	return muxRouter
}

func addRoutes(muxRouter *mux.Router) {
	muxRouter.HandleFunc("/", readCategory).Methods("GET")
}
