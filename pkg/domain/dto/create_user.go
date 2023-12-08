package dto

import (
	"time"
)

type CreateUserRequest struct {
	Username string
	Password []byte
}

type CreateUserResponse struct {
	ID         int
	Username   string
	Password   []byte
	CreatedAt  time.Time
	LastOnline time.Time
	Active     bool
}
