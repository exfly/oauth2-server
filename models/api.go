package models

import (
	"time"
)

// User user for user resources
type User struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	Role      string
	Username  string
}
