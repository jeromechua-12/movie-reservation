package auth

import (
	"time"
)

type Role string

const (
	Admin Role = "admin"
	Customer Role = "customer"
)

type User struct {
	id int
	email string
	passwordHash []byte
	role Role
	createdAt time.Time
}


