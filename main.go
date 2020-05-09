package main

import (
	"log"
	"net/http"
)

func main() {
	if err := http.ListenAndServe(":8080", NewServer()); err != nil {
		log.Fatalln(err)
	}
}
