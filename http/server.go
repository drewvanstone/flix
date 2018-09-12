package http

import (
	"net/http"
)

func NewServeMux() *http.ServeMux { return http.NewServeMux() }

func ListenAndServe(port string, mux *http.ServeMux) {
	http.ListenAndServe(port, mux)
}
