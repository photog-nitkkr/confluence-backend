package events
//
//import (
//	"./categories"
//	"./eventNames"
//	"github.com/gorilla/mux"
//	"net/http"
//)
//
//func EventHandler() *mux.Router {
//	muxRouter := mux.NewRouter()
//
//	muxRouter.PathPrefix("/categories").Handler(http.StripPrefix("/categories", categories.Handler()))
//	muxRouter.PathPrefix("/names").Handler(http.StripPrefix("/names", eventNames.EventNamesRouteHandler()))
//
//	return muxRouter
//}
