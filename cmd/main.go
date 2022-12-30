package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

var localhost = "localhost:8282"

func main() {
	r := chi.NewRouter()
	r.HandleFunc("/{category}", handleConnection)
	http.ListenAndServe(localhost, r)
}

func handleConnection(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
