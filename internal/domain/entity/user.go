package entity

import (
	"database/sql"
	"time"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	BirthDay  time.Time

	Username  string
	Passwd    string
	Active    bool
	LastLogin sql.NullTime

	CreatedAt time.Time
	UpdatedAt sql.NullTime
}
