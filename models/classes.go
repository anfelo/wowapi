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
	"github.com/jackc/pgtype"
)

// Class type definition
type Class struct {
	ID        int          `json:"id"`
	Name      string       `json:"name"`
	RolesJSON pgtype.JSONB `json:"-" db:"roles"`
	Roles     interface{}  `json:"roles" db:"-"`
	URL       string       `json:"url" db:"-"`
	CreatedAt time.Time    `json:"created" db:"created_at"`
	UpdatedAt time.Time    `json:"updated" db:"updated_at"`
}

// GetClasses method that gets all the Classes
func GetClasses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	rows, err := db.Query(
		context.Background(),
		`
		select id, name, roles, updated_at, created_at 
		from classes
		`,
	)
	if err != nil {
		fmt.Println(err)
	}
	classes := []Class{}
	for rows.Next() {
		var class Class
		err := rows.Scan(
			&class.ID,
			&class.Name,
			&class.RolesJSON,
			&class.UpdatedAt,
			&class.CreatedAt,
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to list classes: %v\n", err)
			return
		}
		class.URL = r.Host + r.URL.Path + "/" + strconv.Itoa(class.ID)
		class.Roles = class.RolesJSON.Get()
		classes = append(classes, class)
	}

	encodeResponseAsJSON(classes, w)
}

// GetClass method that gets a Class that matches an id
func GetClass(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var class Class
	row := db.QueryRow(
		context.Background(),
		`
		select id, name, roles, updated_at, created_at 
		from classes 
		where id=$1
		`,
		params["id"],
	)
	err := row.Scan(
		&class.ID,
		&class.Name,
		&class.RolesJSON,
		&class.UpdatedAt,
		&class.CreatedAt,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to retrieve class: %v\n", err)
		return
	}
	class.URL = r.Host + r.URL.Path
	class.Roles = class.RolesJSON.Get()

	encodeResponseAsJSON(class, w)
}

// CreateClass method that creates a new Class
func CreateClass(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var class Class
	_ = json.NewDecoder(r.Body).Decode(&class)
	t := time.Now()
	class.CreatedAt = t
	class.UpdatedAt = t
	_, err := db.Exec(
		context.Background(),
		`
		insert into classes(name, roles, updated_at, created_at) 
		values($1, $2, $3, $4)
		`,
		class.Name,
		class.Roles,
		class.UpdatedAt,
		class.CreatedAt,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to add class: %v\n", err)
		return
	}
	encodeResponseAsJSON(class, w)
}

// UpdateClass method that updates a Class
func UpdateClass(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var class Class
	_ = json.NewDecoder(r.Body).Decode(&class)
	class.UpdatedAt = time.Now()
	_, err := db.Exec(
		context.Background(),
		`
		update classes set name=$1, roles=$2, updated_at=$3
		where id=$4
		`,
		class.Name,
		class.Roles,
		class.UpdatedAt,
		params["id"],
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to update class: %v\n", err)
		return
	}
	class.URL = r.Host + r.URL.Path

	encodeResponseAsJSON(class, w)
}

// DeleteClass method that deletes a Class
func DeleteClass(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	commandTag, err := db.Exec(
		context.Background(),
		"delete from classes where id=$1",
		params["id"],
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to delete class: %v\n", err)
		return
	}
	if commandTag.RowsAffected() != 1 {
		fmt.Fprintf(os.Stderr, "No row found to delete for id: %v\n", params["id"])
		return
	}

	encodeResponseAsJSON(map[string]string{"success": "true"}, w)
}
