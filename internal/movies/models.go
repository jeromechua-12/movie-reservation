package movies

import (
	"time"
)

type Movie struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	ReleaseYear int        `json:"release_year"`
	Genre       string     `json:"genre"`
	Runtime     int        `json:"runtime"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at,omitzero"`
}

func validateMovie(movie Movie) map[string]string {
	errors := make(map[string]string)

	if movie.Title == "" {
		errors["title"] = "title must be provided"
	}

	if movie.Genre == "" {
		errors["genre"] = "genre must be provided"
	}

	if movie.ReleaseYear == 0 {
		errors["release_year"] = "release year must be provided"
	}

	if movie.Runtime == 0 {
		errors["runtime"] = "runtime must be provided"
	}

	if movie.Runtime < 0 {
		errors["runtime"] = "runtime must be a positive integer"
	}

	return errors
}
