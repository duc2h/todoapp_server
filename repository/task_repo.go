package repo

import "github.com/hoangduc02011998/todo_server/models"

type TaskRepo interface {
	FindAll() ([]models.Task, error)
	FindTaskByTask(task string) (models.Task, error)
	Insert(model models.Task) error
	Update(id string, status bool) error
	Delete(id string) error
}
