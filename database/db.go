package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

var db *pgx.Conn

// SetDatabase method that initializes the database connection
func SetDatabase() *pgx.Conn {
	db, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connection to database: %v\n", err)
		os.Exit(1)
	}
	return db
}

// GetConnection method that gets an initialized connection
func GetConnection() *pgx.Conn {
	return db
}
