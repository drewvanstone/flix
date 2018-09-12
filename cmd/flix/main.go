package main

import (
	"github.com/drewvanstone/flix/http"
	"github.com/drewvanstone/flix/postgres"
)

func main() {
	db := postgres.NewDB()
	handle := http.NewFlixHandler(db)

	mux := http.NewServeMux()
	mux.HandleFunc("/movie", handle.Movie)
	mux.HandleFunc("/movies", handle.Movies)

	http.ListenAndServe(":8100", mux)
}
