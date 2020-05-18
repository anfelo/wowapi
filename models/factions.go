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

// Faction type definition
type Faction struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	URL       string    `json:"url" db:"-"`
	CreatedAt time.Time `json:"created" db:"created_at"`
	UpdatedAt time.Time `json:"updated" db:"updated_at"`
}

// GetFactions method that gets all the factions
func GetFactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	rows, err := db.Query(
		context.Background(),
		`
		select id, name, updated_at, created_at 
		from factions
		`,
	)
	if err != nil {
		fmt.Println(err)
	}
	factions := []Faction{}
	for rows.Next() {
		var faction Faction
		err := rows.Scan(
			&faction.ID,
			&faction.Name,
			&faction.UpdatedAt,
			&faction.CreatedAt,
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to list factions: %v\n", err)
			return
		}
		faction.URL = r.Host + r.URL.Path + "/" + strconv.Itoa(faction.ID)
		factions = append(factions, faction)
	}

	encodeResponseAsJSON(factions, w)
}

// GetFaction method that gets a Faction that matches an id
func GetFaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var faction Faction
	row := db.QueryRow(
		context.Background(),
		`
		select id, name, updated_at, created_at 
		from factions 
		where id=$1
		`,
		params["id"],
	)
	err := row.Scan(
		&faction.ID,
		&faction.Name,
		&faction.UpdatedAt,
		&faction.CreatedAt,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to retrieve faction: %v\n", err)
		return
	}
	faction.URL = r.Host + r.URL.Path

	encodeResponseAsJSON(faction, w)
}

// CreateFaction method that creates a new Faction
func CreateFaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var faction Faction
	_ = json.NewDecoder(r.Body).Decode(&faction)
	t := time.Now()
	faction.CreatedAt = t
	faction.UpdatedAt = t
	_, err := db.Exec(
		context.Background(),
		`
		insert into factions(name, updated_at, created_at) 
		values($1, $2, $3)
		`,
		faction.Name,
		faction.UpdatedAt,
		faction.CreatedAt,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to add faction: %v\n", err)
		return
	}
	encodeResponseAsJSON(faction, w)
}

// UpdateFaction method that updates a Faction
func UpdateFaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var faction Faction
	_ = json.NewDecoder(r.Body).Decode(&faction)
	faction.UpdatedAt = time.Now()
	_, err := db.Exec(
		context.Background(),
		`
		update factions set name=$1, updated_at=$2
		where id=$3
		`,
		faction.Name,
		faction.UpdatedAt,
		params["id"],
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to update faction: %v\n", err)
		return
	}
	faction.URL = r.Host + r.URL.Path

	encodeResponseAsJSON(faction, w)
}

// DeleteFaction method that deletes a Faction
func DeleteFaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	commandTag, err := db.Exec(
		context.Background(),
		"delete from factions where id=$1",
		params["id"],
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to delete faction: %v\n", err)
		return
	}
	if commandTag.RowsAffected() != 1 {
		fmt.Fprintf(os.Stderr, "No row found to delete for id: %v\n", params["id"])
		return
	}

	encodeResponseAsJSON(map[string]string{"success": "true"}, w)
}
