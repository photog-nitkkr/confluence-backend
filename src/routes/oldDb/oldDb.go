package oldDb

import (
	"github.com/gorilla/mux"
	. "net/http"
	"oldDb"
)

func OldDbHandler() *mux.Router {
	muxRouter := mux.NewRouter()

	addRoutes(muxRouter)

	return muxRouter
}

func addRoutes(muxRouter *mux.Router) {
	muxRouter.HandleFunc("/", addOldData).Methods("GET")
}

func addOldData(w ResponseWriter, r *Request) {
	category := r.URL.Query()["category"]
	for _, cat := range category {
		oldDb.GetOld(cat)
	}
}
