package server

import (
	"net/http"

	"github.com/gorilla/mux"
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

	http.Handle("/", r)

	return http.ListenAndServe(addr, nil)
}
