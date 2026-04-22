package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/areesh18/devboard/internals/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}
	app.render(w, "home.page.html", nil)

}
func (app *application) logList(w http.ResponseWriter, r *http.Request) {
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
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
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
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
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
func (app *application) logCreateForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, "create_log.page.html", nil)
}
func (app *application) logCreatePost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")
	tag := r.PostForm.Get("tag")

	errors := map[string]string{}
	if strings.TrimSpace(title) == "" {
		errors["title"] = "title cannot be empty"
	} else if len(title) > 100 {
		errors["title"] = "Title cannot be more than 100 characters"
	}

	if strings.TrimSpace(content) == "" {
		errors["content"] = "Content cannot be empty"
	}
	if len(tag) > 30 {
		errors["tag"] = "Tag cannot be more than 30 characters"
	}
	if len(errors) > 0 {
		fmt.Fprint(w, errors)
		return
	}
	id, err := app.logs.Insert(title, content, tag)
	if err != nil {
		app.serverError(w, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/log/%d", id), http.StatusSeeOther)
}
