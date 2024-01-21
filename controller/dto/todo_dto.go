package dto

import "time"

type TodoResponse struct {
	Id          int       `json:"id"`
	UserId      int       `json:"userId"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Deadline    time.Time `json:"deadline"`
	IsDone      bool      `json:"isDone"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type TodoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type TodosResponse struct {
	Todos []TodoResponse `json:"todos"`
}
