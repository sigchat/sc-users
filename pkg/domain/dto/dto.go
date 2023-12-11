package dto

import "time"

type RegisterUserDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UpdateUserDTO struct {
	Username       string
	HashedPassword []byte
	LastOnline     time.Time
	Active         bool
}
