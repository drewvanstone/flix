package web

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/drewvanstone/flix"
)

type Handler struct {
	db flix.StorageService
}

func NewHandler(db flix.StorageService) *Handler {
	return &Handler{db}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.getMovie(w, r)
	case "POST":
		h.addMovie(w, r)
	case "DELETE":
		h.deleteMovie(w, r)
	default:
		fmt.Fprintf(w, "Something went wrong")
		return
	}
}

func (h *Handler) addMovie(w http.ResponseWriter, r *http.Request) {
	movie := r.URL.Query().Get("movie")
	if err := h.db.AddMovie(movie); err != nil {
		fmt.Fprintf(w, "Error creating movie %s. Got error: %v\n", movie, err)
		return
	}
	fmt.Fprintf(w, "Created movie %s\n", movie)
}

func (h *Handler) getMovie(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		fmt.Fprintf(w, "Something went wrong")
		return
	}

	movie, err := h.db.GetMovie(id)
	if err != nil {
		fmt.Fprintf(w, "Error reading movie %v\n", err)
		return
	}
	fmt.Fprintf(w, "Read movie %v\n", *movie)
}

func (h *Handler) deleteMovie(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		fmt.Fprintf(w, "Something went wrong")
		return
	}

	if err := h.db.DeleteMovie(id); err != nil {
		fmt.Fprintf(w, "Error deleting movie %d\n", id)
		return
	}
	fmt.Fprintf(w, "Deleted movie %d\n", id)
}
