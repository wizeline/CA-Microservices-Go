package entity

import "time"

type User struct {
	ID       int
	Name     string
	Username string
	Password string
	Email    string
	BirthDay time.Time
	Active   bool

	CreatedAt time.Time
	UpdatedAt time.Time
}
