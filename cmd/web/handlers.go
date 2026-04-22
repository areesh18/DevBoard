package main

import (
	"html/template"
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
	data := &templateData{
		Logs: logs,
	}
	err = ts.ExecuteTemplate(w, "base", data)
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
	resources, err := app.resources.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}
	data := &templateData{
		Resources: resources,
	}
	files := []string{
		"./ui/html/base.layout.html",
		"./ui/html/resources.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}
	err = ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		app.serverError(w, err)
	}

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
	files := []string{
		"./ui/html/base.layout.html",
		"./ui/html/log.page.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}
	data := &templateData{
		Log: log,
	}
	err = ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		app.serverError(w, err)
	}

}
func (app *application) resourceView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err!=nil || id<1{
		app.notFound(w)
		return
	}
	resource, err:=app.resources.Get(id)
	if err==models.ErrNoRecord{
		app.notFound(w)
		return
	}else if err!=nil{
		app.serverError(w, err)
		return
	}
	data:=&templateData{
		Resource: resource,
	}
	files := []string{
		"./ui/html/base.layout.html",
		"./ui/html/resource.page.html",
	}
	ts, err:=template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}
	err=ts.ExecuteTemplate(w,"base",data)
	if err != nil {
		app.serverError(w, err)
	}
}
