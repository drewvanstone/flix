package main

import (
	"log"

	"github.com/drewvanstone/flix/http"
	"github.com/drewvanstone/flix/postgres"
)

func main() {
	// Set up all dependencies
	db := postgres.NewDB()
	handle := http.NewFlixHandler(db)

	// Set up router
	mux := http.NewServeMux()
	mux.HandleFunc("/movie", handle.Movie)
	mux.HandleFunc("/movies", handle.Movies)

	// Serve routes
	log.Fatal(http.ListenAndServe(":8080", mux))
}
