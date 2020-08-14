package repo

import "github.com/hoangduc02011998/todo-app/server/models"

type UserRepo interface {
	Insert(model models.User) error
	Login(model models.User) (*models.User, error)
	GetByUsername(username string) (*models.User, error)
}
