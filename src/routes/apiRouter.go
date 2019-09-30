package routes

import (
	"net/http"
	"routes/authRoutes"
	"routes/categoryRoutes"
	"routes/developerRoutes"
	"routes/eventDescRoutes"
	"routes/eventNameRoutes"
	"routes/teamConfluRoutes"
	"routes/sponsorRoutes"

	"github.com/gorilla/mux"
)

func ApiMuxRouterInitializer() *mux.Router {
	muxRouter := mux.NewRouter()
	muxRouter.PathPrefix("/events/name").Handler(http.StripPrefix("/events/name", eventNameRoutes.EventHandler()))
	muxRouter.PathPrefix("/events/desc").Handler(http.StripPrefix("/events/desc", eventDescRoutes.EventHandler()))
	muxRouter.PathPrefix("/category").Handler(http.StripPrefix("/category", categoryRoutes.CategoryHandler()))
	muxRouter.PathPrefix("/developers").Handler(http.StripPrefix("/developers", developerRoutes.DeveloperHandler()))
	muxRouter.PathPrefix("/teamConfluence").Handler(http.StripPrefix("/teamConflu", teamConfluRoutes.TeamConfluHandler()))
	muxRouter.PathPrefix("/auth").Handler(http.StripPrefix("/auth", authRoutes.AuthRoutesHandler()))
	muxRouter.PathPrefix("/sponsors").Handler(http.StripPrefix("/sponsors", sponsorRoutes.SponsorsHandler()))
	return muxRouter
}
