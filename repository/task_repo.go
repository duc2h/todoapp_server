package repo

import "github.com/hoangduc02011998/todo-app/server/models"

type TaskRepo interface {
	FindAll() ([]models.Task, error)
	FindTaskByName(task string) (*models.Task, error)
	Insert(model models.Task) error
	Update(name string, status bool) error
	Delete(name string) error
}
