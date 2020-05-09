package models

import (
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Hero type definition
type Hero struct {
	ID    string `json:"id"`
	Isbn  string `json:"isbn"`
	Title string `json:"title"`
}

var (
	heroes []*Hero
)

// GetHeroes method that gets all the heroes
func GetHeroes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	encodeResponseAsJSON(heroes, w)
}

// GetHero method that gets a hero that matches an id
func GetHero(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, hero := range heroes {
		if hero.ID == params["id"] {
			encodeResponseAsJSON(hero, w)
			return
		}
	}
	encodeResponseAsJSON(&Hero{}, w)
}

// CreateHero method that creates a new hero
func CreateHero(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var hero Hero
	_ = json.NewDecoder(r.Body).Decode(&hero)
	hero.ID = strconv.Itoa(rand.Intn(10000000)) // Mock ID - not safe
	heroes = append(heroes, &hero)
	encodeResponseAsJSON(hero, w)
}

// UpdateHero method that updates a hero
func UpdateHero(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, hero := range heroes {
		if hero.ID == params["id"] {
			heroes = append(heroes[:i], heroes[i+1:]...)
			var hero Hero
			_ = json.NewDecoder(r.Body).Decode(&hero)
			hero.ID = params["id"]
			heroes = append(heroes, &hero)
			encodeResponseAsJSON(hero, w)
			break
		}
	}
}

// DeleteHero method that deletes a hero
func DeleteHero(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, hero := range heroes {
		if hero.ID == params["id"] {
			heroes = append(heroes[:i], heroes[i+1:]...)
			break
		}
	}
	encodeResponseAsJSON(heroes, w)
}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
