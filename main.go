package main

import (
	_ "github.com/retrohacker/go-sqlite3"
	"github.com/retrohacker/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	registerStaticRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
