package repo

import "github.com/hoangduc02011998/todo_server/models"

type UserRepo interface {
	Insert(model models.User) error
}
