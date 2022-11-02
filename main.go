package main

import (
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	s := "welcome"
	w.Write([]byte(s))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":3000", mux)
}
