package models

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

// Hero type definition
type Hero struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Title     string    `json:"title"`
	Faction   string    `json:"faction"`
	Race      string    `json:"race"`
	Location  string    `json:"location"`
	CreatedAt time.Time `json:"created" db:"created_at"`
	UpdatedAt time.Time `json:"updated" db:"updated_at"`
}

// GetHeroes method that gets all the heroes
func GetHeroes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	rows, err := db.Query(
		context.Background(),
		`
		select id, name, title, faction, race, location, updated_at, created_at 
		from heroes
		`,
	)
	if err != nil {
		fmt.Println(err)
	}
	var heroes []Hero
	for rows.Next() {
		var hero Hero
		err := rows.Scan(
			&hero.ID,
			&hero.Name,
			&hero.Title,
			&hero.Faction,
			&hero.Race,
			&hero.Location,
			&hero.UpdatedAt,
			&hero.CreatedAt,
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to list heroes: %v\n", err)
			return
		}
		heroes = append(heroes, hero)
	}

	encodeResponseAsJSON(heroes, w)
}

// GetHero method that gets a hero that matches an id
func GetHero(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var hero Hero
	row := db.QueryRow(
		context.Background(),
		`
		select id, name, title, faction, race, location, updated_at, created_at 
		from heroes 
		where id=$1
		`,
		params["id"],
	)
	err := row.Scan(
		&hero.ID,
		&hero.Name,
		&hero.Title,
		&hero.Faction,
		&hero.Race,
		&hero.Location,
		&hero.UpdatedAt,
		&hero.CreatedAt,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to retrieve hero: %v\n", err)
		return
	}

	encodeResponseAsJSON(hero, w)
}

// CreateHero method that creates a new hero
func CreateHero(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var hero Hero
	_ = json.NewDecoder(r.Body).Decode(&hero)
	t := time.Now()
	hero.CreatedAt = t
	hero.UpdatedAt = t
	_, err := db.Exec(
		context.Background(),
		`
		insert into heroes(name, title, faction, race, location, updated_at, created_at) 
		values($1, $2, $3, $4, $5, $6, $7)
		`,
		hero.Name,
		hero.Title,
		hero.Faction,
		hero.Race,
		hero.Location,
		hero.UpdatedAt,
		hero.CreatedAt,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to add hero: %v\n", err)
		return
	}

	encodeResponseAsJSON(hero, w)
}

// UpdateHero method that updates a hero
func UpdateHero(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var hero Hero
	_ = json.NewDecoder(r.Body).Decode(&hero)
	hero.UpdatedAt = time.Now()
	_, err := db.Exec(
		context.Background(),
		`
		update heroes set name=$1, title=$2, faction=$3, race=$4, location=$5, updated_at=$6
		where id=$7
		`,
		hero.Name,
		hero.Title,
		hero.Faction,
		hero.Race,
		hero.Location,
		hero.UpdatedAt,
		params["id"],
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to update hero: %v\n", err)
		return
	}

	encodeResponseAsJSON(hero, w)
}

// DeleteHero method that deletes a hero
func DeleteHero(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	_, err := db.Exec(context.Background(), "delete from heroes where id=$1", params["id"])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to delete hero: %v\n", err)
		return
	}

	encodeResponseAsJSON(map[string]string{"success": "true"}, w)
}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
