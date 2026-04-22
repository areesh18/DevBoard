package main

import (
	"fmt"
	"html/template"
	"net/http"
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
	files := []string{
		"./ui/html/base.layout.html",
		"./ui/html/home.page.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, err)
		return
	}

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
	files := []string{
		"./ui/html/base.layout.html",
		"./ui/html/logs.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}
	data:=&templateData{
		Logs: logs,
	}
	err=ts.ExecuteTemplate(w,"base",data)
	if err != nil {
		app.serverError(w, err)
	}

}

func (app *application) resourceList(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	resources, err := app.resources.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}
	for _, resource := range resources {
		fmt.Fprintf(w, "%v\n", resource)
	}
}
