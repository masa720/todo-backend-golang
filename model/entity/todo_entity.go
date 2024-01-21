package entity

import "time"

type TodoEntity struct {
	Id          int
	Title       string
	Description string
	Deadline    time.Time
	IsDone      bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
