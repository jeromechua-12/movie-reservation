package auth

import (
	"database/sql"
	"errors"
	"log/slog"
	"net/http"

	"github.com/jeromechua-12/movie-reservation/internal/helper"
)

type Handler struct {
	svc    *Service
	logger *slog.Logger
}

func NewHandler(db *sql.DB, logger *slog.Logger) *Handler {
	repo := newRepo(db)
	svc := newService(repo)
	return &Handler{svc: svc, logger: logger}
}

func (h *Handler) RegisterCustomer(w http.ResponseWriter, r *http.Request) {
	// struct to hold expected data from request body
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// parse request body
	err := helper.ReadJSON(w, r, &input)
	if err != nil {
		helper.BadRequestResponse(w, r, h.logger, err)
		return
	}

	// validate email and password
	validationErrors := validateRegistration(input.Email, input.Password)
	if len(validationErrors) > 0 {
		helper.FailedValidationResponse(w, r, h.logger, validationErrors)
		return
	}

	user, err := h.svc.registerUser(r.Context(), input.Email, input.Password, Customer)
	if err != nil {
		switch {
		case errors.Is(err, ErrDuplicateEmail):
			validationErrors["email"] = "email address already exists"
			helper.FailedValidationResponse(w, r, h.logger, validationErrors)
		default:
			helper.ServerErrorResponse(w, r, h.logger)
		}
		return
	}

	err = helper.WriteJSON(w, http.StatusCreated, map[string]User{"user": user}, nil)
	if err != nil {
		helper.ServerErrorResponse(w, r, h.logger)
	}
}
