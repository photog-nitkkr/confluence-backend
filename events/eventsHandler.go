package events

import (
	"github.com/gorilla/mux"
	"./categories"
	"net/http"
	"./eventNames"
)

func EventHandler() *mux.Router {
	muxRouter := mux.NewRouter()

	muxRouter.PathPrefix("/categories").Handler(http.StripPrefix("/categories", categories.Handler()))
	muxRouter.PathPrefix("/names").Handler(http.StripPrefix("/names", eventNames.EventNamesRouteHandler()))

	return muxRouter
}
