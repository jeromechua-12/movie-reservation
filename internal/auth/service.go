package auth

import (
	"context"

	"golang.org/x/crypto/bcrypt"
)


type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (svc *Service) createUser(ctx context.Context, email string, password string, role Role) error {
	// validate credentials

	// hash password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	// insert to DB
	svc.repo.insertUser(ctx, email, string(passwordHash), role)
	return nil
}
