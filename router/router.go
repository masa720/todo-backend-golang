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
	r.PUT("/api/v1/todos/:id", tc.Update)
	r.DELETE("/api/v1/todos/:id", tc.Delete)
	r.POST("/api/v1/todos/:id/done", tc.ToggleIsDone)
	return r
}
