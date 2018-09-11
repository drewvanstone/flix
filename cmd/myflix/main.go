package main

import (
	"net/http"

	"github.com/drewvanstone/minerva/postgres"
	"github.com/drewvanstone/minerva/web"
)

// https://www.codementor.io/anshulsanghi/so-you-wanna-start-developing-web-apps-with-go-huh-handle-handler-handlefunc-eziu2go2t
func main() {
	db := postgres.NewDB()
	moviesHandler := web.NewHandler(db)

	mux := http.NewServeMux()
	mux.Handle("/movie", moviesHandler)

	http.ListenAndServe(":8100", mux)
}
