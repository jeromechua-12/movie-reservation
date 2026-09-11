package auth

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
)

type Repository struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) (*Repository) {
	return &Repository{db: db}
}

func (r *Repository) insertUser(ctx context.Context, email string, passwordHash string, role Role) error {
	query := `INSERT INTO users (email, password_hash, role, created_at)
	VALUES($1, $2, $3, UTC_TIMESTAMP)`

	_, err := r.db.ExecContext(ctx, query, email, passwordHash, role)
	if err != nil {
		return err
	}
	return nil
}
