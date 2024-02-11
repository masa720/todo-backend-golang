package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/masa720/todo-backend-golang/model"
	"github.com/masa720/todo-backend-golang/usecase"
	"gorm.io/gorm"
)

type TodoController interface {
	GetTodos(c *gin.Context)
	GetTodo(c *gin.Context)
	Create(c *gin.Context)
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

func (tc *todoController) GetTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	todo, err := tc.todoUsecase.GetTodo(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, todo)
}

func (tc *todoController) Create(c *gin.Context) {
	var todo *model.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := tc.todoUsecase.Create(todo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todo)
}
