package main

import (
	"fmt"
	"html/template"
	"net/http"
	"runtime/debug"
)

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}
func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}
func (app *application) render(w http.ResponseWriter, page string, data *templateData) {
	  funcMap := template.FuncMap{
        "humanDate": humanDate,
    }
	files := []string{
		"./ui/html/base.layout.html",
		"./ui/html/" + page,
	}
	ts, err := template.New("").Funcs(funcMap).ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		app.serverError(w, err)
	}
}
