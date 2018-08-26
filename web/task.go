package web

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/drewvanstone/minerva"
)

type Handler struct {
	store minerva.StorageService
}

func NewHandler(store minerva.StorageService) *Handler {
	return &Handler{store}
}

func (h *Handler) TaskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.readTask(w, r)
	case "PUT":
		h.createTask(w, r)
	case "DELETE":
		h.deleteTask(w, r)
	default:
		fmt.Fprintf(w, "Something went wrong")
		return
	}
}

func (h *Handler) createTask(w http.ResponseWriter, r *http.Request) {
	task := r.URL.Query().Get("task")
	if err := h.store.Create(task); err != nil {
		fmt.Fprintf(w, "Error creating task %s\n", task)
		return
	}
	fmt.Fprintf(w, "Created task %s\n", task)
}

func (h *Handler) readTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		fmt.Fprintf(w, "Something went wrong")
		return
	}

	task, err := h.store.Read(id)
	if err != nil {
		fmt.Fprintf(w, "Error reading task %v\n", err)
		return
	}
	fmt.Fprintf(w, "Read task %v\n", *task)
}

func (h *Handler) deleteTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		fmt.Fprintf(w, "Something went wrong")
		return
	}

	if err := h.store.Delete(id); err != nil {
		fmt.Fprintf(w, "Error deleting task %d\n", id)
		return
	}
	fmt.Fprintf(w, "Deleted task %d\n", id)
}
