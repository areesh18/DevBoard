package main

import (
	"html/template"
	"path/filepath"
	"time"

	"github.com/areesh18/devboard/internals/models"
)

type templateData struct {
	Log       *models.Log
	Logs      []*models.Log
	Resource  *models.Resource
	Resources []*models.Resource
}

func humanDate(t time.Time) string {
	ist := t.UTC().Add(5*time.Hour + 30*time.Minute)
	return ist.Format("02 Jan 2006 at 15:04")
}
func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err:=filepath.Glob("./ui/html/*.page.html")
	if err != nil {
		return nil, err
	}
	for _,page:=range pages{
		name:=filepath.Base(page)
		funcMap:=template.FuncMap{
			"humanDate":humanDate,
		}
		ts, err:=template.New(name).Funcs(funcMap).ParseGlob("./ui/html/*.layout.html")
		if err != nil {
			return nil, err
		}
		ts, err=ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}
		cache[name]=ts
	}
	return cache, nil
}
