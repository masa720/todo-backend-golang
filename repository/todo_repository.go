package repository

import (
	"github.com/masa720/todo-backend-golang/model"
	"gorm.io/gorm"
)

type TodoRepository interface {
	FindAll() ([]*model.Todo, error)
	FindById(id uint) (*model.Todo, error)
	Create(todo *model.Todo) error
}

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &todoRepository{db}
}

func (tr *todoRepository) FindAll() ([]*model.Todo, error) {
	var todos []*model.Todo
	if err := tr.db.Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func (tr *todoRepository) FindById(id uint) (*model.Todo, error) {
	var todo *model.Todo
	if err := tr.db.First(&todo, id).Error; err != nil {
		return nil, err
	}
	return todo, nil
}

func (tr *todoRepository) Create(todo *model.Todo) error {
	return tr.db.Create(todo).Error
}
