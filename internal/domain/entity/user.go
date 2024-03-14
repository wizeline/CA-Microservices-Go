package entity

import "time"

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	BirthDay  time.Time

	Username  string
	Passwd    string
	Active    bool
	LastLogin time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
}
