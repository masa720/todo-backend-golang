package router

import (
	"github.com/gin-gonic/gin"
	"github.com/masa720/todo-backend-golang/controller"
)

func NewRouter(tc controller.TodoController) *gin.Engine {
	r := gin.Default()

	r.GET("/todos", tc.GetTodos)

	r.GET("/todos/:id", tc.GetTodo)

	return r
}
