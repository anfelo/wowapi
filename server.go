package main

import (
	"context"
	"fmt"
	"os"

	"github.com/anfelo/wowapi/middleware"
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
	r.Use(middleware.AuthMiddleware)
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
