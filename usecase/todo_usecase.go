package usecase

import (
	"github.com/masa720/todo-backend-golang/model"
	"github.com/masa720/todo-backend-golang/repository"
)

type TodoUsecase interface {
	GetTodos() ([]*model.Todo, error)
	GetTodo(id uint) (*model.Todo, error)
}

type todoUsecase struct {
	todoRepository repository.TodoRepository
}

func NewTodoUsecase(tr repository.TodoRepository) TodoUsecase {
	return &todoUsecase{tr}
}

func (tu *todoUsecase) GetTodos() ([]*model.Todo, error) {
	return tu.todoRepository.FindAll()
}

func (tu *todoUsecase) GetTodo(id uint) (*model.Todo, error) {
	return tu.todoRepository.FindById(id)
}
