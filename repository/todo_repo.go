package repository

import (
	"github.com/hoangduc02011998/todo_server/models"
)

type ToDoRepo interface {
	FindToDoByTask(task string) (models.ToDo, error)
	Insert(model models.ToDo) error
	Update(model models.ToDo) error
	Delete(id string) error
}
