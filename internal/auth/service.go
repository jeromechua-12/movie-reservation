package auth

import (
	"context"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo *Repository
}

func newService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (svc *Service) registerUser(ctx context.Context, email string, password string, role Role) (User, error) {
	// hash password using bcrypt
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return User{}, err
	}

	user := User{
		Email:        email,
		PasswordHash: passwordHash,
		Role:         role,
	}

	// insert to DB
	return svc.repo.insertUser(ctx, user)
}
