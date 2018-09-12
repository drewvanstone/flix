package http

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *FlixHandler) Movies(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.getMovies(w, r)
	default:
		fmt.Fprintf(w, "Something went wrong")
		return
	}
}

func (h *FlixHandler) getMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := h.store.GetMovies()
	if err != nil {
		fmt.Fprintf(w, "Error getting movie. Got error %v\n", err)
	}

	json, err := json.Marshal(movies)
	if err != nil {
		fmt.Fprintf(w, "Something went wrong. Got error %v\n", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
