package entity

import "time"

type Movie struct {
	ID        int
	Name      string
	Duration  time.Time
	Release   time.Time
	Genre     string
	Directors []string
	Actors    []string
	Writers   []string
	Country   string

	CreatedAt time.Time
	UpdatedAt time.Time
}
