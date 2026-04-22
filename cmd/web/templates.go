package main

import "github.com/areesh18/devboard/internals/models"

type templateData struct {
	Log *models.Log
	Logs []*models.Log
	Resource *models.Resource
	Resources []*models.Resource
}