package events

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Handler() *mux.Router {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", Abc)
	return muxRouter
}

func Abc(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	fmt.Println(r.RequestURI)
	w.WriteHeader(http.StatusOK)
	_, err := fmt.Fprintf(w, "It's working with mux in different file")

	if err != nil {
		log.Fatalln(err)
	}
}
