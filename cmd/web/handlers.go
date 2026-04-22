package main

import (
	"net/http"
	"strconv"

	"github.com/areesh18/devboard/internals/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	app.render(w, "home.page.html", nil)

}
func (app *application) logList(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	logs, err := app.logs.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}
	data := &templateData{
		Logs: logs,
	}
	app.render(w, "logs.page.html", data)

}

func (app *application) resourceList(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	resources, err := app.resources.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}
	data := &templateData{
		Resources: resources,
	}
	app.render(w, "resources.page.html", data)

}
func (app *application) logView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	log, err := app.logs.Get(id)
	if err == models.ErrNoRecord {
		app.notFound(w)
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}
	data := &templateData{
		Log: log,
	}
	app.render(w, "log.page.html", data)

}
func (app *application) resourceView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	resource, err := app.resources.Get(id)
	if err == models.ErrNoRecord {
		app.notFound(w)
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}
	data := &templateData{
		Resource: resource,
	}
	app.render(w, "resource.page.html", data)
}
