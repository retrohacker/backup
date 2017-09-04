package main

import (
	"github.com/retrohacker/mux"
	"io"
	"net/http"
	"os"
)

func serve(file string) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		file, err := os.Open(file)
		if err != nil {
			res.WriteHeader(http.StatusNotFound)
		} else {
			io.Copy(res, file)
		}
	}
}

func serveBundle(res http.ResponseWriter, req *http.Request) {
	bundle, err := os.Open("./web/bundle.js")
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
	} else {
		io.Copy(res, bundle)
	}
}

func registerStaticRoutes(r *mux.Router) {
	r.HandleFunc("/bundle.js", serve("./web/bundle.js")).Methods("GET")
	r.PathPrefix("/css").Handler(http.FileServer(http.Dir("./web"))).Methods("GET")
	r.PathPrefix("/img").Handler(http.FileServer(http.Dir("./web"))).Methods("GET")
	r.PathPrefix("/fonts").Handler(http.FileServer(http.Dir("./web"))).Methods("GET")
	r.PathPrefix("/").HandlerFunc(serve("./web/index.html")).Methods("GET")
}
