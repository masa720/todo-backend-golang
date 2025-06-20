package repository

import (
	"github.com/masa720/todo-backend-golang/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *model.User) error
	FindByEmail(mail string) (*model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (ur *userRepository) CreateUser(user *model.User) error {
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) FindByEmail(mail string) (*model.User, error) {
	var user model.User
	if err := ur.db.Where("mail = ?", mail).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
