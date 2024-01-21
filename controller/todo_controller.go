package controller

import (
	"encoding/json"
	"net/http"

	"github.com/masa720/todo-backend-golang/controller/dto"
	"github.com/masa720/todo-backend-golang/model/repository"
)

type TodoController interface {
	GetTodos(w http.ResponseWriter, r *http.Request)
}

type todoController struct {
	tr repository.TodoRepository
}

func NewTodoController(tr repository.TodoRepository) TodoController {
	return &todoController{tr}
}

func (tc *todoController) GetTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := tc.tr.GetTodos()
	if err != nil {
		w.WriteHeader(500)
		return
	}

	var todoResponses []dto.TodoResponse
	for _, v := range todos {
		todoResponses = append(todoResponses, dto.TodoResponse{Id: v.Id, Title: v.Title, Description: v.Description, Deadline: v.Deadline, IsDone: v.IsDone, CreatedAt: v.CreatedAt, UpdatedAt: v.UpdatedAt})
	}

	var todosResponse dto.TodosResponse
	todosResponse.Todos = todoResponses

	output, _ := json.MarshalIndent(todosResponse, "", "\t\t")

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}
