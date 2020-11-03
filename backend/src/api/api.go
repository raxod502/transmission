package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func launchMissiles(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Missiles launched\n"))
}

func Setup(r *mux.Router) {
	r.HandleFunc("/launch", launchMissiles)
}
