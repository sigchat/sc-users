package model

import (
	"time"
)

type User struct {
	ID         int
	Username   string
	Password   []byte
	CreatedAt  time.Time
	LastOnline time.Time
	Active     bool
}
