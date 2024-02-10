package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/masa720/todo-backend-golang/controller"
)

func NewRouter(tc controller.TodoController) *gin.Engine {
	r := gin.Default()

	r.GET("/todos", tc.GetTodos)

	r.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello world")
	})

	return r
}
