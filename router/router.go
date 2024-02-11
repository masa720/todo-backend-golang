package router

import (
	"github.com/gin-gonic/gin"
	"github.com/masa720/todo-backend-golang/controller"
)

func NewRouter(tc controller.TodoController) *gin.Engine {
	r := gin.Default()

	r.GET("/api/v1/todos", tc.GetTodos)
	r.POST("/api/v1/todos", tc.Create)
	r.GET("/api/v1/todos/:id", tc.GetTodo)

	return r
}
