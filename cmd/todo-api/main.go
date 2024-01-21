package main

import (
	"net/http"

	"github.com/masa720/todo-backend-golang/controller"
	"github.com/masa720/todo-backend-golang/model/repository"
)

var tr = repository.NewTodoRepository()
var tc = controller.NewTodoController(tr)
var ro = controller.NewRouter(tc)

func main() {
	server := http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/todos", ro.HandleTodosRequest)
	server.ListenAndServe()
}
