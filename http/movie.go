package http

import (
	"fmt"
	"net/http"
	"strconv"
)

func (h *FlixHandler) Movie(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.getMovie(w, r)
	case "PUT":
		h.addMovie(w, r)
	case "DELETE":
		h.deleteMovie(w, r)
	default:
		fmt.Fprintf(w, "Something went wrong")
		return
	}
}

func (h *FlixHandler) addMovie(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	if err := h.store.AddMovie(title); err != nil {
		fmt.Fprintf(w, "Error creating movie %h. Got error: %v\n", title, err)
		return
	}
	fmt.Fprintf(w, "Created movie %s\n", title)
}

func (h *FlixHandler) getMovie(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		fmt.Fprintf(w, "Something went wrong")
		return
	}

	movie, err := h.store.GetMovie(id)
	if err != nil {
		fmt.Fprintf(w, "Error reading movie %v\n", err)
		return
	}
	fmt.Fprintf(w, "Read movie %v\n", *movie)
}

func (h *FlixHandler) deleteMovie(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		fmt.Fprintf(w, "Something went wrong")
		return
	}

	if err := h.store.DeleteMovie(id); err != nil {
		fmt.Fprintf(w, "Error deleting movie %d\n", id)
		return
	}
	fmt.Fprintf(w, "Deleted movie %d\n", id)
}
