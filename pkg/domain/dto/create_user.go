package dto

import "time"

type CreateUserDTO struct {
	Username       string
	HashedPassword []byte
}

type UpdateUserDTO struct {
	Username       string
	HashedPassword []byte
	LastOnline     time.Time
	Active         bool
}
