package model

import "time"

type Todo struct {
	Id          int
	Title       string
	Description string
	Deadline    time.Time
	IsDone      bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type TodoResponse struct {
	Todos []*Todo `json:"todos"`
}
