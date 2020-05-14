package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/anfelo/wowapi/models"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
)

// NewServer method that sets a new server
func NewServer() *mux.Router {
	r := mux.NewRouter()
	SetHeroesRoutes(r)
	SetRacesRoutes(r)
	SetFactionsRoutes(r)
	SetClassesRoutes(r)
	return r
}

func connectToDataBase() *pgx.Conn {
	db, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connection to database: %v\n", err)
		os.Exit(1)
	}
	models.SetDatabase(db)
	return db
}

// ErrHandler method that handles http errors
func ErrHandler(err error) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, err.Error())
	})
}
