package model

import (
	"time"
)

type User struct {
	ID            int        `json:"id"`
	Username      string     `json:"username"`
	Password      []byte     `json:"-"`
	CreatedAt     time.Time  `json:"created_at"`
	LastUpdatedAt time.Time  `json:"last_updated_at"`
	LastOnline    *time.Time `json:"last_online,omitempty"`
	Active        bool       `json:"active"`
}

type Session struct {
	ID          int    `json:"session_id"`
	AccessToken string `json:"access_token"`
}
