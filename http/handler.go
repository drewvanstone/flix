package http

import "github.com/drewvanstone/flix"

type FlixHandler struct {
	store flix.StorageService
}

func NewFlixHandler(store flix.StorageService) *FlixHandler {
	return &FlixHandler{store}
}
