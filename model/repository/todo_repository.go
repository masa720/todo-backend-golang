package repository

import (
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/masa720/todo-backend-golang/model/entity"
)

type TodoRepository interface {
	GetTodos() (todos []entity.TodoEntity, err error)
	// InsertTodo(todo entity.TodoEntity) (id int, err error)
	// UpdateTodo(todo entity.TodoEntity) (err error)
	// DeleteTodo(id int) (err error)
}

type todoRepository struct {
}

func NewTodoRepository() TodoRepository {
	return &todoRepository{}
}

func (tr *todoRepository) GetTodos() (todos []entity.TodoEntity, err error) {
	todos = []entity.TodoEntity{}
	rows, err := Db.Query("SELECT id, title, description, deadline, is_done, created_at, updated_at FROM todos ORDER BY id DESC")
	if err != nil {
		log.Print(err)
		return
	}

	for rows.Next() {
		todo := entity.TodoEntity{}
		err = rows.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Deadline, &todo.IsDone, &todo.CreatedAt, &todo.UpdatedAt)
		if err != nil {
			log.Print(err)
			return
		}
		todos = append(todos, todo)
	}

	return
}
