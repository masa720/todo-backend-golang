package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/masa720/todo-backend-golang/model"
	"github.com/masa720/todo-backend-golang/usecase"
)

type TodoController interface {
	GetTodos(c *gin.Context)
}

type todoController struct {
	todoUsecase usecase.TodoUsecase
}

func NewTodoController(tu usecase.TodoUsecase) TodoController {
	return &todoController{tu}
}

func (tc *todoController) GetTodos(c *gin.Context) {
	todos, err := tc.todoUsecase.GetTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := model.TodoResponse{Todos: todos}
	c.JSON(http.StatusOK, response)
}
