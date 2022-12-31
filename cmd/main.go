package main

import (
	"SkillboxDiploma/pkg/result"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
)

var localhost = "localhost:8282"

func main() {
	r := chi.NewRouter()
	r.Get("/", handleConnection)
	http.ListenAndServe(localhost, r)
}

func handleConnection(w http.ResponseWriter, r *http.Request) {
	result := result.GetResultData()

	content, err := json.Marshal(result)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(content)
}
