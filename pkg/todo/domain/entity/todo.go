package entity

import "time"

type Todo struct {
	ID          string
	Title       string
	Description string
	Done        bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
