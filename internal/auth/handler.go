package auth

import (
	"net/http"
)

type Handler struct {
	svc *Service
}

func NewHandler(svc *Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) createUser(w http.ResponseWriter, r *http.Request) {
	// parse data

	// call service
}
