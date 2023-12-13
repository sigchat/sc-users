package dto

import (
	"time"
)

type RegisterUserRequestDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterUserResponseDTO struct {
	UserID      int    `json:"user_id"`
	AccessToken string `json:"access_token"`
}

type LoginUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginUserResponse struct {
	UserID      int    `json:"user_id"`
	AccessToken string `json:"access_token"`
}

type UpdateUserDTO struct {
	Username       string
	HashedPassword []byte
	LastOnline     time.Time
	Active         bool
}

type UserInfoDTO struct {
	ID            int        `json:"id"`
	Username      string     `json:"username"`
	CreatedAt     time.Time  `json:"created_at"`
	LastUpdatedAt time.Time  `json:"last_updated_at"`
	LastOnline    *time.Time `json:"last_online,omitempty"`
	Active        bool       `json:"active"`
}
