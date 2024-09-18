package model

import (
	"database/sql"
	"fmt"
	"time"
)

type RepoUser struct {
	ID        int64        `db:"id"`
	Name      string       `db:"name"`
	Email     string       `db:"email"`
	Password  string       `db:"password"`
	Role      Role         `db:"role"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

// Role /
type Role int64

const (
	UserRole  Role = iota // UserRole /
	AdminRole             // AdminRole /
)

func (r *Role) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	switch src := value.(type) {
	case int64:
		*r = Role(src)
		return nil
	}

	return fmt.Errorf("cannot scan %T", value)
}
