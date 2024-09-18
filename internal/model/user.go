package model

import (
	"database/sql"
	"time"
)

type User struct {
	ID        int64
	Name      string
	Email     string
	Password  string
	Role      Role
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

// Role /
type Role int64

const (
	UserRole  Role = iota // UserRole /
	AdminRole             // AdminRole /
)
