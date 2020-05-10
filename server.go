package main

import (
	"github.com/anfelo/wowapi/database"
	"github.com/gorilla/mux"
)

// NewServer method that sets a new server
func NewServer() *mux.Router {
	r := mux.NewRouter()
	SetHeroesRoutes(r)
	database.SetDatabase()
	return r
}
