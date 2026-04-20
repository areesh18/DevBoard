package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/logs", logList)
	mux.HandleFunc("/resources", resourceList)

	log.Printf("Starting server on port %s", *addr)
	fs := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fs))
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
