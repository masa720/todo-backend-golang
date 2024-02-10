package repository

import (
	"github.com/masa720/todo-backend-golang/model"
	"gorm.io/gorm"
)

type TodoRepository interface {
	GetTodos() ([]*model.Todo, error)
}

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &todoRepository{db}
}

func (tr *todoRepository) GetTodos() ([]*model.Todo, error) {
	var todos []*model.Todo
	if err := tr.db.Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}
