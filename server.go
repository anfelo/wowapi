package main

import "github.com/gorilla/mux"

// NewServer method that sets a new server
func NewServer() *mux.Router {
	r := mux.NewRouter()
	SetHeroesRoutes(r)
	return r
}
