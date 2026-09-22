package auth

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/lib/pq"
)

var (
	ErrDuplicateEmail = errors.New("duplicate email")
	ErrEmailNotFound  = errors.New("email not found")
)

type Repository struct {
	db *sql.DB
}

func newRepo(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) insertUser(ctx context.Context, user User) (User, error) {
	query := `INSERT INTO users (email, password_hash, role)
	VALUES ($1, $2, $3)
	RETURNING id, created_at, updated_at`

	args := []any{user.Email, user.PasswordHash, user.Role}

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	err := r.db.QueryRowContext(ctx, query, args...).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if pqErr, ok := errors.AsType[*pq.Error](err); ok {
			if pqErr.Code == "23505" && pqErr.Constraint == "users_email_key" {
				return User{}, ErrDuplicateEmail
			}
		}
		return User{}, err
	}

	return user, nil
}

func (r *Repository) getUserByEmail(ctx context.Context, email string) (User, error) {
	query := `SELECT id, email, password_hash, role, created_at, updated_at
			  FROM users
			  WHERE email = $1`

	var user User

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.Role,
		&user.CreatedAt,
		&user.PasswordHash,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return User{}, ErrEmailNotFound
		default:
			return User{}, err
		}
	}

	return user, nil
}
