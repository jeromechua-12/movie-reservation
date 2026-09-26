package auth

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var ErrPasswordMismatch = errors.New("password do not match")

type Service struct {
	repo *Repository
}

func newService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (svc *Service) registerUser(ctx context.Context, email, password string, role Role) (User, error) {
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

func (svc *Service) authenticateUser(ctx context.Context, email, password string) (User, error) {
	user, err := svc.repo.getUserByEmail(ctx, email)
	if err != nil {
		return User{}, err
	}

	// validate hashed password
	err = bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(password))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return User{}, ErrPasswordMismatch
		default:
			return User{}, err
		}

	}

	return user, nil
}
