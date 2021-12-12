package model

import "time"

type Entry struct {
	ID        int
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
