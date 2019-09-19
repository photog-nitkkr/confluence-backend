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
	muxRouter.HandleFunc("/login", login).Methods("POST")
	//muxRouter.HandleFunc("/signUp", readCategory).Methods("POST")
}

