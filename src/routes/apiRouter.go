package routes

import (
	"net/http"
	"routes/categoryRoutes"
	"routes/developerRoutes"
	"routes/eventDescRoutes"
	"routes/eventNameRoutes"

	"github.com/gorilla/mux"
)

func ApiMuxRouterInitializer() *mux.Router {
	muxRouter := mux.NewRouter()
	muxRouter.PathPrefix("/events/name").Handler(http.StripPrefix("/events/name", eventNameRoutes.EventHandler()))
	muxRouter.PathPrefix("/events/desc").Handler(http.StripPrefix("/events/desc", eventDescRoutes.EventHandler()))
	muxRouter.PathPrefix("/category").Handler(http.StripPrefix("/category", categoryRoutes.CategoryHandler()))
	muxRouter.PathPrefix("/developers").Handler(http.StripPrefix("/developers", developerRoutes.DeveloperHandler()))
	return muxRouter
}
