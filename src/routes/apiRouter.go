package routes

import (
	"github.com/gorilla/mux"
	"net/http"
	"routes/categoryRoutes"
	"routes/eventDescRoutes"
	"routes/eventNameRoutes"
)

func ApiMuxRouterInitializer() *mux.Router {
	muxRouter := mux.NewRouter()
	muxRouter.PathPrefix("/events/name").Handler(http.StripPrefix("/events/name", eventNameRoutes.EventHandler()))
	muxRouter.PathPrefix("/events/desc").Handler(http.StripPrefix("/events/desc", eventDescRoutes.EventHandler()))
	muxRouter.PathPrefix("/category").Handler(http.StripPrefix("/category", categoryRoutes.CategoryHandler()))
	return muxRouter
}
