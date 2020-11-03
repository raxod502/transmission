package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/raxod502/transmission/backend/src/api"
)

func serveHtml(name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./frontend/html/"+name+".html")
	})
}

func Start(addr string) error {
	r := mux.NewRouter()

	css := http.FileServer(http.Dir("./frontend/css"))
	js := http.FileServer(http.Dir("./frontend/js/out"))

	r.Handle("/", serveHtml("index"))
	r.PathPrefix("/css/").Handler(http.StripPrefix("/css/", css))
	r.PathPrefix("/js/").Handler(http.StripPrefix("/js/", js))
	api.Setup(r.PathPrefix("/api/v1/").Subrouter())

	http.Handle("/", r)

	return http.ListenAndServe(addr, nil)
}
