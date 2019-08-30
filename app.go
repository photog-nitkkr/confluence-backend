package main

import (
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func main() {
	http.HandleFunc("/", Hello)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}