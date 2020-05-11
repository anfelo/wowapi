package models

import (
	"github.com/jackc/pgx/v4"
)

var db *pgx.Conn

// SetDatabase method that initializes the database connection
func SetDatabase(database *pgx.Conn) {
	db = database
}
