package movies

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/lib/pq"
)

var (
	ErrDuplicateMovie = errors.New("duplicate movie")
	ErrMovieNotFound  = errors.New("movie not found")
)

type Repository struct {
	db *sql.DB
}

func newRepo(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) insertMovie(ctx context.Context, movie Movie) (Movie, error) {
	query := `INSERT INTO movies (title, release_year, genre, runtime)
	VALUES ($1, $2, $3, $4)
	RETURNING id, created_at, updated_at`

	args := []any{movie.Title, movie.ReleaseYear, movie.Genre, movie.Runtime}

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	err := r.db.QueryRowContext(ctx, query, args...).Scan(&movie.ID, &movie.CreatedAt, &movie.UpdatedAt)
	if err != nil {
		if pqErr, ok := errors.AsType[*pq.Error](err); ok {
			if pqErr.Code == "23505" && pqErr.Constraint == "movies_title_release_year_key" {
				return Movie{}, ErrDuplicateMovie
			}
		}
		return Movie{}, err
	}

	return movie, nil
}

func (r *Repository) getMovieByID(ctx context.Context, id int) (Movie, error) {
	query := `SELECT id, title, release_year, genre, runtime, created_at, updated_at
	FROM movies
	WHERE id = $1`

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var movie Movie

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&movie.ID,
		&movie.Title,
		&movie.ReleaseYear,
		&movie.Genre,
		&movie.Runtime,
		&movie.CreatedAt,
		&movie.UpdatedAt,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return Movie{}, ErrMovieNotFound
		default:
			return Movie{}, err
		}
	}

	return movie, nil
}
