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
	CreatedAt     time.Time  `json:"createdAt"`
	LastUpdatedAt time.Time  `json:"lastUpdatedAt"`
	LastOnline    *time.Time `json:"lastOnline,omitempty"`
	Status        string     `json:"status"`
	Active        bool       `json:"active"`
}
