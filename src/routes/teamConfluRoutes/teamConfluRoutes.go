package teamConfluRoutes

import (
	"github.com/gorilla/mux"
)

func TeamConfluHandler() *mux.Router {
	muxRouter := mux.NewRouter()

	addRoutes(muxRouter)

	return muxRouter
}

func addRoutes(muxRouter *mux.Router) {
	muxRouter.HandleFunc("/roles/", getTeamConfluRoles).Methods("GET")
	muxRouter.HandleFunc("/", writeTeamConflu).Methods("POST")
	muxRouter.HandleFunc("/", readTeamConflu).Methods("GET")
}
