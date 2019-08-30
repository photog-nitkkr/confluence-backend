package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HeyHandler)
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func HeyHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "It's working with mux")
}
