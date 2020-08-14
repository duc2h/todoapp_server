package biz

import (
	"github.com/hoangduc02011998/todo-app/server/models"
	repo "github.com/hoangduc02011998/todo-app/server/repository"
	"github.com/hoangduc02011998/todo-app/server/repository/repoIml"
)

func InitRepo() {
	InitUserRepo()
	InitTaskRepo()
}

var userRepo repo.UserRepo
var taskRepo repo.TaskRepo

func InitUserRepo() {
	userRepo = repoIml.NewUserRepo(models.UserDB.Collection)
}

func InitTaskRepo() {
	taskRepo = repoIml.NewTaskRepo(models.TaskDB.Collection)
}
