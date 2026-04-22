package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic,app.logRequest, secureHeaders)

	mux := pat.New()
	mux.Get("/", http.HandlerFunc(app.home))
	mux.Get("/logs", http.HandlerFunc(app.logList))
	mux.Get("/resources", http.HandlerFunc(app.resourceList))
	mux.Get("/log/create",http.HandlerFunc(app.logCreateForm))
	mux.Get("/log/:id", http.HandlerFunc(app.logView))
	mux.Get("/resource/:id", http.HandlerFunc(app.resourceView))
	fs := http.FileServer(http.Dir("./ui/static"))
	mux.Get("/static/", http.StripPrefix("/static", fs))
	return standardMiddleware.Then(mux)
}
