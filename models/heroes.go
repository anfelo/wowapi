package models

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/anfelo/wowapi/database"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
)

// Hero type definition
type Hero struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Title    string `json:"title"`
	Faction  string `json:"faction"`
	Race     string `json:"race"`
	Location string `json:"location"`
}

var conn *pgx.Conn = database.GetConnection()

// GetHeroes method that gets all the heroes
func GetHeroes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	rows, _ := conn.Query(context.Background(), "select * from heroes")

	var heroes []Hero
	for rows.Next() {
		err := rows.Scan(&heroes)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to list heroes: %v\n", err)
			return
		}
	}

	encodeResponseAsJSON(heroes, w)
}

// GetHero method that gets a hero that matches an id
func GetHero(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	var hero Hero
	row := conn.QueryRow(context.Background(), "select * from heroes where id=$1", params["id"])
	err := row.Scan(&hero)
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
	_, err := conn.Exec(context.Background(), "insert into heroes values($1)", hero)
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
	_, err := conn.Exec(
		context.Background(),
		`update heroes set name=$1, title=$2, faction=$3, race=$4, location=$5  where id=$6`,
		hero.Name,
		hero.Title,
		hero.Faction,
		hero.Race,
		hero.Location,
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
	_, err := conn.Exec(context.Background(), "delete from heroes where id=$1", params["id"])
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
