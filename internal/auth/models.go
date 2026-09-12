package auth

import (
	"time"
)

type Role string

const (
	Admin    Role = "admin"
	Customer Role = "customer"
)

type User struct {
	ID           int        `json:"id"`
	Email        string     `json:"email"`
	PasswordHash []byte     `json:"-"`
	Role         Role       `json:"role"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at,omitzero"`
}
