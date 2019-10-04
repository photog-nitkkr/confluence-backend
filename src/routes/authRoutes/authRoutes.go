package authRoutes

import (
	"github.com/gorilla/mux"
)

func AuthRoutesHandler() *mux.Router {
	muxRouter := mux.NewRouter()

	addRoutes(muxRouter)

	return muxRouter
}

func addRoutes(muxRouter *mux.Router) {
	muxRouter.HandleFunc("/login/", login).Methods("POST")
	muxRouter.HandleFunc("/event/", eventRegister).Methods("PUT")
	muxRouter.HandleFunc("/onBoard/", userSignUp).Methods("PUT")
	muxRouter.HandleFunc("/event/", getRegisteredEvents).Methods("GET")
}

