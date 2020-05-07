package repository

import (
	"github.com/hoangduc02011998/todo_server/models"
)

type ToDoRepo interface {
	FindAll() ([]models.ToDo, error)
	FindToDoByTask(task string) (models.ToDo, error)
	Insert(model models.ToDo) error
	Update(id string, status bool) error
	Delete(id string) error
}
