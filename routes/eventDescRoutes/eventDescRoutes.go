package eventDescRoutes

import (
	"github.com/gorilla/mux"
)

func EventHandler() *mux.Router {
	muxRouter := mux.NewRouter()

	addRoutes(muxRouter)

	return muxRouter
}

func addRoutes(muxRouter *mux.Router) {
	muxRouter.HandleFunc("/", writeEvent).Methods("POST")
}