package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v4"
)

var conn *pgx.Conn

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	db := connectToDataBase()
	defer db.Close(context.Background())

	if err := http.ListenAndServe(":"+port, NewServer()); err != nil {
		log.Fatalln(err)
	}
}
