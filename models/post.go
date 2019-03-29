package models

import "time"

type Post struct {
	ID        uint64
	Text      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
