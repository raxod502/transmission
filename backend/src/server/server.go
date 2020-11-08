package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/raxod502/transmission/backend/src/api"
)

func serveHtml(name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./frontend/static/html/"+name+".html")
	})
}

func Start(addr string) error {
	r := mux.NewRouter()

	static := http.FileServer(http.Dir("./frontend/out"))

	r.Handle("/", serveHtml("index"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", static))
	api.Setup(r.PathPrefix("/api/v1/").Subrouter())

	http.Handle("/", r)

	return http.ListenAndServe(addr, nil)
}
