package movies

import (
	"database/sql"
	"errors"
	"fmt"
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

func (h *Handler) CreateMovie(w http.ResponseWriter, r *http.Request) {
	// struct to hold movie data from request body
	var input struct {
		Title       string `json:"title"`
		ReleaseYear int    `json:"release_year"`
		Genre       string `json:"genre"`
		Runtime     int    `json:"runtime"`
	}

	// parse request body
	err := helper.ReadJSON(w, r, &input)
	if err != nil {
		helper.BadRequestResponse(w, r, h.logger, err)
		return
	}

	movie := Movie{
		Title:       input.Title,
		ReleaseYear: input.ReleaseYear,
		Genre:       input.Genre,
		Runtime:     input.Runtime,
	}

	// validate input
	validationErrors := validateMovie(movie)
	if len(validationErrors) > 0 {
		helper.FailedValidationResponse(w, r, h.logger, validationErrors)
		return
	}

	movie, err = h.svc.createMovie(r.Context(), movie)
	if err != nil {
		switch {
		case errors.Is(err, ErrDuplicateMovie):
			validationErrors["title, release year"] = "duplicate title and release year exists"
			helper.FailedValidationResponse(w, r, h.logger, validationErrors)
		default:
			helper.ServerErrorResponse(w, r, h.logger, err)
		}
		return
	}

	// add location header to new resource
	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("v1/movies/%d", movie.ID))

	err = helper.WriteJSON(w, http.StatusCreated, map[string]Movie{"movie": movie}, nil)
	if err != nil {
		helper.ServerErrorResponse(w, r, h.logger, err)
	}
}
