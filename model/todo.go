package model

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Deadline    time.Time      `json:"deadline"`
	IsDone      bool           `json:"isDone" gorm:"default:false"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

type TodoResponse struct {
	Todos []*Todo `json:"todos"`
}
