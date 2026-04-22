package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/logs", app.logList)
	mux.HandleFunc("/resources", app.resourceList)
	mux.HandleFunc("/log", app.logView)
	fs := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fs))
	return mux
}
