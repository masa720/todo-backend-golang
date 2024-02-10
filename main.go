package main

import (
	"github.com/masa720/todo-backend-golang/controller"
	"github.com/masa720/todo-backend-golang/database"
	"github.com/masa720/todo-backend-golang/repository"
	"github.com/masa720/todo-backend-golang/router"
	"github.com/masa720/todo-backend-golang/usecase"
)

func main() {
	db := database.Init()
	tr := repository.NewTodoRepository(db)
	tu := usecase.NewTodoUsecase(tr)
	tc := controller.NewTodoController(tu)
	r := router.NewRouter(tc)
	r.Run(":8080")
}
