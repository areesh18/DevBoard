package main

import (
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
	return t.Format("02 Jan 2006 at 15:04")
}
