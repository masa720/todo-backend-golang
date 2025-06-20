package router

import (
	"github.com/gin-gonic/gin"
	"github.com/masa720/todo-backend-golang/controller"
	"github.com/masa720/todo-backend-golang/middleware"
)

func NewRouter(tc controller.TodoController, uc controller.UserController) *gin.Engine {
	r := gin.Default()

	// 認証が必要なエンドポイント
	authorized := r.Group("/api/v1")
	authorized.Use(middleware.AuthMiddleware())
	{
		authorized.GET("/todos", tc.GetTodos)
		authorized.POST("/todos", tc.Create)
		authorized.GET("/todos/:id", tc.GetTodo)
		authorized.PUT("/todos/:id", tc.Update)
		authorized.DELETE("/todos/:id", tc.Delete)
		authorized.POST("/todos/:id/done", tc.ToggleIsDone)
	}

	// 認証不要なエンドポイント
	r.POST("/api/v1/sign_up", uc.Signup)
	r.POST("/api/v1/sign_in", uc.Signin)
	return r
}
