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
