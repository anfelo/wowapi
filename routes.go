package main

import (
	"github.com/anfelo/wowapi/models"
	"github.com/gorilla/mux"
)

// SetHeroesRoutes method for initializing heroes routes
func SetHeroesRoutes(r *mux.Router) {
	r.HandleFunc("/api/heroes", models.GetHeroes).Methods("GET")
	r.HandleFunc("/api/heroes/{id}", models.GetHero).Methods("GET")
	r.HandleFunc("/api/heroes", models.CreateHero).Methods("POST")
	r.HandleFunc("/api/heroes/{id}", models.UpdateHero).Methods("PUT")
	r.HandleFunc("/api/heroes/{id}", models.DeleteHero).Methods("DELETE")
}

// SetRacesRoutes method for initializing races routes
func SetRacesRoutes(r *mux.Router) {
	r.HandleFunc("/api/races", models.GetRaces).Methods("GET")
	r.HandleFunc("/api/races/{id}", models.GetRace).Methods("GET")
	r.HandleFunc("/api/races", models.CreateRace).Methods("POST")
	r.HandleFunc("/api/races/{id}", models.UpdateRace).Methods("PUT")
	r.HandleFunc("/api/races/{id}", models.DeleteRace).Methods("DELETE")
}
