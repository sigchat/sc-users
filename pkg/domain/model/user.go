package model

import (
	"time"
)

type Role string

const (
	RoleUser  Role = "user"
	RoleAdmin Role = "admin"
)

type Status string

const (
	StatusOnline  Status = "online"
	StatusOffline Status = "offline"
	StatusDND     Status = "dnd"
	StatusAFK     Status = "afk"
)

type User struct {
	ID            int        `json:"id"`
	Username      string     `json:"username"`
	Password      []byte     `json:"-"`
	CreatedAt     time.Time  `json:"createdAt"`
	LastUpdatedAt time.Time  `json:"lastUpdatedAt"`
	LastOnline    *time.Time `json:"lastOnline,omitempty"`
	Friends       []*User    `json:"-"`
	Status        Status     `json:"status"`
	Role          Role       `json:"role"`
	Active        bool       `json:"active"`
}

type Session struct {
	ID          int       `json:"sessionID"`
	AccessToken string    `json:"accessToken"`
	ExpiresAt   time.Time `json:"-"`
}
