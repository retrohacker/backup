package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

const (
	usage = `
	USAGE: backup PATH

	PATH - path to the directory you would like to sync with the cloud`
)

func argError(msg string) {
	fmt.Fprintf(
		os.Stderr,
		"%v\n%v\n",
		usage,
		msg,
	)
}

func main() {
	flag.Parse()
	flags := flag.Args()
	if len(flags) != 1 {
		argError(fmt.Sprintf("Expected 1 argument, instead saw %v", len(flags)))
		os.Exit(1)
	}

	dir, e := filepath.Abs(flags[0])
	if e != nil {
		argError(fmt.Sprintf("%v", e))
	}

	fmt.Println("Indexing ", dir, "...")
	fc := make(chan *File, 100)
	fh := make(chan *HashFile, 100)
	go List(dir, fc)
	go Hash(fc, fh)
	for file := range fh {
		fmt.Println(file.path, " ", file.hash)
	}
	/*
		r := mux.NewRouter()
		registerStaticRoutes(r)
		http.Handle("/", r)
		log.Fatal(http.ListenAndServe(":8989", nil))
	*/
}
