package models

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Race type definition
type Race struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	IsAllied  bool      `json:"allied" db:"is_allied"`
	URL       string    `json:"url" db:"-"`
	CreatedAt time.Time `json:"created" db:"created_at"`
	UpdatedAt time.Time `json:"updated" db:"updated_at"`
}

// GetRaces method that gets all the races
func GetRaces(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	rows, err := db.Query(
		context.Background(),
		`
		select id, name, is_allied, updated_at, created_at 
		from races
		`,
	)
	if err != nil {
		fmt.Println(err)
	}
	races := []Race{}
	for rows.Next() {
		var race Race
		err := rows.Scan(
			&race.ID,
			&race.Name,
			&race.IsAllied,
			&race.UpdatedAt,
			&race.CreatedAt,
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to list races: %v\n", err)
			return
		}
		race.URL = r.Host + r.URL.Path + "/" + strconv.Itoa(race.ID)
		races = append(races, race)
	}

	encodeResponseAsJSON(races, w)
}

// GetRace method that gets a Race that matches an id
func GetRace(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var race Race
	row := db.QueryRow(
		context.Background(),
		`
		select id, name, is_allied, updated_at, created_at 
		from races 
		where id=$1
		`,
		params["id"],
	)
	err := row.Scan(
		&race.ID,
		&race.Name,
		&race.IsAllied,
		&race.UpdatedAt,
		&race.CreatedAt,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to retrieve race: %v\n", err)
		return
	}
	race.URL = r.Host + r.URL.Path

	encodeResponseAsJSON(race, w)
}

// CreateRace method that creates a new Race
func CreateRace(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var race Race
	_ = json.NewDecoder(r.Body).Decode(&race)
	t := time.Now()
	race.CreatedAt = t
	race.UpdatedAt = t
	_, err := db.Exec(
		context.Background(),
		`
		insert into races(name, is_allied, updated_at, created_at) 
		values($1, $2, $3, $4)
		`,
		race.Name,
		race.IsAllied,
		race.UpdatedAt,
		race.CreatedAt,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to add race: %v\n", err)
		return
	}
	encodeResponseAsJSON(race, w)
}

// UpdateRace method that updates a Race
func UpdateRace(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var race Race
	_ = json.NewDecoder(r.Body).Decode(&race)
	race.UpdatedAt = time.Now()
	_, err := db.Exec(
		context.Background(),
		`
		update races set name=$1, is_allied=$2, updated_at=$3
		where id=$4
		`,
		race.Name,
		race.IsAllied,
		race.UpdatedAt,
		params["id"],
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to update race: %v\n", err)
		return
	}
	race.URL = r.Host + r.URL.Path

	encodeResponseAsJSON(race, w)
}

// DeleteRace method that deletes a Race
func DeleteRace(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	commandTag, err := db.Exec(
		context.Background(),
		"delete from races where id=$1",
		params["id"],
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to delete race: %v\n", err)
		return
	}
	if commandTag.RowsAffected() != 1 {
		fmt.Fprintf(os.Stderr, "No row found to delete for id: %v\n", params["id"])
		return
	}

	encodeResponseAsJSON(map[string]string{"success": "true"}, w)
}
